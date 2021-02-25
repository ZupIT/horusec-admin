//+build wireinject

package router

import (
	"github.com/go-chi/chi"
	"github.com/google/wire"
	"github.com/tiagoangelozup/horusec-admin/internal/http/handler"
	"github.com/tiagoangelozup/horusec-admin/internal/http/middleware"
	"github.com/tiagoangelozup/horusec-admin/internal/http/render"
)

var providers = wire.NewSet(
	chi.NewRouter,
	handler.NewAuth,
	handler.NewConfigEditing,
	handler.NewConfigReading,
	handler.NewHealth,
	middleware.NewAuthorizer,
	newRoutes,
	render.New,
	wire.Struct(new(handlers), "*"),
	wire.Struct(new(router), "*"),
)

func newRouter() *router {
	wire.Build(providers)
	return nil
}
