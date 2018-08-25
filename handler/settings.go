package handler

import (
	"bytes"
	"encoding/base64"
	"image/png"
	"net/http"

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
	user := checkSession(r)
	if user != "" {
		// define variables
		special := 0
		var reloadPage bool

		// change language
		lang := r.PostFormValue("language")
		if lang != "" {
			// change language cookie
			http.SetCookie(w, &http.Cookie{Name: "secpass_lang", Value: lang})

			// reload page
			reloadPage = true
			return
		}

		// change password
		currentPassword, newPassword, repeatNewPassword := r.PostFormValue("currentPassword"), r.PostFormValue("newPassword"), r.PostFormValue("repeatNewPassword")
		if currentPassword != "" && len(newPassword) >= 8 && len(repeatNewPassword) >= 8 {
			// check passwords match
			if newPassword == repeatNewPassword {
				// define variables to write into
				var username string
				var passwordHash string

				// read user from database
				errQuery := conf.DB.QueryRow(conf.GetSQL("login"), user).Scan(&username, &passwordHash)
				shorts.Check(errQuery, true)

				// compare passwords
				if bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(currentPassword)) == nil && user == username {
					// generate bcrypt hash
					password, errPassword := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost+1)
					shorts.Check(errPassword, true)

					// change password in db and check error
					_, errExec := conf.DB.Exec(conf.GetSQL("change_password"), string(password), user)
					shorts.Check(errExec, true)

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
				// delete account in db, sessions, cookies
				conf.DB.Exec(conf.GetSQL("delete_account"), user)
				for sessionUUID, session := range spuser.Sessions {
					if session.User == user {
						delete(spuser.Sessions, sessionUUID)
					}
				}

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
			var oneTimePasswordWrong bool

			// enable two factor authentication
			oneTimePassword, twoFactorSecret := r.PostFormValue("oneTimePassword"), r.PostFormValue("twoFactorSecret")
			if twoFactorSecret != "" && len(oneTimePassword) >= 6 && len(oneTimePassword) <= 8 {
				// validate one-time password
				if totp.Validate(oneTimePassword, twoFactorSecret) {
					// write to db
					spuser.EnableTwoFactor(user, twoFactorSecret)

					// reload page
					reloadPage = true
					return
				}

				// one-time password is not correct
				oneTimePasswordWrong = true
			}

			// two-factor authentication key generation
			key, errKey := totp.Generate(totp.GenerateOpts{Issuer: "secpass", AccountName: user, Algorithm: otp.AlgorithmSHA512})
			shorts.Check(errKey, true)

			// image generation
			image, errImg := key.Image(200, 200)
			shorts.Check(errImg, true)

			// image encoding
			buf := bytes.Buffer{}
			png.Encode(&buf, image)
			encodedImage := base64.StdEncoding.EncodeToString(buf.Bytes())

			// set two-factor data
			twoFactorData = TwoFactorData{Image: encodedImage, Secret: key.Secret(), Disabled: true, OneTimePasswordWrong: oneTimePasswordWrong}
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
				return
			}

			// set two-factor data
			twoFactorData = TwoFactorData{OneTimePasswordWrong: true}
		}

		if reloadPage {
			w.Header().Set("location", r.URL.Path)
			w.WriteHeader(http.StatusSeeOther)
		}

		// execute template
		shorts.Check(tpl.ExecuteTemplate(w, "settings.html", Data{User: user, Lang: getLang(r), Special: special, TwoFactor: twoFactorData}), false)
	}

	// redirect to login
	redirect(w, "/login")
}
