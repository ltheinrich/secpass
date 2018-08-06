package user

import "time"

// Users map
var Users = map[int]User{}

// Sessions map
var Sessions = map[string]Session{}

// User structure
type User struct {
	ID   int
	Name string
	Lang string
}

// Session structure
type Session struct {
	UUID   string
	UserID int
	Valid  time.Time
}

// IsValid check session is valid
func (s Session) IsValid() bool {
	return s.Valid.Before(time.Now())
}

// CleanupSessions delete sessions no longer valid
func CleanupSessions() {
	for uuid, session := range Sessions {
		if !session.IsValid() {
			delete(Sessions, uuid)
		}
	}
}
