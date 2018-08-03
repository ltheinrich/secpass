package handler

import (
	"net/http"

	"lheinrich.de/extgo/shorts"
)

// Index function
func Index(w http.ResponseWriter, r *http.Request) {
	// execute template
	shorts.Check(tpl.ExecuteTemplate(w, "index.gohtml", nil))
}
