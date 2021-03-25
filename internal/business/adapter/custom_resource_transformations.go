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

func (cr *CustomResource) toGeneral() *core.General {
	admin := cr.toAdmin()

	var enabled bool
	if cr.Spec.Global != nil {
		enabled = cr.Spec.Global.EnableAdmin
	}

	secret := "horusec-jwt"
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

func (cr *CustomResource) toKeycloak() *core.Keycloak {
	var basePath, clientID, clientSecret, realm string
	var otp bool
	reactApp := cr.toKeycloakReactApp()

	if k := cr.GetKeycloak(); k != nil {
		basePath = k.InternalURL
		realm = k.Realm
		otp = k.OTP
	}

	if cc := cr.GetKeycloakConfidentialCredentials(); cc != nil {
		clientID = cc.ID
		clientSecret = cc.ID
	}

	if basePath != "" || clientID != "" || clientSecret != "" || realm != "" || otp || reactApp != nil {
		return &core.Keycloak{
			BasePath:         basePath,
			ClientID:         clientID,
			ClientSecret:     clientSecret,
			Realm:            realm,
			OTP:              otp,
			KeycloakReactApp: reactApp,
		}
	}

	return nil
}

func (cr *CustomResource) toKeycloakReactApp() *core.KeycloakReactApp {
	var clientID, realm, basePath string

	if k := cr.GetKeycloak(); k != nil {
		basePath = k.PublicURL
		realm = k.Realm
	}

	pc := cr.GetKeycloakPublicCredentials()
	if pc != nil {
		clientID = pc.ID
	}

	if clientID != "" || realm != "" || basePath != "" {
		return &core.KeycloakReactApp{ClientID: clientID, Realm: realm, BasePath: basePath}
	}

	return nil
}

func (cr *CustomResource) toManager() *core.Manager {
	apiEndpoint := "http://api.local/"
	analyticEndpoint := "http://analytic.local/"
	accountEndpoint := "http://account.local/"
	authEndpoint := "http://auth.local/"
	managerEndpoint := "http://manager.local/"
	managerPath := "/horusec"

	if u := cr.GetAPIURL(); u != nil {
		apiEndpoint = u.String()
	}
	if u := cr.GetAnalyticURL(); u != nil {
		analyticEndpoint = u.String()
	}
	if u := cr.GetAccountURL(); u != nil {
		accountEndpoint = u.String()
	}
	if u := cr.GetAuthURL(); u != nil {
		authEndpoint = u.String()
	}
	if u := cr.GetManagerURL(); u != nil {
		managerEndpoint = u.String()
	}

	return &core.Manager{
		APIEndpoint:      apiEndpoint,
		AnalyticEndpoint: analyticEndpoint,
		AccountEndpoint:  accountEndpoint,
		AuthEndpoint:     authEndpoint,
		ManagerEndpoint:  managerEndpoint,
		ManagerPath:      managerPath, // TODO: make manager path configurable
	}
}