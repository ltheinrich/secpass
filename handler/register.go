package handler

import (
	"database/sql"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"lheinrich.de/secpass/conf"
	"lheinrich.de/secpass/shorts"
	"lheinrich.de/secpass/user"
)

// Register function
func Register(w http.ResponseWriter, r *http.Request) {
	// output text if equals special
	special := -1

	// define values
	name, password, repeat := r.PostFormValue("name"), r.PostFormValue("password"), r.PostFormValue("repeat")

	// check for input
	if name != "" && password != "" && repeat != "" {
		// check whether passwords match
		if password == repeat {
			// check whether name already exists
			errQuery := conf.DB.QueryRow("SELECT password FROM users WHERE name = $1", name).Scan(nil)
			shorts.Check(errQuery, false)

			// name does not exist
			if errQuery == sql.ErrNoRows {
				// hash password and insert user
				passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost+1)
				_, errExec := conf.DB.Exec("INSERT INTO users (name, password) VALUES ($1, $2)", name, string(passwordHash))
				shorts.Check(errExec, true)

				// redirect and return
				redirect(w, "/login")
				return
			}

			// name exists, print
			special = -2
		} else {
			// passwords does not match, print
			special = -3
		}
	}

	// execute template
	shorts.Check(tpl.ExecuteTemplate(w, "register.html", user.User{ID: special, Lang: "de"}), false)
}
