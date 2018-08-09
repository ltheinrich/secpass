package handler

import (
	"html/template"
	"net/http"
	"strings"

	"lheinrich.de/secpass/conf"
	"lheinrich.de/secpass/shorts"
	"lheinrich.de/secpass/user"
)

// Data to pass into template
type Data struct {
	User    string
	Lang    string
	Special int
}

var (
	// define functions
	funcs = template.FuncMap{
		"config": func(key string) string {
			// split key string by . to get group.key and return value
			keys := strings.Split(key, ".")
			return conf.Config[keys[0]][keys[1]]
		},
		"lang": func(keys ...interface{}) string {
			return (*conf.Lang[keys[0].(string)])[keys[1].(string)]
		}}

	// define templates
	tpl *template.Template
)

// name function
func name(w http.ResponseWriter, r *http.Request) {
	// execute template
	shorts.Check(tpl.ExecuteTemplate(w, "name", nil), false)
}

// LoadTemplates parse
func LoadTemplates() {
	// parse templates and check error
	var err error
	tpl, err = template.New("").Funcs(funcs).ParseGlob(conf.Config["webserver"]["templatesDirectory"])
	shorts.Check(err, true)
}

// redirect to location
func redirect(w http.ResponseWriter, location string) {
	w.Header().Set("location", location)
	w.WriteHeader(307)
}

// cookiesExist check whether the cookies exist
func cookiesExist(r *http.Request, names ...string) bool {
	for _, name := range names {
		_, err := r.Cookie(name)
		if err != nil {
			return false
		}
	}
	return true
}

// getLang get language or default language
func getLang(r *http.Request) string {
	lang, err := r.Cookie("secpass_lang")
	if err != nil {
		return conf.Config["app"]["defaultLanguage"]
	}
	return lang.Value
}

// checkSession check whether logged in and return username or empty string
func checkSession(r *http.Request) string {
	if cookiesExist(r, "secpass_uuid", "secpass_name") {
		cookieUUID, _ := r.Cookie("secpass_uuid")
		cookieName, _ := r.Cookie("secpass_name")
		uuid := cookieUUID.Value
		name := cookieName.Value
		user := user.Sessions[uuid].User

		if user == name {
			return user
		}
	}
	return ""
}
