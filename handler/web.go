package handler

import (
	"io"
	"mime"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"lheinrich.de/secpass/conf"
	"lheinrich.de/secpass/shorts"
)

// Web function
func Web(w http.ResponseWriter, r *http.Request) {
	// open file and check for error
	_, fileName := path.Split(r.URL.Path)
	file, errFile := os.Open(conf.Config["webserver"]["webDirectory"] + "/" + fileName)
	shorts.Check(errFile)
	defer file.Close()

	// set content-type
	w.Header().Set("content-type", mime.TypeByExtension(filepath.Ext(fileName))+"; charset=utf-8")

	// write out file
	_, errCopy := io.Copy(w, file)
	shorts.Check(errCopy)
}
