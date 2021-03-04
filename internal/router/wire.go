//+build wireinject

package router

import (
	"github.com/ZupIT/horusec-admin/internal/core"
	"github.com/ZupIT/horusec-admin/internal/kubernetes"
	"github.com/ZupIT/horusec-admin/internal/router/api"
	"github.com/ZupIT/horusec-admin/internal/router/handler"
	"github.com/ZupIT/horusec-admin/internal/router/middleware"
	"github.com/ZupIT/horusec-admin/internal/router/page"
	"github.com/ZupIT/horusec-admin/internal/router/render"
	"github.com/ZupIT/horusec-admin/internal/router/static"
	"github.com/go-chi/chi"
	"github.com/google/wire"
)

var providers = wire.NewSet(
	api.NewSet,
	chi.NewRouter,
	core.NewConfigService,
	handler.NewAuth,
	handler.NewConfigEditing,
	handler.NewConfigReading,
	handler.NewDefaultRender,
	handler.NewHealth,
	kubernetes.NewHorusecManagerClient,
	kubernetes.NewRestConfig,
	middleware.NewAuthorizer,
	page.NewSet,
	render.New,
	static.ListAssets,
	wire.Bind(new(handler.ConfigReader), new(*core.ConfigService)),
	wire.Bind(new(handler.ConfigWriter), new(*core.ConfigService)),
	wire.Struct(new(api.Handlers), "*"),
	wire.Struct(new(router), "*"),
)

func newRouter() (*router, error) {
	wire.Build(providers)
	return nil, nil
}
