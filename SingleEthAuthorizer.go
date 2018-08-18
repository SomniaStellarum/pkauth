package pkauth

import "net/http"

type SingleEthAuthorizer struct {
	Storage DB
	SigninURL string
	PKSignURL string
	h http.Handler
}

func (s SingleEthAuthorizer) Authorize(h http.Handler) http.Handler {
	s2 := SingleEthAuthorizer{}(s.Storage, s.SigninURL, h)
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
	err = s.sessionValid(c.String())
	if err != nil {
		http.Redirect(w, r, s.SigninURL, 302)
	}
	s.h.ServeHTTP(w, r)
}


func (s SingleEthAuthorizer) sessionValid(session_id string) error {

	return nil
}