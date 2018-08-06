package pkauth

import "net/http"

type SingleEthAuthorizer struct {
	Storage DB
	RedirectURL string
	h http.Handler
}

func (s SingleEthAuthorizer) Authorize(h http.Handler) http.Handler {
	s2 := SingleEthAuthorizer{}(s.Storage, s.RedirectURL, h)
	return s2
}

func (s SingleEthAuthorizer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Checks:
	// check for signed message: fail redirect to signin
	// check it's valid then: fail redirect to signin
	// check for session token in cookies: fail create cookie
}
