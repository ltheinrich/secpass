package handler

import (
	"net/http"

	"lheinrich.de/secpass/shorts"
)

// Settings function
func Settings(w http.ResponseWriter, r *http.Request) {
	// check session
	user := checkSession(r)
	if user != "" {
		// change language
		lang := r.PostFormValue("language")
		if lang != "" {
			http.SetCookie(w, &http.Cookie{Name: "secpass_lang", Value: lang})
			w.Header().Set("location", r.URL.Path)
			w.WriteHeader(http.StatusSeeOther)
			return
		}

		// execute template
		shorts.Check(tpl.ExecuteTemplate(w, "settings.html", Data{User: user, Lang: getLang(r)}), false)
	}

	// redirect to login
	redirect(w, "/login")
}
