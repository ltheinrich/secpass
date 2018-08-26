package handler

import (
	"net/http"

	"lheinrich.de/secpass/spuser"

	"lheinrich.de/secpass/conf"

	"lheinrich.de/secpass/shorts"
)

// Index function
func Index(w http.ResponseWriter, r *http.Request) {
	// check session
	user := checkSession(r)
	if user != "" {
		// define special and pwns
		special := 0
		var pwns []string

		// haveibeenpwned integration
		if r.URL.Path == "/haveibeenpwned" {
			account := r.FormValue("account")
			pwnedList := spuser.PwnedList(account)
			if len(pwnedList) == 0 {
				special = -1
			} else {
				special = -2
				pwns = pwnedList
			}
		}

		// add password
		name, password := r.PostFormValue("name"), r.PostFormValue("password")
		if name != "" && password != "" {
			// insert into db
			_, err := conf.DB.Exec(conf.GetSQL("add_password"), name, password, user)
			shorts.Check(err, true)
		}

		// edit password
		passwordEditID, passwordEditInput := r.PostFormValue("passwordEditIDAfter"), r.PostFormValue("passwordEditInputAfter")
		if passwordEditID != "" && passwordEditInput != "" {
			// update db
			_, err := conf.DB.Exec(conf.GetSQL("edit_password"), passwordEditInput, passwordEditID, user)
			shorts.Check(err, true)
		}

		// delete password
		passwordDeleteInput := r.PostFormValue("passwordDeleteInput")
		if passwordDeleteInput != "" {
			// delete from db
			_, err := conf.DB.Exec(conf.GetSQL("delete_password"), passwordDeleteInput, user)
			shorts.Check(err, true)
		}

		// execute template
		shorts.Check(tpl.ExecuteTemplate(w, "index.html", Data{User: user, Lang: getLang(r), Passwords: passwords(user), Special: special, Pwns: pwns}), false)
	}

	// redirect to login
	redirect(w, "/login")
}

// return passwords as map
func passwords(user string) map[string]string {
	// query db and check for error
	rows, errQuery := conf.DB.Query(conf.GetSQL("passwords"), user)
	shorts.Check(errQuery, true)

	passwordMap := map[string]string{}

	// loop through rows
	for rows.Next() {
		// define variables to write into
		var name string
		var password string

		// read from rows
		errScan := rows.Scan(&name, &password)
		shorts.Check(errScan, true)

		// put into map
		passwordMap[name] = password
	}

	return passwordMap
}
