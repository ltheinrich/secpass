package handler

import (
	"bytes"
	"encoding/base64"
	"image/png"
	"net/http"
	"strconv"
	"time"

	"github.com/pquerna/otp"
	"github.com/pquerna/otp/totp"

	"lheinrich.de/secpass/conf"
	"lheinrich.de/secpass/spuser"

	"golang.org/x/crypto/bcrypt"

	"lheinrich.de/secpass/shorts"
)

// Settings function
func Settings(w http.ResponseWriter, r *http.Request) {
	// check session
	user := checkSession(w, r)
	if user != "" {
		// define variables
		special := 0
		var reloadPage bool

		// set special from url
		specialValue := r.FormValue("special")
		if specialValue != "" {
			special, _ = strconv.Atoi(specialValue)
		}

		// change language
		lang := r.PostFormValue("language")
		if lang != "" {
			// define date
			date := time.Now()
			year := strconv.Itoa(date.Year() + 1)
			month := date.Month().String()
			day := strconv.Itoa(date.Day())
			expires, _ := time.Parse("20060102", year+month+day)

			// change language cookie
			http.SetCookie(w, &http.Cookie{Name: "secpass_lang", Value: lang, Expires: expires})

			// reload page
			reloadPage = true
		}

		// change password
		currentPassword, newPassword, repeatNewPassword := r.PostFormValue("currentPassword"), r.PostFormValue("newPassword"), r.PostFormValue("repeatNewPassword")
		if currentPassword != "" && len(newPassword) >= 8 && len(repeatNewPassword) >= 8 {
			// check passwords match
			if newPassword == repeatNewPassword {
				// define variables to write into
				var username string
				var passwordHash string
				var secret string
				var key string

				// read user from database
				errQuery := conf.DB.QueryRow(conf.GetSQL("login"), user).Scan(&username, &passwordHash, &secret, &key)
				shorts.Check(errQuery)

				// compare passwords
				if bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(currentPassword)) == nil && user == username {
					// generate bcrypt hash
					password, errPassword := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost+1)
					shorts.Check(errPassword)

					// re-encrypt key
					decryptedKey := shorts.Decrypt(key, shorts.GenerateKey(currentPassword))
					encryptedKey := shorts.Encrypt(decryptedKey, shorts.GenerateKey(newPassword))

					// change password in db and check error
					_, errExec := conf.DB.Exec(conf.GetSQL("change_password"), string(password), encryptedKey, user)
					shorts.Check(errExec)

					// password changed successfully
					special = -1
				} else {
					// current password is wrong
					special = -2
				}
			} else {
				// passwords are not equal to each other
				special = -3
			}
		}

		// delete account forever
		if r.URL.Path == "/settings/delete_forever" {
			deleteForever := r.PostFormValue("delete_forever")
			if deleteForever == "delete_account_forever" {
				// delete passwords in db
				_, errPasswords := conf.DB.Exec(conf.GetSQL("delete_passwords"), user)
				shorts.Check(errPasswords)

				// delete categories in db
				_, errCategories := conf.DB.Exec(conf.GetSQL("delete_categories"), user)
				shorts.Check(errCategories)

				// delete account in db
				_, errAccount := conf.DB.Exec(conf.GetSQL("delete_account"), user)
				shorts.Check(errAccount)

				// delete session
				for sessionUUID, session := range spuser.Sessions {
					if session.User == user {
						delete(spuser.Sessions, sessionUUID)
					}
				}

				// delete cookie
				deleteAllCookies(w)

				// redirect to index page
				redirect(w, "/")
				return
			}
			// display checkbox
			special = -4
		}

		var twoFactorData TwoFactorData

		// check whether two-factor authentication is disabled (== "") or enabled
		if spuser.TwoFactorSecret(user) == "" {
			var skip bool
			var oneTimePasswordWrong bool

			// enable two factor authentication
			oneTimePassword, twoFactorSecret := r.PostFormValue("oneTimePassword"), r.PostFormValue("twoFactorSecret")
			if twoFactorSecret != "" && (len(oneTimePassword) == 6 || len(oneTimePassword) == 8) {
				// validate one-time password
				if totp.Validate(oneTimePassword, twoFactorSecret) {
					// write to db
					spuser.EnableTwoFactor(user, twoFactorSecret)

					// reload page
					reloadPage = true

					// skip key generation
					skip = true
				} else {
					// one-time password is not correct
					oneTimePasswordWrong = true
				}
			}

			if !skip {
				// two-factor authentication key generation
				key, errKey := totp.Generate(totp.GenerateOpts{Issuer: "secpass", AccountName: user, Algorithm: otp.AlgorithmSHA512})
				shorts.Check(errKey)

				// image generation
				image, errImg := key.Image(200, 200)
				shorts.Check(errImg)

				// image encoding
				buf := bytes.Buffer{}
				png.Encode(&buf, image)
				encodedImage := base64.StdEncoding.EncodeToString(buf.Bytes())

				// set two-factor data
				twoFactorData = TwoFactorData{Image: encodedImage, Secret: key.Secret(), Disabled: true, OneTimePasswordWrong: oneTimePasswordWrong}
			}
		} else if r.PostFormValue("disableTwoFactor") == "disableTwoFactorAuthentication" {
			// define variables
			oneTimePassword := r.PostFormValue("oneTimePassword")
			twoFactorSecret := spuser.TwoFactorSecret(user)

			// validate one-time password
			if len(oneTimePassword) >= 6 && len(oneTimePassword) <= 8 && totp.Validate(oneTimePassword, twoFactorSecret) {
				// disable two-factor authentication
				spuser.DisableTwoFactor(user)

				// reload page
				reloadPage = true
			} else {
				// set two-factor data
				twoFactorData = TwoFactorData{OneTimePasswordWrong: true}
			}
		}

		if reloadPage {
			redirect(w, "/settings?special="+strconv.Itoa(special))
			return
		}

		// execute template
		shorts.Check(tpl.ExecuteTemplate(w, "settings.html", Data{User: user, Lang: getLang(r), Special: special, TwoFactor: twoFactorData}))
	}

	// redirect to login
	redirect(w, "/login")
}
