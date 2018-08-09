package handler

import (
	"net/http"

	"lheinrich.de/secpass/shorts"
)

// Index function
func Index(w http.ResponseWriter, r *http.Request) {
	// check session
	user := checkSession(r)
	if user != "" {
		// execute template
		shorts.Check(tpl.ExecuteTemplate(w, "index.html", Data{User: user, Lang: getLang(r)}), false)
	}

	// redirect to login
	redirect(w, "/login")
}
