package handler

import (
	"html/template"
	"net/http"
	"strings"

	"lheinrich.de/secpass/conf"
	"lheinrich.de/secpass/shorts"
)

var (
	// define functions
	funcs = template.FuncMap{
		"config": func(key string) string {
			// split key string by . to get group.key and return value
			keys := strings.Split(key, ".")
			return conf.Config[keys[0]][keys[1]]
		},
		"lang": func(key string) string {
			// split key string by . to get language.key and return value
			keys := strings.Split(key, ".")
			return (*conf.Lang[keys[0]])[keys[1]]
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

func redirect(w http.ResponseWriter, location string) {
	w.Header().Set("location", location)
	w.WriteHeader(307)
}
