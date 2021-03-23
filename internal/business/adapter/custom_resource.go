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
	api "github.com/ZupIT/horusec-admin/pkg/api/install/v1alpha1"
	"github.com/ZupIT/horusec-admin/pkg/core"
	"net/url"
)

type CustomResource api.HorusecManager

func ForCustomResource(hm *api.HorusecManager) *CustomResource {
	return (*CustomResource)(hm)
}

func (cr *CustomResource) GetAPIEndpoint() string {
	if cr.Spec.Components == nil || cr.Spec.Components.API == nil || cr.Spec.Components.API.Ingress == nil {
		return ""
	}

	ingress := cr.Spec.Components.API.Ingress
	u := &url.URL{Scheme: ingress.Scheme, Host: ingress.Host}
	return u.String()
}

func (cr *CustomResource) GetAnalyticEndpoint() string {
	if cr.Spec.Components == nil || cr.Spec.Components.Analytic == nil || cr.Spec.Components.Analytic.Ingress == nil {
		return ""
	}

	ingress := cr.Spec.Components.Analytic.Ingress
	u := &url.URL{Scheme: ingress.Scheme, Host: ingress.Host}
	return u.String()
}

func (cr *CustomResource) GetAccountEndpoint() string {
	if cr.Spec.Components == nil || cr.Spec.Components.Account == nil || cr.Spec.Components.Account.Ingress == nil {
		return ""
	}

	ingress := cr.Spec.Components.Account.Ingress
	u := &url.URL{Scheme: ingress.Scheme, Host: ingress.Host}
	return u.String()
}

func (cr *CustomResource) GetAuthEndpoint() string {
	if cr.Spec.Components == nil || cr.Spec.Components.Auth == nil || cr.Spec.Components.Auth.Ingress == nil {
		return ""
	}

	ingress := cr.Spec.Components.Auth.Ingress
	u := &url.URL{Scheme: ingress.Scheme, Host: ingress.Host}
	return u.String()
}

func (cr *CustomResource) GetManagerEndpoint() string {
	if cr.Spec.Components == nil || cr.Spec.Components.Manager == nil || cr.Spec.Components.Manager.Ingress == nil {
		return ""
	}

	ingress := cr.Spec.Components.Manager.Ingress
	u := &url.URL{Scheme: ingress.Scheme, Host: ingress.Host}
	return u.String()
}

func (cr *CustomResource) ToConfiguration() (*core.Configuration, error) {
	jo := jsonObject{
		"react_app_horusec_endpoint_account":  cr.GetAccountEndpoint(),
		"react_app_horusec_endpoint_analytic": cr.GetAnalyticEndpoint(),
		"react_app_horusec_endpoint_api":      cr.GetAPIEndpoint(),
		"react_app_horusec_endpoint_auth":     cr.GetAuthEndpoint(),
		"react_app_horusec_endpoint_manager":  cr.GetManagerEndpoint(),
	}

	var configuration *core.Configuration
	if err := jo.unmarshal(&configuration); err != nil {
		return nil, err
	}

	return configuration, nil
}
