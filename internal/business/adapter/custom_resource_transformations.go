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

import "github.com/ZupIT/horusec-admin/pkg/core"

const (
	defaultSecretName       = "horusec-jwt" // nolint:gosec // its not a hardcoded credential
	defaultAPIEndpoint      = "http://api.local/"
	defaultAnalyticEndpoint = "http://analytic.local/"
	defaultAccountEndpoint  = "http://account.local/"
	defaultAuthEndpoint     = "http://auth.local/"
	defaultManagerEndpoint  = "http://manager.local/"
	defaultManagerPath      = "/horusec"
)

func (cr *CustomResource) toGeneral() *core.General {
	admin := cr.toAdmin()

	var enabled bool
	if cr.Spec.Global != nil {
		enabled = cr.Spec.Global.EnableAdmin
	}

	secret := defaultSecretName
	if jwt := cr.GetJWT(); jwt != nil {
		secret = cr.Spec.Global.JWT.SecretName
	}

	return &core.General{EnableApplicationAdmin: enabled, JwtSecretKey: secret, ApplicationAdminData: admin}
}

func (cr *CustomResource) toAdmin() *core.Admin {
	// TODO: implement Admin integration
	return nil
}

func (cr *CustomResource) toAuth() *core.Auth {
	k := cr.toKeycloak()
	t := cr.GetAuthType()
	l := cr.toLDAP()

	if k != nil || l != nil || t != "" {
		return &core.Auth{Type: t, Keycloak: k, LDAP: l}
	}

	return nil
}

func (cr *CustomResource) toLDAP() *core.LDAP {
	// TODO: implement LDAP configurations
	return nil
}

//nolint:funlen // the number of statements cannot be reduced
func (cr *CustomResource) toKeycloak() *core.Keycloak {
	keycloak := &core.Keycloak{KeycloakReactApp: cr.toKeycloakReactApp()}

	if k := cr.GetKeycloak(); k != nil {
		keycloak.BasePath = k.InternalURL
		keycloak.Realm = k.Realm
		keycloak.OTP = k.OTP
	}

	if cc := cr.GetKeycloakConfidentialCredentials(); cc != nil {
		keycloak.ClientID = cc.ID
		keycloak.ClientSecret = cc.Secret
	}

	if (core.Keycloak{}) == *keycloak {
		return nil
	}

	return keycloak
}

func (cr *CustomResource) toKeycloakReactApp() *core.KeycloakReactApp {
	reactApp := new(core.KeycloakReactApp)
	if k := cr.GetKeycloak(); k != nil {
		reactApp.BasePath = k.PublicURL
		reactApp.Realm = k.Realm
	}
	if pc := cr.GetKeycloakPublicCredentials(); pc != nil {
		reactApp.ClientID = pc.ID
	}
	if (core.KeycloakReactApp{}) == *reactApp {
		return nil
	}

	return reactApp
}

//nolint:funlen,gocyclo // the number of statements and cyclomatic complexity cannot be reduced
func (cr *CustomResource) toManager() *core.Manager {
	m := &core.Manager{
		APIEndpoint:      defaultAPIEndpoint,
		AnalyticEndpoint: defaultAnalyticEndpoint,
		AccountEndpoint:  defaultAccountEndpoint,
		AuthEndpoint:     defaultAuthEndpoint,
		ManagerEndpoint:  defaultManagerEndpoint,
		ManagerPath:      defaultManagerPath, // TODO: make manager path configurable
	}

	if u := cr.GetAPIURL(); u != nil {
		m.APIEndpoint = u.String()
	}
	if u := cr.GetAnalyticURL(); u != nil {
		m.AnalyticEndpoint = u.String()
	}
	if u := cr.GetAccountURL(); u != nil {
		m.AccountEndpoint = u.String()
	}
	if u := cr.GetAuthURL(); u != nil {
		m.AuthEndpoint = u.String()
	}
	if u := cr.GetManagerURL(); u != nil {
		m.ManagerEndpoint = u.String()
	}

	return m
}
