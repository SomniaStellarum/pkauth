package pkauth

import (
	"encoding/json"
	"time"
	"errors"
)

type Putter interface {
	Put(key []byte, value []byte) error
}

type Getter interface {
	Get(key []byte) ([]byte, error)
}

type Deleter interface {
	Delete(key []byte) error
}

type DB interface {
	Putter
	Getter
	Deleter
}

func VerifySession(d Getter, session_id string) error {
	v, err := d.Get([]byte(session_id))
	if err != nil {
		return err
	}
	var s *Session
	err = json.Unmarshal(v, s)
	if err != nil {
		return err
	}
	if (!s.SessionVerified)|(s.SessionTimeout>time.Now()) {
		return errors.New("pkauth: Unauthorized Session")
	}
	return nil
}

func GetSession(d Getter, session_id string) (*Session, error) {
	v, err := d.Get([]byte(session_id))
	if err != nil {
		return err
	}
	var s *Session
	err = json.Unmarshal(v, s)
	if err != nil {
		return err
	}
	return s, nil
}

func GetUser(d Getter, username string) (*User, error) {
	v, err := d.Get([]byte(username))
	if err != nil {
		return err
	}
	var u *User
	err = json.Unmarshal(v, u)
	if err != nil {
		return err
	}
	return u, nil
}