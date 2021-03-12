// Copyright 2021 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package router

import (
	"net/http"

	"github.com/ZupIT/horusec-admin/internal/logger"
	"github.com/ZupIT/horusec-admin/internal/router/api"
	internal "github.com/ZupIT/horusec-admin/internal/router/middleware"
	"github.com/ZupIT/horusec-admin/internal/router/page"
	"github.com/ZupIT/horusec-admin/internal/router/static"
	"github.com/ZupIT/horusec-admin/pkg/core"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/heptiolabs/healthcheck"
	"github.com/thedevsaddam/renderer"
)

type router struct {
	*chi.Mux
	authz  *internal.Authorizer
	render *renderer.Render
	tracer *internal.TraceInitializer

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

	r.Use(middleware.Recoverer)
	r.routeHealthcheckEndpoints()
	r.routeAPIs()
	r.routePages()
	r.serveStaticAssets()
	r.routeErrors()

	return r.Mux, nil
}

func (r *router) routeHealthcheckEndpoints() {
	h := healthcheck.NewHandler()
	r.Handle("/live", http.HandlerFunc(h.LiveEndpoint))
	r.Handle("/ready", http.HandlerFunc(h.ReadyEndpoint))
}

func (r *router) routeAPIs() {
	router := chi.NewRouter()
	router.Use(r.tracer.Initialize)
	if logger.IsTrace() {
		router.Use(middleware.RequestLogger(logger.NewRequestFormatter()))
	}
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
	r.Assets.Serve(r)
}

func (r *router) routeErrors() {
	r.NotFound(func(w http.ResponseWriter, _ *http.Request) {
		err := r.render.HTML(w, http.StatusNotFound, "not-found", nil)
		if err != nil {
			panic(err)
		}
	})
}
