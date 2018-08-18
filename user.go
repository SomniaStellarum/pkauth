package pkauth

import (
	"github.com/ethereum/go-ethereum/common"
	"time"
)

type User struct {
	Username string `json:"username"`
	Address common.Address `json:"address"`
	Password []byte `json:"password"`
	SessionID string `json:"session_id"`
	SessionTimeout time.Time `json:"sessiontimeout"`
	SessionVerified bool `json:"sessionverified"`
}

func (u User) confirmPassword() (bool, error) {


	return true, nil
}

