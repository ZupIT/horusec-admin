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

	api "github.com/ZupIT/horusec-admin/pkg/api/install/v1alpha1"
	"github.com/ZupIT/horusec-admin/pkg/core"
)

type Manager core.Manager

func (m *Manager) setAPI(api *api.API) {
	if api != nil && api.Ingress != nil {
		u := &url.URL{Scheme: api.Ingress.Scheme, Host: api.Ingress.Host}
		m.APIEndpoint = u.String()
	}
}

func (m *Manager) setAnalytic(analytic *api.Analytic) {
	if analytic != nil && analytic.Ingress != nil {
		u := &url.URL{Scheme: analytic.Ingress.Scheme, Host: analytic.Ingress.Host}
		m.APIEndpoint = u.String()
	}
}

func (m *Manager) setAccount(account *api.Account) {
	if account != nil && account.Ingress != nil {
		u := &url.URL{Scheme: account.Ingress.Scheme, Host: account.Ingress.Host}
		m.APIEndpoint = u.String()
	}
}

func (m *Manager) setAuth(auth *api.Auth) {
	if auth != nil && auth.Ingress != nil {
		u := &url.URL{Scheme: auth.Ingress.Scheme, Host: auth.Ingress.Host}
		m.APIEndpoint = u.String()
	}
}

// nolint:funlen // newManager method needs to set all manager endpoints
func newManager(cr *api.HorusecManager) *core.Manager {
	mng := new(Manager)
	components := cr.Spec.Components
	if components != nil {
		mng.setAPI(components.API)
		mng.setAnalytic(components.Analytic)
		mng.setAccount(components.Account)
		mng.setAuth(components.Auth)
	}
	return (*core.Manager)(mng)
}
