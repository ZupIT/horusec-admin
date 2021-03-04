// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/ZupIT/horusec-admin/internal/business"
	"github.com/ZupIT/horusec-admin/internal/kubernetes"
	"github.com/ZupIT/horusec-admin/internal/router"
	"github.com/ZupIT/horusec-admin/pkg/core"
	"github.com/go-chi/chi"
	"github.com/google/wire"
)

// Injectors from wire.go:

func newRouter() (*chi.Mux, error) {
	config, err := kubernetes.NewRestConfig()
	if err != nil {
		return nil, err
	}
	horusecManagerInterface, err := kubernetes.NewHorusecManagerClient(config)
	if err != nil {
		return nil, err
	}
	configService := business.NewConfigService(horusecManagerInterface)
	mux, err := router.New(configService, configService)
	if err != nil {
		return nil, err
	}
	return mux, nil
}

// wire.go:

var providers = wire.NewSet(business.NewConfigService, kubernetes.NewHorusecManagerClient, kubernetes.NewRestConfig, router.New, wire.Bind(new(core.ConfigurationReader), new(*business.ConfigService)), wire.Bind(new(core.ConfigurationWriter), new(*business.ConfigService)))