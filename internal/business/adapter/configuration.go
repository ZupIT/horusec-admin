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
	"github.com/ZupIT/horusec-admin/pkg/core"
	"net/url"
)

type Configuration core.Configuration

func ForConfiguration(configuration *core.Configuration) *Configuration {
	return (*Configuration)(configuration)
}

func (c *Configuration) GetAccountURL() (*url.URL, error) {
	u, err := url.Parse(c.AccountEndpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to parse Account URL: %w", err)
	}

	return u, nil
}

func (c *Configuration) GetAnalyticURL() (*url.URL, error) {
	u, err := url.Parse(c.AnalyticEndpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to parse Analytic URL: %w", err)
	}

	return u, nil
}

func (c *Configuration) GetAPIURL() (*url.URL, error) {
	u, err := url.Parse(c.APIEndpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to parse API URL: %w", err)
	}

	return u, nil
}

func (c *Configuration) GetAuthURL() (*url.URL, error) {
	u, err := url.Parse(c.AuthEndpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to parse Auth URL: %w", err)
	}

	return u, nil
}

func (c *Configuration) GetManagerURL() (*url.URL, error) {
	u, err := url.Parse(c.ManagerEndpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to parse Manager URL: %w", err)
	}

	return u, nil
}

func (c Configuration) ToCustomResource() (*api.HorusecManager, error) {
	m := jsonObject{}

	if u, err := c.GetAccountURL(); err == nil {
		m["account"] = jsonObject{"ingress": jsonObject{"host": u.Host, "scheme": u.Scheme}}
	} else {
		return nil, err
	}
	if u, err := c.GetAnalyticURL(); err == nil {
		m["analytic"] = jsonObject{"ingress": jsonObject{"host": u.Host, "scheme": u.Scheme}}
	} else {
		return nil, err
	}
	if u, err := c.GetAPIURL(); err == nil {
		m["api"] = jsonObject{"ingress": jsonObject{"host": u.Host, "scheme": u.Scheme}}
	} else {
		return nil, err
	}
	if u, err := c.GetAuthURL(); err == nil {
		m["auth"] = jsonObject{"ingress": jsonObject{"host": u.Host, "scheme": u.Scheme}}
	} else {
		return nil, err
	}
	if u, err := c.GetManagerURL(); err == nil {
		m["manager"] = jsonObject{"ingress": jsonObject{"host": u.Host, "scheme": u.Scheme}}
	} else {
		return nil, err
	}

	jo := jsonObject{"components": m}
	var spec api.HorusecManagerSpec
	if err := jo.unmarshal(&spec); err != nil {
		return nil, err
	}

	return &api.HorusecManager{Spec: spec}, nil
}
