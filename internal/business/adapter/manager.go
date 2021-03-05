package adapter

import (
	"net/url"

	"github.com/ZupIT/horusec-admin/pkg/core"

	api "github.com/ZupIT/horusec-admin/pkg/api/install/v1alpha1"
)

func newManager(cr *api.HorusecManager) *core.Manager {
	mng := new(core.Manager)
	components := cr.Spec.Components
	if components != nil {
		mng.APIEndpoint = (&url.URL{Scheme: components.API.Ingress.Scheme, Host: components.API.Ingress.Host}).String()
		mng.AnalyticEndpoint = (&url.URL{Scheme: components.Analytic.Ingress.Scheme, Host: components.Analytic.Ingress.Host}).String()
		mng.AccountEndpoint = (&url.URL{Scheme: components.Account.Ingress.Scheme, Host: components.Account.Ingress.Host}).String()
		mng.AuthEndpoint = (&url.URL{Scheme: components.Auth.Ingress.Scheme, Host: components.Auth.Ingress.Host}).String()
	}
	return mng
}
