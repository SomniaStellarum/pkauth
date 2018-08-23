package pkauth

import (
	"net/http"
	"github.com/gin-gonic/gin/json"
)

type SingleEthAuthorizer struct {
	Storage DB
	SigninURL string
	h http.Handler
}

func (s SingleEthAuthorizer) Authorize(hd http.Handler) http.Handler {
	s2 := SingleEthAuthorizer{
		Storage: s.Storage,
		SigninURL: s.SigninURL,
		h: hd,
	}
	return s2
}

func (s SingleEthAuthorizer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Checks:
	// check for cookie, if session token is valid and verified, continue to page
	// For unverified
	// check for signed message: fail redirect to signin
	// check it's valid then: fail redirect to signin
	// check for session token in cookies: fail create cookie
	c, err := r.Cookie("session_id")
	if err == http.ErrNoCookie {
		//redirect to signin
		http.Redirect(w, r, s.SigninURL, 302)
	}
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	// Get user data, get if session verified and not past timeout
	err = VerifySession(s.Storage, c.String())
	if err != nil {
		err = s.verifyPKSignature(r, c)
		if err != nil {
			http.Redirect(w, r, s.SigninURL, 302)
		}
	}
	s.h.ServeHTTP(w, r)
}

func (s SingleEthAuthorizer) verifyPKSignature(r *http.Request, c *http.Cookie) error {
	session, err := GetSession(s.Storage, c.String())
	if err != nil {
		return err
	}
	user, err := GetUser(s.Storage, session.Username)
	if err != nil {
		return err
	}
	var uSignature *User
	d := json.NewDecoder(r.Body)
	err = d.Decode(uSignature)
	if err != nil {
		return err
	}
	err = user.confirmPKSignature(uSignature)
	if err != nil {
		return err
	}
	return nil
}