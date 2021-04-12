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

func (c *Configuration) toGlobal() *api.Global {
	g := &api.Global{
		Administrator: c.toAdministrator(),
		JWT:           c.toJWT(),
		Broker:        c.toBroker(),
		Database:      c.toDatabase(),
		Keycloak:      c.toKeycloak(),
	}

	if (api.Global{}) == *g {
		return nil
	}

	return g
}

func (c *Configuration) toAdministrator() *api.Administrator {
	if c.General != nil && c.General.ApplicationAdminData != nil && c.General.EnableApplicationAdmin {
		return &api.Administrator{
			Enabled:  c.General.EnableApplicationAdmin,
			Username: c.General.ApplicationAdminData.Username,
			Email:    c.General.ApplicationAdminData.Email,
			Password: c.General.ApplicationAdminData.Password,
		}
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

//nolint:funlen // the number of statements cannot be reduced
func (c *Configuration) toKeycloak() *api.Keycloak {
	k := &api.Keycloak{Clients: c.toClients()}

	if keycloak := c.GetKeycloak(); keycloak != nil {
		k.InternalURL = keycloak.BasePath
		k.Realm = keycloak.Realm
		k.OTP = keycloak.OTP
	}

	if reactApp := c.GetKeycloakReactApp(); reactApp != nil {
		k.PublicURL = reactApp.BasePath
		if k.Realm == "" {
			k.Realm = reactApp.Realm
		}
	}

	if (api.Keycloak{}) == *k {
		return nil
	}

	return k
}

func (c *Configuration) toClients() *api.Clients {
	client := &api.Clients{
		Public:       c.GetKeycloakPublicCredentials(),
		Confidential: c.GetKeycloakConfidentialCredentials(),
	}

	if (api.Clients{}) == *client {
		return nil
	}

	return client
}

//nolint:funlen,gocyclo // the number of statements and cyclomatic complexity cannot be reduced
func (c *Configuration) toComponents() (*api.Components, error) {
	components := new(api.Components)

	if account, err := c.toAccount(); err == nil {
		components.Account = account
	} else {
		return nil, err
	}

	if analytic, err := c.toAnalytic(); err == nil {
		components.Analytic = analytic
	} else {
		return nil, err
	}

	if rapi, err := c.toAPI(); err == nil {
		components.API = rapi
	} else {
		return nil, err
	}

	if auth, err := c.toAuth(); err == nil {
		components.Auth = auth
	} else {
		return nil, err
	}

	if manager, err := c.toManager(); err == nil {
		components.Manager = manager
	} else {
		return nil, err
	}

	if (api.Components{} == *components) {
		return nil, nil
	}

	return components, nil
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
