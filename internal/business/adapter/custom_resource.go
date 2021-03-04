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
		err := cr.SetAPIEndpoint(mng.APIEndpoint)
		if err != nil {
			return nil, err
		}

		err = cr.SetAnalyticEndpoint(mng.AnalyticEndpoint)
		if err != nil {
			return nil, err
		}

		err = cr.SetAccountEndpoint(mng.AccountEndpoint)
		if err != nil {
			return nil, err
		}

		err = cr.SetAuthEndpoint(mng.AuthEndpoint)
		if err != nil {
			return nil, err
		}
	}

	return cr, nil
}

func (c *CustomResource) SetAPIEndpoint(endpoint string) error {
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

func (c *CustomResource) SetAnalyticEndpoint(endpoint string) error {
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

func (c *CustomResource) SetAccountEndpoint(endpoint string) error {
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

func (c *CustomResource) SetAuthEndpoint(endpoint string) error {
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
