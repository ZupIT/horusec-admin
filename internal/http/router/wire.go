//+build wireinject

package router

import (
	"github.com/ZupIT/horusec-admin/internal/core"
	"github.com/ZupIT/horusec-admin/internal/http/handler"
	"github.com/ZupIT/horusec-admin/internal/http/middleware"
	"github.com/ZupIT/horusec-admin/internal/http/render"
	"github.com/ZupIT/horusec-admin/internal/kubernetes"
	"github.com/go-chi/chi"
	"github.com/google/wire"
)

var providers = wire.NewSet(
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
	render.New,
	newAPIs,
	newPages,
	scanAssets,
	wire.Bind(new(handler.ConfigReader), new(*core.ConfigService)),
	wire.Bind(new(handler.ConfigWriter), new(*core.ConfigService)),
	wire.Struct(new(apiHandlers), "*"),
	wire.Struct(new(router), "*"),
)

func newRouter(handler.ConfigReader, handler.ConfigWriter) (*router, error) {
	wire.Build(providers)
	return nil, nil
}
