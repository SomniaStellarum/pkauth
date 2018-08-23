package pkauth

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"errors"
)

type User struct {
	Username string `json:"username"`
	Address common.Address `json:"address"`
	Password []byte `json:"password"`
	SessionID string `json:"session_id"`
	MessageHash []byte `json:"messagehash"`
	SignatureHash string `json:"signaturehash"`
}

func (u User) confirmPKSignature(uSignature *User) error {
	b := crypto.VerifySignature(u.Address.Bytes(), u.MessageHash, []byte(uSignature.SignatureHash))
	if !b {
		return errors.New("pkauth: PK Signature not valid")
	}
	return nil
}

