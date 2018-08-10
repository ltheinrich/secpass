package handler

import (
	"io"
	"net/http"
	"os"
	"path"
	"strings"

	"lheinrich.de/secpass/conf"
	"lheinrich.de/secpass/shorts"
)

// Image function
func Image(w http.ResponseWriter, r *http.Request) {
	// set content-type
	extension := strings.Split(r.URL.Path, ".")
	w.Header().Set("content-type", "image/"+extension[len(extension)-1]+"; charset=utf-8")

	// open file and check for error
	_, fileName := path.Split(r.URL.Path)
	file, errFile := os.Open(conf.Config["webserver"]["imageDirectory"] + "/" + fileName)
	shorts.Check(errFile, false)
	defer file.Close()

	// write out file
	_, errCopy := io.Copy(w, file)
	shorts.Check(errCopy, false)
}
