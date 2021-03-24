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
	"fmt"
	api "github.com/ZupIT/horusec-admin/pkg/api/install/v1alpha1"
	"net/url"

	"github.com/ZupIT/horusec-admin/pkg/core"
)

type Configuration core.Configuration

func ForConfiguration(configuration *core.Configuration) *Configuration {
	return (*Configuration)(configuration)
}

func (c Configuration) ToCustomResource() (*api.HorusecManager, error) {
	components := jsonObject{}

	if aj, err := c.accountJSON(); err != nil {
		return nil, err
	} else if aj != nil {
		components["account"] = aj
	}

	if aj, err := c.analyticJSON(); err != nil {
		return nil, err
	} else if aj != nil {
		components["analytic"] = aj
	}

	if aj, err := c.apiJSON(); err != nil {
		return nil, err
	} else if aj != nil {
		components["api"] = aj
	}

	if aj, err := c.authJSON(); err != nil {
		return nil, err
	} else if aj != nil {
		components["auth"] = aj
	}

	if mj, err := c.managerJSON(); err != nil {
		return nil, err
	} else if mj != nil {
		components["manager"] = mj
	}

	jo := jsonObject{}
	if k, err := c.keycloakJSON(); err != nil {
		return nil, err
	} else if k != nil {
		jo["global"] = jsonObject{"keycloak": k}
	}

	if len(components) > 0 {
		jo["components"] = components
	}

	var spec api.HorusecManagerSpec
	if err := jo.unmarshal(&spec); err != nil {
		return nil, err
	}

	return &api.HorusecManager{Spec: spec}, nil
}

func (c *Configuration) GetAuthType() string {
	if c.Auth == nil || c.Auth.Type == "horusec" {
		return ""
	}

	return c.Auth.Type
}

func (c *Configuration) GetAccountURL() (*url.URL, error) {
	if c.Manager == nil {
		return nil, nil
	}

	u, err := url.Parse(c.AccountEndpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to parse Account URL: %w", err)
	}

	return u, nil
}

func (c *Configuration) GetAnalyticURL() (*url.URL, error) {
	if c.Manager == nil {
		return nil, nil
	}

	u, err := url.Parse(c.AnalyticEndpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to parse Analytic URL: %w", err)
	}

	return u, nil
}

func (c *Configuration) GetAPIURL() (*url.URL, error) {
	if c.Manager == nil {
		return nil, nil
	}

	u, err := url.Parse(c.APIEndpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to parse API URL: %w", err)
	}

	return u, nil
}

func (c *Configuration) GetAuthURL() (*url.URL, error) {
	if c.Manager == nil {
		return nil, nil
	}

	u, err := url.Parse(c.AuthEndpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to parse Auth URL: %w", err)
	}

	return u, nil
}

func (c *Configuration) GetManagerURL() (*url.URL, error) {
	if c.Manager == nil {
		return nil, nil
	}

	u, err := url.Parse(c.ManagerEndpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to parse Manager URL: %w", err)
	}

	return u, nil
}

func (c Configuration) accountJSON() (jsonObject, error) {
	jo := jsonObject{}

	u, err := c.GetAccountURL()
	if err != nil {
		return nil, err
	}

	if u != nil {
		jo["ingress"] = jsonObject{"host": u.Host, "scheme": u.Scheme}
	}

	if len(jo) > 0 {
		return jo, nil
	}

	return nil, nil
}

func (c Configuration) analyticJSON() (jsonObject, error) {
	jo := jsonObject{}

	u, err := c.GetAnalyticURL()
	if err != nil {
		return nil, err
	}

	if u != nil {
		jo["ingress"] = jsonObject{"host": u.Host, "scheme": u.Scheme}
	}

	if len(jo) > 0 {
		return jo, nil
	}

	return nil, nil
}

func (c Configuration) apiJSON() (jsonObject, error) {
	jo := jsonObject{}

	u, err := c.GetAPIURL()
	if err != nil {
		return nil, err
	}

	if u != nil {
		jo["ingress"] = jsonObject{"host": u.Host, "scheme": u.Scheme}
	}

	if len(jo) > 0 {
		return jo, nil
	}

	return nil, nil
}

func (c Configuration) authJSON() (jsonObject, error) {
	jo := jsonObject{}

	if at := c.GetAuthType(); at != "" {
		jo["type"] = at
	}

	u, err := c.GetAuthURL()
	if err != nil {
		return nil, err
	}

	if u != nil {
		jo["ingress"] = jsonObject{"host": u.Host, "scheme": u.Scheme}
	}

	if len(jo) > 0 {
		return jo, nil
	}

	return nil, nil
}

func (c Configuration) managerJSON() (jsonObject, error) {
	jo := jsonObject{}

	u, err := c.GetManagerURL()
	if err != nil {
		return nil, err
	}

	if u != nil {
		jo["ingress"] = jsonObject{"host": u.Host, "scheme": u.Scheme}
	}

	if len(jo) > 0 {
		return jo, nil
	}

	return nil, nil
}

func (c Configuration) keycloakJSON() (jsonObject, error) {
	if c.Auth == nil || c.Auth.Keycloak == nil {
		return nil, nil
	}

	jo := jsonObject{}

	if c.Auth.Keycloak.BasePath != "" {
		jo["internalURL"] = c.Auth.Keycloak.BasePath
	}
	if c.Auth.Keycloak.Realm != "" {
		jo["realm"] = c.Auth.Keycloak.Realm
	}
	if c.Auth.Keycloak.OTP {
		jo["otp"] = c.Auth.Keycloak.OTP
	}

	confidentialClient := jsonObject{}
	if c.Auth.Keycloak.ClientID != "" {
		confidentialClient["id"] = c.Auth.Keycloak.ClientID
	}
	if c.Auth.Keycloak.ClientSecret != "" {
		confidentialClient["secret"] = c.Auth.Keycloak.ClientSecret
	}

	clients := jsonObject{}
	if len(confidentialClient) > 0 {
		clients["confidential"] = confidentialClient
	}

	if c.Auth.Keycloak.KeycloakReactApp != nil {
		if c.Auth.Keycloak.KeycloakReactApp.ClientID != "" {
			clients["public"] = jsonObject{"id": c.Auth.Keycloak.KeycloakReactApp.ClientID}
		}
		if c.Auth.Keycloak.KeycloakReactApp.Realm != "" && jo["realm"] == nil {
			jo["realm"] = c.Auth.Keycloak.KeycloakReactApp.Realm
		}
		if c.Auth.Keycloak.KeycloakReactApp.BasePath != "" {
			jo["publicURL"] = c.Auth.Keycloak.KeycloakReactApp.BasePath
		}
	}

	if len(clients) > 0 {
		jo["clients"] = clients
	}

	if len(jo) > 0 {
		return jo, nil
	}

	return nil, nil
}
