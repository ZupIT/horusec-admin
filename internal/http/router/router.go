package router

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/thedevsaddam/renderer"
	internal "github.com/tiagoangelozup/horusec-admin/internal/http/middleware"
)

type router struct {
	*chi.Mux
	authz  *internal.Authorizer
	render *renderer.Render

	APIs   []*API
	Assets []*LocalFileSystem
	Pages  []*Page
}

// New creates the router with all API routes and the static files handler.
func New() (*chi.Mux, error) {
	r, err := newRouter()
	if err != nil {
		return nil, err
	}
	r.Use(middleware.Logger)
	r.routeAPIs()
	r.routePages()
	r.servesStaticFiles()
	r.routeErrors()

	return r.Mux, nil
}

func (r *router) routeAPIs() {
	api := chi.NewRouter()
	for _, route := range r.APIs {
		handlerFunc := route.Handler
		if route.Authenticated {
			handlerFunc = r.authz.Authorize(handlerFunc)
		}
		api.Method(route.Method, route.Pattern, handlerFunc)
	}
	r.Mount("/api", api)
}

func (r *router) routePages() {
	view := chi.NewRouter()
	for _, route := range r.Pages {
		view.Method(http.MethodGet, route.Pattern, route.Handler)
	}
	r.Mount("/view", view)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/view", http.StatusMovedPermanently)
	})
}

func (r *router) servesStaticFiles() {
	for _, a := range r.Assets {
		a.serve(r.Mux)
	}
}

func (r *router) routeErrors() {
	r.NotFound(func(w http.ResponseWriter, _ *http.Request) {
		err := r.render.HTML(w, http.StatusNotFound, "not-found", nil)
		if err != nil {
			log.Fatal(err)
		}
	})
}
