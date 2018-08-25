package handler

import (
	"net/http"

	"lheinrich.de/secpass/shorts"
)

// Index function
func Index(w http.ResponseWriter, r *http.Request) {
	// check session
	user := checkSession(r)
	if user != "" {
		// define passwords (test)
		passwords := map[string]string{}
		passwords["Alvin"] = "{\"iv\":\"taJ7ZsRkO9bij+lLr377zA==\",\"v\":1,\"iter\":10000,\"ks\":256,\"ts\":64,\"mode\":\"ccm\",\"adata\":\"\",\"cipher\":\"aes\",\"salt\":\"vzH3IITZG4Q=\",\"ct\":\"tuvHwCBfvXbLV/VypLwWN5TWYHU5\"}"

		// execute template
		shorts.Check(tpl.ExecuteTemplate(w, "index.html", Data{User: user, Lang: getLang(r), Passwords: passwords}), false)
	}

	// redirect to login
	redirect(w, "/login")
}
