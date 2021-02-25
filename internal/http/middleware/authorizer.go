package middleware

import "net/http"

type Authorizer struct{}

func NewAuthorizer() *Authorizer {
	return new(Authorizer)
}

func (a *Authorizer) Authorize(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// TODO: handle access token
		next.ServeHTTP(w, r)
	})
}
