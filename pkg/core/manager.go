package core

import (
	"net/url"

	"github.com/ZupIT/horusec-admin/pkg/api/install/v1alpha1"
)

type Manager struct {
	APIEndpoint      string `json:"react_app_horusec_endpoint_api"`
	AnalyticEndpoint string `json:"react_app_horusec_endpoint_analytic"`
	AccountEndpoint  string `json:"react_app_horusec_endpoint_account"`
	AuthEndpoint     string `json:"react_app_horusec_endpoint_auth"`
	ManagerPath      string `json:"react_app_horusec_manager_path"`
}

func newManager(cr *v1alpha1.HorusecManager) *Manager {
	components := cr.Spec.Components
	api := (&url.URL{Scheme: components.API.Ingress.Scheme, Host: components.API.Ingress.Host}).String()
	analytic := (&url.URL{Scheme: components.Analytic.Ingress.Scheme, Host: components.Analytic.Ingress.Host}).String()
	account := (&url.URL{Scheme: components.Account.Ingress.Scheme, Host: components.Account.Ingress.Host}).String()
	auth := (&url.URL{Scheme: components.Auth.Ingress.Scheme, Host: components.Auth.Ingress.Host}).String()
	return &Manager{
		APIEndpoint:      api,
		AnalyticEndpoint: analytic,
		AccountEndpoint:  account,
		AuthEndpoint:     auth,
		ManagerPath:      "", // TODO: add this field to the Operator
	}
}
