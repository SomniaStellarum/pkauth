package pkauth

import "net/http"

type Authorizer interface {
	Authorize(h http.Handler) http.Handler
}
