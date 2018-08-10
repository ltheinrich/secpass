package handler

import (
	"io"
	"net/http"
	"os"
	"path"

	"lheinrich.de/secpass/conf"
	"lheinrich.de/secpass/shorts"
)

// JS function
func JS(w http.ResponseWriter, r *http.Request) {
	// set content-type
	w.Header().Set("content-type", "text/javascript; charset=utf-8")

	// open file and check for error
	_, fileName := path.Split(r.URL.Path)
	file, errFile := os.Open(conf.Config["webserver"]["jsDirectory"] + "/" + fileName)
	shorts.Check(errFile, false)
	defer file.Close()

	// write out file
	_, errCopy := io.Copy(w, file)
	shorts.Check(errCopy, false)
}
