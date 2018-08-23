package pkauth

import "time"

type Session struct {
	Username string `json:"username"`
	SessionID string `json:"session_id"`
	SessionTimeout time.Time `json:"sessiontimeout"`
	SessionVerified bool `json:"sessionverified"`
}
