package handler

import (
	"io"
	"mime"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"

	"lheinrich.de/lheinrich/secpass/internal/pkg/conf"
	"lheinrich.de/lheinrich/secpass/internal/pkg/shorts"
)

// Web function
func Web(w http.ResponseWriter, r *http.Request) {
	// define file name and open file
	directory, fileName := path.Split(r.URL.Path)
	file, errFile := os.Open(conf.Config["webserver"]["webDirectory"] + getWebfileDirectory(directory) + fileName)

	// check for error and defer close
	shorts.Check(errFile)
	defer file.Close()

	// set content-type
	w.Header().Set("content-type", mime.TypeByExtension(filepath.Ext(fileName))+"; charset=utf-8")

	// write out file
	_, errCopy := io.Copy(w, file)
	shorts.Check(errCopy)
}

// detect css, js oder images directory
func getWebfileDirectory(directory string) string {
	if strings.HasPrefix(directory, "/web/css/") {
		// css file
		return "/css/"
	} else if strings.HasPrefix(directory, "/web/js/") {
		// js file
		return "/js/"
	} else if strings.HasPrefix(directory, "/web/images/") {
		// image file
		return "/images/"
	}

	// just a file
	return "/"
}
