package handler

import "net/http"

type Auth struct{}

func NewAuth() *Auth {
	return new(Auth)
}

func (h *Auth) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}
