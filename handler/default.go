package handler

import (
	"html/template"
	"net/http"
	"strings"

	"lheinrich.de/extgo/shorts"

	"lheinrich.de/secpass/conf"
)

var (
	// define functions
	funcs = template.FuncMap{"config": func(key string) string {
		// split key string by . to get group.key
		keys := strings.Split(key, ".")
		return conf.Config[keys[0]][keys[1]]
	}}

	// load templates
	tpl, _ = template.New("").Funcs(funcs).ParseGlob("templates/*/*")
)

// name function
func name(w http.ResponseWriter, r *http.Request) {
	// execute template
	shorts.Check(tpl.ExecuteTemplate(w, "name", nil))
}
