package handler

import (
	"net/http"
	"strconv"

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

		// define form values
		categoryID, _ := strconv.Atoi(r.PostFormValue("categoryID"))
		categoryNew, categoryDelete, categoryName := r.PostFormValue("categoryNew"), r.PostFormValue("categoryDelete"), r.PostFormValue("categoryName")

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

		// create category
		if categoryNew != "" {
			_, err := conf.DB.Exec(conf.GetSQL("add_category"), categoryNew, user)
			shorts.Check(err)
		}

		// edit category
		if categoryName != "" {
			_, err := conf.DB.Exec(conf.GetSQL("edit_category"), categoryName, categoryID, user)
			shorts.Check(err)
		}

		// delete category
		if categoryDelete != "" {
			_, errPasswords := conf.DB.Exec(conf.GetSQL("delete_category_passwords"), categoryDelete, user)
			_, errCategory := conf.DB.Exec(conf.GetSQL("delete_category"), categoryDelete, user)
			shorts.Check(errPasswords)
			shorts.Check(errCategory)
		}

		// execute template
		shorts.Check(tpl.ExecuteTemplate(w, "index.html", Data{User: user, Lang: getLang(r), Passwords: getPasswords(user),
			Categories: getCategories(user), DefaultCategory: getDefaultCategory(user), Special: special, Pwns: pwns}))
	}

	// redirect to login
	redirect(w, "/login")
}

// return passwords as list
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
		var category int

		// read from rows
		errScan := rows.Scan(&id, &title, &name, &mail, &password, &url, &backupCode, &notes, &category)
		shorts.Check(errScan)

		// put into list
		passwordList = append(passwordList, Password{ID: id, Title: title, Name: name, Mail: mail, Value: password, URL: url, BackupCode: backupCode, Notes: notes, Category: category})
	}

	return passwordList
}

// return categories as list
func getCategories(user string) []Category {
	// query db and check for error
	rows, errQuery := conf.DB.Query(conf.GetSQL("get_categories"), user)
	shorts.Check(errQuery)

	categoryList := []Category{}

	// loop through rows
	for rows.Next() {
		// define variables to write into
		var id int
		var name string

		// read fron rows
		errScan := rows.Scan(&id, &name)
		shorts.Check(errScan)

		// put into list
		categoryList = append(categoryList, Category{ID: id, Name: name})
	}

	return categoryList
}

// return Category
func getCategory(id int, user string) Category {
	// define variable to write into
	var name string

	// query db and read
	err := conf.DB.QueryRow(conf.GetSQL("get_category"), id, user).Scan(&name)
	shorts.Check(err)

	return Category{ID: id, Name: name}
}
