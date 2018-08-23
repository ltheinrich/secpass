package user

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"lheinrich.de/secpass/shorts"
)

// Sessions map
var Sessions = map[string]Session{}

// Session structure
type Session struct {
	User    string
	Expires time.Time
}

// IsValid check session is valid
func (s Session) IsValid() bool {
	return s.Expires.Before(time.Now())
}

// CleanupSessions delete sessions no longer valid
func CleanupSessions() {
	for uuid, session := range Sessions {
		if !session.IsValid() {
			delete(Sessions, uuid)
		}
	}
}

// PwnedList titles from https://haveibeenpwned.com/api/v2/breachedaccount/<NAME>
func PwnedList(name string) []string {
	// http request and check error
	resp, err := http.DefaultClient.Get("https://haveibeenpwned.com/api/v2/breachedaccount/" + name)
	shorts.Check(err, false)

	// read body and check error
	read, errRead := ioutil.ReadAll(resp.Body)
	shorts.Check(errRead, false)

	// struct for json
	type titleStruct struct {
		Title string
	}

	// unmarshal json
	titleList := []titleStruct{}
	json.Unmarshal(read, &titleList)

	// append to slice from struct
	titles := []string{}
	for _, title := range titleList {
		titles = append(titles, title.Title)
	}

	// return
	return titles
}
