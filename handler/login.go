package handler

import (
	"io"
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"lheinrich.de/secpass/conf"
	"lheinrich.de/secpass/shorts"
	"lheinrich.de/secpass/user"
)

// Login function
func Login(w http.ResponseWriter, r *http.Request) {
	// output login data wrong
	special := -1

	// define values
	name, password := r.PostFormValue("name"), r.PostFormValue("password")

	// check for input
	if name != "" && password != "" {
		// read password hash from database and check for error
		var passwordHash string
		err := conf.DB.QueryRow("SELECT password FROM users WHERE name = $1", name).Scan(&passwordHash)
		shorts.Check(err, true)

		// compare passwords
		if bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password)) == nil {
			// login done
			io.WriteString(w, "true")
			return
		} else {
			special = -2
		}
	}

	// execute template
	shorts.Check(tpl.ExecuteTemplate(w, "login.html", user.User{ID: special, Lang: conf.Lang["de"]}), false)
}
