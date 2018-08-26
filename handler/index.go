package handler

import (
	"net/http"

	"lheinrich.de/secpass/spuser"

	"lheinrich.de/secpass/conf"

	"lheinrich.de/secpass/shorts"
)

// Password structure
type Password struct {
	Title    string
	Name     string
	Password string
}

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
		title, name, password := r.PostFormValue("title"), r.PostFormValue("name"), r.PostFormValue("password")
		if name != "" && password != "" {
			// insert into db
			_, err := conf.DB.Exec(conf.GetSQL("add_password"), title, name, password, user)
			shorts.Check(err)
		}

		// edit password
		passwordEditTitle, passwordEditID, passwordEditInput := r.PostFormValue("passwordEditTitleAfter"), r.PostFormValue("passwordEditIDAfter"), r.PostFormValue("passwordEditInputAfter")
		if passwordEditID != "" && passwordEditInput != "" {
			// update db
			_, err := conf.DB.Exec(conf.GetSQL("edit_password"), passwordEditInput, passwordEditTitle, passwordEditID, user)
			shorts.Check(err)
		}

		// delete password
		passwordDeleteTitle, passwordDeleteInput := r.PostFormValue("passwordDeleteTitle"), r.PostFormValue("passwordDeleteInput")
		if passwordDeleteTitle != "" && passwordDeleteInput != "" {
			// delete from db
			_, err := conf.DB.Exec(conf.GetSQL("delete_password"), passwordDeleteTitle, passwordDeleteInput, user)
			shorts.Check(err)
		}

		// execute template
		shorts.Check(tpl.ExecuteTemplate(w, "index.html", Data{User: user, Lang: getLang(r), Passwords: passwords(user), Special: special, Pwns: pwns}))
	}

	// redirect to login
	redirect(w, "/login")
}

// return passwords as map
func passwords(user string) []Password {
	// query db and check for error
	rows, errQuery := conf.DB.Query(conf.GetSQL("passwords"), user)
	shorts.Check(errQuery)

	passwordList := []Password{}

	// loop through rows
	for rows.Next() {
		// define variables to write into
		var title string
		var name string
		var password string

		// read from rows
		errScan := rows.Scan(&title, &name, &password)
		shorts.Check(errScan)

		// put into list
		passwordList = append(passwordList, Password{Title: title, Name: name, Password: password})
	}

	return passwordList
}
