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

	"github.com/ZupIT/horusec-admin/pkg/api/install/v1alpha1"
)

func (cr *CustomResource) IsAdministratorEnabled() bool {
	if cr.Spec.Global != nil && cr.Spec.Global.Administrator != nil {
		return cr.Spec.Global.Administrator.Enabled
	}

	return false
}

func (cr *CustomResource) GetJWT() *v1alpha1.JWT {
	if cr.Spec.Global != nil {
		return cr.Spec.Global.JWT
	}

	return nil
}

func (cr *CustomResource) GetAuthType() string {
	if cr.Spec.Components == nil || cr.Spec.Components.Auth == nil || cr.Spec.Components.Auth.Type == "" {
		return "horusec"
	}

	return cr.Spec.Components.Auth.Type
}

func (cr *CustomResource) GetKeycloak() *v1alpha1.Keycloak {
	if cr.Spec.Global != nil {
		return cr.Spec.Global.Keycloak
	}

	return nil
}

func (cr *CustomResource) GetKeycloakPublicCredentials() *v1alpha1.ClientCredentials {
	if cr.Spec.Global != nil && cr.Spec.Global.Keycloak != nil && cr.Spec.Global.Keycloak.Clients != nil {
		return cr.Spec.Global.Keycloak.Clients.Public
	}

	return nil
}

func (cr *CustomResource) GetKeycloakConfidentialCredentials() *v1alpha1.ClientCredentials {
	if cr.Spec.Global != nil && cr.Spec.Global.Keycloak != nil && cr.Spec.Global.Keycloak.Clients != nil {
		return cr.Spec.Global.Keycloak.Clients.Confidential
	}

	return nil
}

func (cr *CustomResource) GetAccountURL() *url.URL {
	if cr.Spec.Components == nil || cr.Spec.Components.Account == nil || cr.Spec.Components.Account.Ingress == nil {
		return nil
	}

	return &url.URL{Scheme: cr.Spec.Components.Account.Ingress.Scheme, Host: cr.Spec.Components.Account.Ingress.Host}
}

func (cr *CustomResource) GetAnalyticURL() *url.URL {
	if cr.Spec.Components == nil || cr.Spec.Components.Analytic == nil || cr.Spec.Components.Analytic.Ingress == nil {
		return nil
	}

	return &url.URL{Scheme: cr.Spec.Components.Analytic.Ingress.Scheme, Host: cr.Spec.Components.Analytic.Ingress.Host}
}

func (cr *CustomResource) GetAPIURL() *url.URL {
	if cr.Spec.Components == nil || cr.Spec.Components.API == nil || cr.Spec.Components.API.Ingress == nil {
		return nil
	}

	return &url.URL{Scheme: cr.Spec.Components.API.Ingress.Scheme, Host: cr.Spec.Components.API.Ingress.Host}
}

func (cr *CustomResource) GetAuthURL() *url.URL {
	if cr.Spec.Components == nil || cr.Spec.Components.Auth == nil || cr.Spec.Components.Auth.Ingress == nil {
		return nil
	}

	return &url.URL{Scheme: cr.Spec.Components.Auth.Ingress.Scheme, Host: cr.Spec.Components.Auth.Ingress.Host}
}

func (cr *CustomResource) GetManagerURL() *url.URL {
	if cr.Spec.Components == nil || cr.Spec.Components.Manager == nil || cr.Spec.Components.Manager.Ingress == nil {
		return nil
	}

	return &url.URL{Scheme: cr.Spec.Components.Manager.Ingress.Scheme, Host: cr.Spec.Components.Manager.Ingress.Host}
}
