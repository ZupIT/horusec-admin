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

package router

import (
	"github.com/ZupIT/horusec-admin/internal/authz"
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
	middleware.NewTracer,
	authz.New,
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
