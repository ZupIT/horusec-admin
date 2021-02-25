package router

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	internal "github.com/tiagoangelozup/horusec-admin/internal/http/middleware"
)

type router struct {
	*chi.Mux
	authz  *internal.Authorizer
	Routes []*Route
}

// New creates the router with all API routes and the static files handler.
func New() *chi.Mux {
	r := newRouter()
	r.Use(middleware.Logger)

	api := chi.NewRouter()
	for _, route := range r.Routes {
		handlerFunc := route.Handler
		if route.Authenticated {
			handlerFunc = r.authz.Authorize(handlerFunc)
		}
		api.Method(route.Method, route.Pattern, handlerFunc)
	}
	r.Mount("/api", api)

	return r.Mux
}
