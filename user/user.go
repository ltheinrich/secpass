package user

import "time"

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
