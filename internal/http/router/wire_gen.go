// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

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

// Injectors from wire.go:

func newRouter() (*router, error) {
	mux := chi.NewRouter()
	authorizer := middleware.NewAuthorizer()
	rendererRender := render.New()
	auth := handler.NewAuth()
	config, err := kubernetes.NewRestConfig()
	if err != nil {
		return nil, err
	}
	horusecManagerInterface, err := kubernetes.NewHorusecManagerClient(config)
	if err != nil {
		return nil, err
	}
	configService := core.NewConfigService(horusecManagerInterface)
	configEditing := handler.NewConfigEditing(rendererRender, configService)
	configReading := handler.NewConfigReading(rendererRender, configService)
	health := handler.NewHealth()
	routerApiHandlers := &apiHandlers{
		Auth:          auth,
		ConfigEditing: configEditing,
		ConfigReading: configReading,
		Health:        health,
	}
	v := newAPIs(routerApiHandlers)
	v2, err := scanAssets()
	if err != nil {
		return nil, err
	}
	defaultRender := handler.NewDefaultRender(rendererRender)
	v3 := newPages(defaultRender)
	routerRouter := &router{
		Mux:    mux,
		authz:  authorizer,
		render: rendererRender,
		APIs:   v,
		Assets: v2,
		Pages:  v3,
	}
	return routerRouter, nil
}

// wire.go:

var providers = wire.NewSet(chi.NewRouter, core.NewConfigService, handler.NewAuth, handler.NewConfigEditing, handler.NewConfigReading, handler.NewDefaultRender, handler.NewHealth, kubernetes.NewHorusecManagerClient, kubernetes.NewRestConfig, middleware.NewAuthorizer, render.New, newAPIs,
	newPages,
	scanAssets, wire.Bind(new(handler.ConfigReader), new(*core.ConfigService)), wire.Bind(new(handler.ConfigWriter), new(*core.ConfigService)), wire.Struct(new(apiHandlers), "*"), wire.Struct(new(router), "*"),
)
