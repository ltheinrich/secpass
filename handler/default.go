package handler

import (
	"html/template"
	"net/http"
	"time"

	"lheinrich.de/secpass/conf"
	"lheinrich.de/secpass/shorts"
	"lheinrich.de/secpass/spuser"
)

// Data to pass into template
type Data struct {
	User            string
	Lang            string
	Special         int
	LoggedOut       bool
	TwoFactor       TwoFactorData
	Entry           Password
	Passwords       []Password
	Pwns            []string
	Categories      []Category
	DefaultCategory Category
}

// Category to pass with Data
type Category struct {
	ID   int
	Name string
}

// TwoFactorData data for two-factor authentication
type TwoFactorData struct {
	Image                string
	Secret               string
	Disabled             bool
	OneTimePasswordWrong bool
}

var (
	// define functions
	funcs = template.FuncMap{
		"config": func(keys ...string) string {
			return conf.Config[keys[0]][keys[1]]
		},
		"lang": func(keys ...interface{}) string {
			return conf.Lang[keys[0].(string)][keys[1].(string)]
		},
		"languages": func() []string {
			languages := []string{}
			for lang := range conf.Lang {
				languages = append(languages, lang)
			}

			return languages
		}}

	// define templates
	tpl *template.Template

	// let cookies expire
	expiresCookie = time.Now().Add(-100 * time.Hour)
)

// name function
func name(w http.ResponseWriter, r *http.Request) {
	// execute template
	shorts.Check(tpl.ExecuteTemplate(w, "name", nil))
}

// LoadTemplates parse
func LoadTemplates() {
	// parse templates and check error
	var err error
	tpl, err = template.New("").Funcs(funcs).ParseGlob(conf.Config["webserver"]["templatesDirectory"])
	shorts.Check(err)
}

// redirect to location
func redirect(w http.ResponseWriter, location string) {
	// change location and write status code
	w.Header().Set("location", location)
	w.WriteHeader(http.StatusSeeOther)
}

// redirect to location temporary
func redirectTemp(w http.ResponseWriter, location string) {
	// change location and write status code
	w.Header().Set("location", location)
	w.WriteHeader(http.StatusTemporaryRedirect)
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
	// get cookie and use default it not existing
	lang, err := r.Cookie("secpass_lang")
	if err != nil {
		return conf.Config["app"]["defaultLanguage"]
	}

	// return value
	return lang.Value
}

// checkSession check whether logged in and return username or empty string
func checkSession(w http.ResponseWriter, r *http.Request) string {
	if cookiesExist(r, "secpass_uuid", "secpass_name") {
		// define cookies
		cookieUUID, _ := r.Cookie("secpass_uuid")
		cookieName, _ := r.Cookie("secpass_name")
		cookieHash, _ := r.Cookie("secpass_hash")

		// get cookie values
		uuid := cookieUUID.Value
		name := cookieName.Value

		// session data
		session := spuser.Sessions[uuid]
		user := session.User

		if user == name {
			// define expires time
			expires := time.Now().Add(10 * time.Minute)

			// change session expires time
			session.Expires = expires
			spuser.Sessions[uuid] = session

			// change cookie expires time
			cookieUUID.Expires = expires
			cookieName.Expires = expires
			cookieHash.Expires = expires

			// update cookies
			http.SetCookie(w, cookieUUID)
			http.SetCookie(w, cookieName)
			http.SetCookie(w, cookieHash)

			// return user
			return user
		}
	}

	// return logged out
	return ""
}

// cookie return cookie value
func cookie(r *http.Request, name string) string {
	// get cookie and check for error
	cookie, err := r.Cookie(name)
	shorts.Check(err)

	// return value
	return cookie.Value
}

// delete cookie with specified name
func deleteCookie(w http.ResponseWriter, name string) {
	// define cookie and delete
	cookie := http.Cookie{Name: name, Value: "null", Path: "/", MaxAge: -1, Expires: expiresCookie}
	http.SetCookie(w, &cookie)
}

// delete cookies with specified names
func deleteCookies(w http.ResponseWriter, names []string) {
	// loop through names
	for _, name := range names {
		// delete cookie
		deleteCookie(w, name)
	}
}

// delete all secpass cookies
func deleteAllCookies(w http.ResponseWriter) {
	// define cookies and delete them
	cookies := []string{"secpass_hash", "secpass_uuid", "secpass_name", "secpass_lang"}
	deleteCookies(w, cookies)
}
