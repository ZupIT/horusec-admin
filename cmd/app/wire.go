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
