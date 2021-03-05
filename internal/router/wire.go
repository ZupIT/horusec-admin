//+build wireinject

package router

import (
	"github.com/ZupIT/horusec-admin/internal/router/api"
	"github.com/ZupIT/horusec-admin/internal/router/handler"
	"github.com/ZupIT/horusec-admin/internal/router/middleware"
	"github.com/ZupIT/horusec-admin/internal/router/page"
	"github.com/ZupIT/horusec-admin/internal/router/render"
	"github.com/ZupIT/horusec-admin/internal/router/static"
	"github.com/ZupIT/horusec-admin/pkg/core"
	"github.com/go-chi/chi"
	"github.com/google/wire"
)

var providers = wire.NewSet(
	api.NewSet,
	chi.NewRouter,
	handler.NewAuth,
	handler.NewConfigEditing,
	handler.NewConfigReading,
	handler.NewDefaultRender,
	handler.NewHealth,
	middleware.NewAuthorizer,
	page.NewSet,
	render.New,
	static.ListAssets,
	wire.Struct(new(api.Handlers), "*"),
	wire.Struct(new(router), "*"),
)

func newRouter(reader core.ConfigurationReader, writer core.ConfigurationWriter) (*router, error) {
	wire.Build(providers)
	return nil, nil
}
