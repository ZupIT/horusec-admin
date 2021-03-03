//+build wireinject

package router

import (
	"github.com/ZupIT/horusec-admin/internal/http/handler"
	"github.com/ZupIT/horusec-admin/internal/http/middleware"
	"github.com/ZupIT/horusec-admin/internal/http/render"
	"github.com/go-chi/chi"
	"github.com/google/wire"
)

var providers = wire.NewSet(
	chi.NewRouter,
	handler.NewAuth,
	handler.NewConfigEditing,
	handler.NewConfigReading,
	handler.NewDefaultRender,
	handler.NewHealth,
	middleware.NewAuthorizer,
	render.New,
	newAPIs,
	newPages,
	scanAssets,
	wire.Struct(new(apiHandlers), "*"),
	wire.Struct(new(router), "*"),
)

func newRouter() (*router, error) {
	wire.Build(providers)
	return nil, nil
}
