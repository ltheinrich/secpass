package handler

import (
	"database/sql"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"lheinrich.de/secpass/conf"
	"lheinrich.de/secpass/shorts"
)

// Register function
func Register(w http.ResponseWriter, r *http.Request) {
	// check logged in and redirect to index if so
	if checkSession(w, r) != "" {
		redirect(w, "/")
		return
	}

	// output text if equals special
	special := 0

	// define values
	name, password, repeat := r.PostFormValue("name"), r.PostFormValue("password"), r.PostFormValue("repeat")

	// check for input
	if name != "" && password != "" && repeat != "" && len(password) >= 8 && len(repeat) >= 8 {
		// check whether passwords match
		if password == repeat {
			// check whether name already exists
			var queryName string
			errQuery := conf.DB.QueryRow(conf.GetSQL("get_password"), name).Scan(&queryName)

			// name does not exist
			if errQuery == sql.ErrNoRows {
				// generate random key
				key := shorts.Encrypt(shorts.UUID(), shorts.GenerateKey(shorts.UUID()))

				// encrypt key
				encryptedKey := shorts.Encrypt(key, shorts.GenerateKey(password))

				// hash password
				passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost+1)

				// insert user into db and add default category
				_, errExecRegister := conf.DB.Exec(conf.GetSQL("register"), name, string(passwordHash), "", encryptedKey)
				_, errExecADC := conf.DB.Exec(conf.GetSQL("add_default_category"), name)

				// check for error
				shorts.Check(errExecRegister)
				shorts.Check(errExecADC)

				// redirect and return
				redirectTemp(w, "/login")
				return
			}

			// check error
			shorts.Check(errQuery)

			// name exists, print
			special = 1
		} else {
			// passwords does not match, print
			special = 2
		}
	}

	// execute template
	shorts.Check(tpl.ExecuteTemplate(w, "register.html", Data{User: "", Lang: getLang(r), Special: special, LoggedOut: true}))
}
