package handler

import (
	"net/http"
	"time"

	"golang.org/x/crypto/bcrypt"

	"lheinrich.de/secpass/conf"
	"lheinrich.de/secpass/shorts"
	"lheinrich.de/secpass/user"
)

// Login function
func Login(w http.ResponseWriter, r *http.Request) {
	// check logged in and redirect to index if so
	if checkSession(r) != "" {
		redirect(w, "/")
		return
	}

	// output login data wrong
	special := 0

	// define values
	name, password := r.PostFormValue("name"), r.PostFormValue("password")

	// check for input
	if name != "" && password != "" {
		// read password hash from database and check for error
		var username string
		var passwordHash string
		err := conf.DB.QueryRow("SELECT name, password FROM users WHERE name = $1", name).Scan(&username, &passwordHash)
		shorts.Check(err, true)

		// compare passwords
		if bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password)) == nil {
			// login done

			// generate uuid and define session expiration date
			uuid := shorts.UUID()
			expires := time.Now().Add(time.Hour)

			// define cookies
			cookieUUID := http.Cookie{Name: "secpass_uuid", Value: uuid, Expires: expires}
			cookieName := http.Cookie{Name: "secpass_name", Value: username, Expires: expires}

			// add session and set cookies
			user.Sessions[uuid] = user.Session{User: username, Expires: expires}
			http.SetCookie(w, &cookieUUID)
			http.SetCookie(w, &cookieName)

			// redirect to index
			redirect(w, "/")
			return
		}
		special = 1
	}

	// execute template
	shorts.Check(tpl.ExecuteTemplate(w, "login.html", Data{User: "", Lang: getLang(r), Special: special}), false)
}
