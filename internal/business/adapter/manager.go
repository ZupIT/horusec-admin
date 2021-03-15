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

package adapter

import (
	"net/url"

	"github.com/ZupIT/horusec-admin/pkg/core"

	api "github.com/ZupIT/horusec-admin/pkg/api/install/v1alpha1"
)

// nolint:funlen // newManager method needs to set all manager endpoints
func newManager(cr *api.HorusecManager) *core.Manager {
	mng := new(core.Manager)
	components := cr.Spec.Components
	if components != nil {
		mng.APIEndpoint = (&url.URL{
			Scheme: components.API.Ingress.Scheme, Host: components.API.Ingress.Host,
		}).String()
		mng.AnalyticEndpoint = (&url.URL{
			Scheme: components.Analytic.Ingress.Scheme, Host: components.Analytic.Ingress.Host,
		}).String()
		mng.AccountEndpoint = (&url.URL{
			Scheme: components.Account.Ingress.Scheme, Host: components.Account.Ingress.Host,
		}).String()
		mng.AuthEndpoint = (&url.URL{
			Scheme: components.Auth.Ingress.Scheme, Host: components.Auth.Ingress.Host,
		}).String()
	}
	return mng
}
