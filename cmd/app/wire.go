//+build wireinject

package main

import (
	"github.com/ZupIT/horusec-admin/internal/business"
	"github.com/ZupIT/horusec-admin/internal/kubernetes"
	"github.com/ZupIT/horusec-admin/internal/router"
	"github.com/ZupIT/horusec-admin/pkg/core"
	"github.com/go-chi/chi"
	"github.com/google/wire"
)

var providers = wire.NewSet(
	business.NewConfigService,
	kubernetes.NewHorusecManagerClient,
	kubernetes.NewRestConfig,
	router.New,
	wire.Bind(new(core.ConfigurationReader), new(*business.ConfigService)),
	wire.Bind(new(core.ConfigurationWriter), new(*business.ConfigService)),
)

func newRouter() (*chi.Mux, error) {
	wire.Build(providers)
	return nil, nil
}
