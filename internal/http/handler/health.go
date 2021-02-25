package handler

import "net/http"

type Health struct{}

func NewHealth() *Health {
	return new(Health)
}

func (h *Health) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNoContent)
}
