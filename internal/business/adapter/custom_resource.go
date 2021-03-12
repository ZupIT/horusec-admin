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
)

type CustomResource api.HorusecManager

func NewCustomResource(cfg *Configuration) (*CustomResource, error) {
	cr := new(CustomResource)

	if mng := cfg.Manager; mng != nil {
		err := cr.setAPIEndpoint(mng.APIEndpoint)
		if err != nil {
			return nil, err
		}

		err = cr.setAnalyticEndpoint(mng.AnalyticEndpoint)
		if err != nil {
			return nil, err
		}

		err = cr.setAccountEndpoint(mng.AccountEndpoint)
		if err != nil {
			return nil, err
		}

		err = cr.setAuthEndpoint(mng.AuthEndpoint)
		if err != nil {
			return nil, err
		}
	}

	return cr, nil
}

func (c *CustomResource) setAPIEndpoint(endpoint string) error {
	u, err := url.Parse(endpoint)
	if err != nil {
		return fmt.Errorf("fail to update the API endpoint: %w", err)
	}

	if c.Spec.Components == nil {
		c.Spec.Components = new(api.Components)
	}
	if c.Spec.Components.API == nil {
		c.Spec.Components.API = new(api.API)
	}

	component := c.Spec.Components.API
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

	if c.Spec.Components == nil {
		c.Spec.Components = new(api.Components)
	}
	if c.Spec.Components.Analytic == nil {
		c.Spec.Components.Analytic = new(api.Analytic)
	}

	component := c.Spec.Components.Analytic
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

	if c.Spec.Components == nil {
		c.Spec.Components = new(api.Components)
	}
	if c.Spec.Components.Account == nil {
		c.Spec.Components.Account = new(api.Account)
	}

	component := c.Spec.Components.Account
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

	if c.Spec.Components == nil {
		c.Spec.Components = new(api.Components)
	}
	if c.Spec.Components.Auth == nil {
		c.Spec.Components.Auth = new(api.Auth)
	}

	component := c.Spec.Components.Auth
	if component.Ingress == nil {
		component.Ingress = new(api.Ingress)
	}
	component.Ingress.Host = u.Host
	component.Ingress.Scheme = u.Scheme

	return nil
}
