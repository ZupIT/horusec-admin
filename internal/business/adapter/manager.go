package adapter

import (
	"net/url"

	"github.com/ZupIT/horusec-admin/pkg/core"

	api "github.com/ZupIT/horusec-admin/pkg/api/install/v1alpha1"
)

func newManager(cr *api.HorusecManager) *core.Manager {
	components := cr.Spec.Components
	api := (&url.URL{Scheme: components.API.Ingress.Scheme, Host: components.API.Ingress.Host}).String()
	analytic := (&url.URL{Scheme: components.Analytic.Ingress.Scheme, Host: components.Analytic.Ingress.Host}).String()
	account := (&url.URL{Scheme: components.Account.Ingress.Scheme, Host: components.Account.Ingress.Host}).String()
	auth := (&url.URL{Scheme: components.Auth.Ingress.Scheme, Host: components.Auth.Ingress.Host}).String()
	return &core.Manager{
		APIEndpoint:      api,
		AnalyticEndpoint: analytic,
		AccountEndpoint:  account,
		AuthEndpoint:     auth,
		ManagerPath:      "", // TODO: add this field to the Operator
	}
}
