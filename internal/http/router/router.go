package router

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	internal "github.com/tiagoangelozup/horusec-admin/internal/http/middleware"
	"net/http"
)

type router struct {
	*chi.Mux
	authz  *internal.Authorizer
	Pages  []*Page
	Routes []*Route
}

// New creates the router with all API routes and the static files handler.
func New() (*chi.Mux, error) {
	r, err := newRouter()
	if err != nil {
		return nil, err
	}
	r.Use(middleware.Logger)

	// routing apis
	api := chi.NewRouter()
	for _, route := range r.Routes {
		handlerFunc := route.Handler
		if route.Authenticated {
			handlerFunc = r.authz.Authorize(handlerFunc)
		}
		api.Method(route.Method, route.Pattern, handlerFunc)
	}
	r.Mount("/api", api)

	// routing views
	view := chi.NewRouter()
	for _, pg := range r.Pages {
		view.Get(pg.Pattern, pg.Render)
	}
	r.Mount("/view", view)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/view", http.StatusMovedPermanently)
	})

	return r.Mux, nil
}
