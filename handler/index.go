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
	user := checkSession(w, r)
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

		// execute template
		shorts.Check(tpl.ExecuteTemplate(w, "index.html", Data{User: user, Lang: getLang(r), Passwords: getPasswords(user), Special: special, Pwns: pwns}))
	}

	// redirect to login
	redirect(w, "/login")
}

// return getPasswords as map
func getPasswords(user string) []Password {
	// query db and check for error
	rows, errQuery := conf.DB.Query(conf.GetSQL("passwords"), user)
	shorts.Check(errQuery)

	passwordList := []Password{}

	// loop through rows
	for rows.Next() {
		// define variables to write into
		var id int
		var title string
		var name string
		var mail string
		var password string
		var url string
		var backupCode string
		var notes string

		// read from rows
		errScan := rows.Scan(&id, &title, &name, &mail, &password, &url, &backupCode, &notes)
		shorts.Check(errScan)

		// put into list
		passwordList = append(passwordList, Password{ID: id, Title: title, Name: name, Mail: mail, Value: password, URL: url, BackupCode: backupCode, Notes: notes})
	}

	return passwordList
}
