package handler

import (
	"net/http"

	"lheinrich.de/secpass/shorts"
	"lheinrich.de/secpass/user"
)

// Index function
func Index(w http.ResponseWriter, r *http.Request) {
	// execute template
	shorts.Check(tpl.ExecuteTemplate(w, "index.html", user.User{ID: -1, Lang: "de"}), false)
}
