package router

import (
	"net/http"

	"github.com/ZupIT/horusec-admin/internal/router/page"
	"github.com/ZupIT/horusec-admin/internal/router/static"
	"github.com/ZupIT/horusec-admin/pkg/core"

	"github.com/ZupIT/horusec-admin/internal/router/api"

	"github.com/ZupIT/horusec-admin/internal/logger"
	internal "github.com/ZupIT/horusec-admin/internal/router/middleware"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/thedevsaddam/renderer"
)

type router struct {
	*chi.Mux
	authz  *internal.Authorizer
	render *renderer.Render

	APIs   api.Set
	Assets static.Assets
	Pages  page.Set
}

// New creates the router with all API routes and the static files handler.
func New(reader core.ConfigurationReader, writer core.ConfigurationWriter) (*chi.Mux, error) {
	r, err := newRouter(reader, writer)
	if err != nil {
		return nil, err
	}
	if logger.IsTrace() {
		r.Use(middleware.RequestLogger(logger.NewRequestFormatter()))
	}
	r.Use(middleware.Recoverer)
	r.routeAPIs()
	r.routePages()
	r.serveStaticAssets()
	r.routeErrors()

	return r.Mux, nil
}

func (r *router) routeAPIs() {
	router := chi.NewRouter()
	for _, route := range r.APIs {
		handlerFunc := route.Handler
		if route.Authenticated {
			handlerFunc = r.authz.Authorize(handlerFunc)
		}
		router.Method(route.Method, route.Pattern, handlerFunc)
	}
	r.Mount("/api", router)
}

func (r *router) routePages() {
	router := chi.NewRouter()
	for _, route := range r.Pages {
		router.Method(http.MethodGet, route.Pattern, route.Handler)
	}
	r.Mount("/view", router)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/view", http.StatusMovedPermanently)
	})
}

func (r *router) serveStaticAssets() {
	for _, a := range r.Assets {
		a.Serve(r.Mux)
	}
}

func (r *router) routeErrors() {
	r.NotFound(func(w http.ResponseWriter, _ *http.Request) {
		err := r.render.HTML(w, http.StatusNotFound, "not-found", nil)
		if err != nil {
			panic(err)
		}
	})
}
