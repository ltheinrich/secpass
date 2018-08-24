package handler

import (
	"net/http"

	spuser "lheinrich.de/secpass/user"

	"lheinrich.de/secpass/conf"

	"golang.org/x/crypto/bcrypt"

	"lheinrich.de/secpass/shorts"
)

// Settings function
func Settings(w http.ResponseWriter, r *http.Request) {
	// check session
	user := checkSession(r)
	if user != "" {
		// define special
		special := 0

		// change language
		lang := r.PostFormValue("language")
		if lang != "" {
			// change language cookie
			http.SetCookie(w, &http.Cookie{Name: "secpass_lang", Value: lang})

			// reload page
			w.Header().Set("location", r.URL.Path)
			w.WriteHeader(http.StatusSeeOther)
			return
		}

		// change password
		currentPassword, newPassword, repeatNewPassword := r.PostFormValue("currentPassword"), r.PostFormValue("newPassword"), r.PostFormValue("repeatNewPassword")
		if currentPassword != "" && newPassword != "" && repeatNewPassword != "" && len(newPassword) >= 8 && len(repeatNewPassword) >= 8 {
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

		// execute template
		shorts.Check(tpl.ExecuteTemplate(w, "settings.html", Data{User: user, Lang: getLang(r), Special: special}), false)
	}

	// redirect to login
	redirect(w, "/login")
}
