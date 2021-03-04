package core

import (
	"fmt"
	"net/url"

	"github.com/ZupIT/horusec-admin/pkg/api/install/v1alpha1"
)

type customResource v1alpha1.HorusecManager

func (c *customResource) SetAPIEndpoint(endpoint string) error {
	u, err := url.Parse(endpoint)
	if err != nil {
		return fmt.Errorf("fail to update the API endpoint: %w", err)
	}

	component := c.Spec.Components.API
	component.Ingress.Host = u.Host
	component.Ingress.Scheme = u.Scheme

	return nil
}

func (c *customResource) SetAnalyticEndpoint(endpoint string) error {
	u, err := url.Parse(endpoint)
	if err != nil {
		return fmt.Errorf("fail to update the Analytic endpoint: %w", err)
	}

	component := c.Spec.Components.Analytic
	component.Ingress.Host = u.Host
	component.Ingress.Scheme = u.Scheme

	return nil
}

func (c *customResource) SetAccountEndpoint(endpoint string) error {
	u, err := url.Parse(endpoint)
	if err != nil {
		return fmt.Errorf("fail to update the Account endpoint: %w", err)
	}

	component := c.Spec.Components.Account
	component.Ingress.Host = u.Host
	component.Ingress.Scheme = u.Scheme

	return nil
}

func (c *customResource) SetAuthEndpoint(endpoint string) error {
	u, err := url.Parse(endpoint)
	if err != nil {
		return fmt.Errorf("fail to update the Auth endpoint: %w", err)
	}

	component := c.Spec.Components.Auth
	component.Ingress.Host = u.Host
	component.Ingress.Scheme = u.Scheme

	return nil
}

func newCR(cfg *Configuration) (*customResource, error) {
	cr := new(customResource)

	mng := cfg.Manager
	if mng != nil {
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
