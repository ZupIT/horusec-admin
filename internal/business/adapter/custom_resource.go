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
	"net/url"

	api "github.com/ZupIT/horusec-admin/pkg/api/install/v1alpha1"
	"github.com/ZupIT/horusec-admin/pkg/core"
)

type CustomResource api.HorusecManager

func NewCustomResource(cfg *Configuration) (*CustomResource, error) {
	cr := new(CustomResource)
	if err := cr.setManager(cfg.Manager); err != nil {
		return nil, err
	}

	return cr, nil
}

// nolint:funlen,gocyclo // setManager method needs to set all manager fields
func (c *CustomResource) setManager(manager *core.Manager) error {
	if manager == nil {
		return nil
	} else if err := c.setAPIEndpoint(manager.APIEndpoint); err != nil {
		return err
	} else if err := c.setAnalyticEndpoint(manager.AnalyticEndpoint); err != nil {
		return err
	} else if err := c.setAccountEndpoint(manager.AccountEndpoint); err != nil {
		return err
	} else if err := c.setAuthEndpoint(manager.AuthEndpoint); err != nil {
		return err
	}

	return nil
}

func (c *CustomResource) setAPIEndpoint(endpoint string) error {
	u, err := url.Parse(endpoint)
	if err != nil {
		return fmt.Errorf("fail to update the API endpoint: %w", err)
	}

	component := c.initAPI()
	if component.Ingress == nil {
		component.Ingress = new(api.Ingress)
	}
	component.Ingress.Host = u.Host
	component.Ingress.Scheme = u.Scheme

	return nil
}

func (c *CustomResource) setAnalyticEndpoint(endpoint string) error {
	u, err := url.Parse(endpoint)
	if err != nil {
		return fmt.Errorf("fail to update the Analytic endpoint: %w", err)
	}

	component := c.initAnalytic()
	if component.Ingress == nil {
		component.Ingress = new(api.Ingress)
	}
	component.Ingress.Host = u.Host
	component.Ingress.Scheme = u.Scheme

	return nil
}

func (c *CustomResource) setAccountEndpoint(endpoint string) error {
	u, err := url.Parse(endpoint)
	if err != nil {
		return fmt.Errorf("fail to update the Account endpoint: %w", err)
	}

	component := c.initAccount()
	if component.Ingress == nil {
		component.Ingress = new(api.Ingress)
	}
	component.Ingress.Host = u.Host
	component.Ingress.Scheme = u.Scheme

	return nil
}

func (c *CustomResource) setAuthEndpoint(endpoint string) error {
	u, err := url.Parse(endpoint)
	if err != nil {
		return fmt.Errorf("fail to update the Auth endpoint: %w", err)
	}

	component := c.initAuth()
	if component.Ingress == nil {
		component.Ingress = new(api.Ingress)
	}
	component.Ingress.Host = u.Host
	component.Ingress.Scheme = u.Scheme

	return nil
}

func (c *CustomResource) initAPI() *api.API {
	if c.Spec.Components == nil {
		c.Spec.Components = new(api.Components)
	}
	if c.Spec.Components.API == nil {
		c.Spec.Components.API = new(api.API)
	}
	return c.Spec.Components.API
}

func (c *CustomResource) initAnalytic() *api.Analytic {
	if c.Spec.Components == nil {
		c.Spec.Components = new(api.Components)
	}
	if c.Spec.Components.Analytic == nil {
		c.Spec.Components.Analytic = new(api.Analytic)
	}
	return c.Spec.Components.Analytic
}

func (c *CustomResource) initAccount() *api.Account {
	if c.Spec.Components == nil {
		c.Spec.Components = new(api.Components)
	}
	if c.Spec.Components.Account == nil {
		c.Spec.Components.Account = new(api.Account)
	}
	return c.Spec.Components.Account
}

func (c *CustomResource) initAuth() *api.Auth {
	if c.Spec.Components == nil {
		c.Spec.Components = new(api.Components)
	}
	if c.Spec.Components.Auth == nil {
		c.Spec.Components.Auth = new(api.Auth)
	}
	return c.Spec.Components.Auth
}
