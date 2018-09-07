package handler

import (
	"database/sql"
	"net/http"
	"strconv"

	"lheinrich.de/secpass/conf"

	"lheinrich.de/secpass/shorts"
)

// Password structure
type Password struct {
	ID         int
	Title      string
	Name       string
	Mail       string
	Value      string
	URL        string
	BackupCode string
	Notes      string
}

// Entry function
func Entry(w http.ResponseWriter, r *http.Request) {
	// check session
	user := checkSession(w, r)
	if user != "" {
		// define special and form values
		special := 0
		id, _ := strconv.Atoi(r.FormValue("id"))
		title, name, mail, password := r.PostFormValue("title"), r.PostFormValue("name"), r.PostFormValue("mail"), r.PostFormValue("password")
		url, backupCode, notes := r.PostFormValue("url"), r.PostFormValue("backupCode"), r.PostFormValue("notes")

		// sql query trash
		var trash string

		// add password
		if id == 0 && title != "" && (name != "" || mail != "") && len(password) >= 4 {
			// check if already exists
			errQuery := conf.DB.QueryRow(conf.GetSQL("password"), id, user).Scan(&trash, &trash, &trash, &trash, &trash, &trash, &trash)

			if errQuery == sql.ErrNoRows {
				// insert into db
				_, errExec := conf.DB.Exec(conf.GetSQL("add_password"), title, name, mail, password, url, backupCode, notes, user)
				shorts.Check(errExec)
				redirect(w, "/")
				return
			} else if errQuery == nil {
				// entry already exists
				special = -1
			} else {
				// check error
				shorts.Check(errQuery)
			}
		}

		// edit password
		if id != 0 && title != "" && (name != "" || mail != "") && len(password) >= 4 {
			// check if exists
			errQuery := conf.DB.QueryRow(conf.GetSQL("password"), id, user).Scan(&trash, &trash, &trash, &trash, &trash, &trash, &trash)

			if errQuery == nil {
				// update db
				_, errExec := conf.DB.Exec(conf.GetSQL("edit_password"), title, name, mail, password, url, backupCode, notes, id, user)
				shorts.Check(errExec)
			} else if errQuery == sql.ErrNoRows {
				// entry not exists
				special = -2
			} else {
				// check error
				shorts.Check(errQuery)
			}
		}

		// delete password
		delete := r.PostFormValue("delete")
		if id != 0 && delete == "delete" {
			// delete from db
			_, err := conf.DB.Exec(conf.GetSQL("delete_password"), id, user)
			shorts.Check(err)
			redirect(w, "/")
		}

		// execute template
		shorts.Check(tpl.ExecuteTemplate(w, "entry.html", Data{User: user, Lang: getLang(r), Entry: getPassword(id, user), Special: special}))
	}

	// redirect to login
	redirect(w, "/login")
}

// return password
func getPassword(id int, user string) Password {
	// query db and check for error
	rows, errQuery := conf.DB.Query(conf.GetSQL("password"), id, user)
	shorts.Check(errQuery)

	// return if exist
	if rows.Next() {
		// define variables to write into
		var title string
		var name string
		var mail string
		var password string
		var url string
		var backupCode string
		var notes string

		// read from row
		errScan := rows.Scan(&title, &name, &mail, &password, &url, &backupCode, &notes)
		shorts.Check(errScan)

		// return
		return Password{ID: id, Title: title, Name: name, Mail: mail, Value: password, URL: url, BackupCode: backupCode, Notes: notes}
	}

	// return empty password
	return Password{}
}
