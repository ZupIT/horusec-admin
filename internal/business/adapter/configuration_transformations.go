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

import api "github.com/ZupIT/horusec-admin/pkg/api/install/v1alpha1"

// nolint:gocyclo // if all Global fields are empty then it should return nil
func (c *Configuration) toGlobal() *api.Global {
	jwt := c.toJWT()
	broker := c.toBroker()
	database := c.toDatabase()
	keycloak := c.toKeycloak()

	var enableAdmin bool
	if c.General != nil {
		enableAdmin = c.General.EnableApplicationAdmin
	}

	if enableAdmin || jwt != nil || broker != nil || database != nil || keycloak != nil {
		return &api.Global{EnableAdmin: enableAdmin, JWT: jwt, Broker: broker, Database: database, Keycloak: keycloak}
	}

	return nil
}

func (c *Configuration) toJWT() *api.JWT {
	if c.General != nil && c.General.JwtSecretKey != "" {
		return &api.JWT{SecretName: c.General.JwtSecretKey}
	}

	return nil
}

func (c *Configuration) toBroker() *api.Broker {
	// TODO: implement RabbitMQ configurations
	return nil
}

func (c *Configuration) toDatabase() *api.Database {
	// TODO: implement PostgreSQL configurations
	return nil
}

//nolint:funlen,gocyclo // the complexity for this object conversion cannot be reduced
func (c *Configuration) toKeycloak() *api.Keycloak {
	var otp bool
	var publicURL, internalURL, realm string
	clients := c.toClients()

	if keycloak := c.GetKeycloak(); keycloak != nil {
		internalURL = keycloak.BasePath
		realm = keycloak.Realm
		otp = keycloak.OTP
	}

	reactApp := c.GetKeycloakReactApp()
	if reactApp != nil {
		publicURL = reactApp.BasePath
		if realm == "" {
			realm = reactApp.Realm
		}
	}

	if publicURL != "" || internalURL != "" || realm != "" || otp || clients != nil {
		return &api.Keycloak{PublicURL: publicURL, InternalURL: internalURL, Realm: realm, OTP: otp, Clients: clients}
	}

	return nil
}

//nolint:funlen,gocyclo // these linters are not feasible for objects adapters
func (c *Configuration) toClients() *api.Clients {
	var confidential, public *api.ClientCredentials

	keycloak := c.GetKeycloak()
	if keycloak != nil && (keycloak.ClientID != "" || keycloak.ClientSecret != "") {
		confidential = &api.ClientCredentials{ID: keycloak.ClientID, Secret: keycloak.ClientSecret}
	}

	reactApp := c.GetKeycloakReactApp()
	if reactApp != nil && reactApp.ClientID != "" {
		public = &api.ClientCredentials{ID: reactApp.ClientID}
	}

	if confidential != nil || public != nil {
		return &api.Clients{Confidential: confidential, Public: public}
	}

	return nil
}

//nolint:funlen,gocyclo // these linters are not feasible for objects adapters
func (c *Configuration) toComponents() (*api.Components, error) {
	account, err := c.toAccount()
	if err != nil {
		return nil, err
	}

	analytic, err := c.toAnalytic()
	if err != nil {
		return nil, err
	}

	rapi, err := c.toAPI()
	if err != nil {
		return nil, err
	}

	auth, err := c.toAuth()
	if err != nil {
		return nil, err
	}

	manager, err := c.toManager()
	if err != nil {
		return nil, err
	}

	if account != nil || analytic != nil || rapi != nil || auth != nil || manager != nil {
		return &api.Components{Account: account, Analytic: analytic, API: rapi, Auth: auth, Manager: manager}, nil
	}

	return nil, nil
}

func (c *Configuration) toAccount() (*api.Account, error) {
	u, err := c.GetAccountURL()
	if err != nil {
		return nil, err
	}

	if u != nil {
		return &api.Account{Ingress: &api.Ingress{Scheme: u.Scheme, Host: u.Host}}, nil
	}

	return nil, nil
}

func (c *Configuration) toAnalytic() (*api.Analytic, error) {
	u, err := c.GetAnalyticURL()
	if err != nil {
		return nil, err
	}

	if u != nil {
		return &api.Analytic{Ingress: &api.Ingress{Scheme: u.Scheme, Host: u.Host}}, nil
	}

	return nil, nil
}

func (c *Configuration) toAPI() (*api.API, error) {
	u, err := c.GetAPIURL()
	if err != nil {
		return nil, err
	}

	if u != nil {
		return &api.API{Ingress: &api.Ingress{Scheme: u.Scheme, Host: u.Host}}, nil
	}

	return nil, nil
}

//nolint:funlen // breaking the method 'toAuth' is infeasible
func (c *Configuration) toAuth() (*api.Auth, error) {
	u, err := c.GetAuthURL()
	if err != nil {
		return nil, err
	}

	var i *api.Ingress
	if u != nil {
		i = &api.Ingress{Scheme: u.Scheme, Host: u.Host}
	}

	at := c.GetAuthType()
	if i != nil || at != "" {
		return &api.Auth{Type: at, Ingress: i}, nil
	}

	return nil, nil
}

func (c *Configuration) toManager() (*api.Manager, error) {
	u, err := c.GetManagerURL()
	if err != nil {
		return nil, err
	}

	if u != nil {
		return &api.Manager{Ingress: &api.Ingress{Scheme: u.Scheme, Host: u.Host}}, nil
	}

	return nil, nil
}
