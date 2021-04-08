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
	"encoding/json"
	"testing"

	"github.com/ZupIT/horusec-admin/pkg/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConfiguration_ToCustomResource(t *testing.T) {
	t.Run("SHOULD marshal to expected json WHEN auth type is set to a non-default value", func(t *testing.T) {
		expected := "{\"components\":{\"auth\":{\"type\":\"keycloak\"}}}"

		cfg := &core.Configuration{Auth: &core.Auth{Type: "keycloak"}}

		cr, err := ForConfiguration(cfg).ToCustomResource()
		require.NoError(t, err)

		b, err := json.Marshal(cr.Spec)
		require.NoError(t, err)

		assert.Equal(t, expected, string(b))
	})

	t.Run("SHOULD marshal to an empty json WHEN auth type is set to default value 'horusec'", func(t *testing.T) {
		expected := "{}"

		cfg := &core.Configuration{Auth: &core.Auth{Type: "horusec"}}

		cr, err := ForConfiguration(cfg).ToCustomResource()
		require.NoError(t, err)

		b, err := json.Marshal(cr.Spec)
		require.NoError(t, err)

		assert.Equal(t, expected, string(b))
	})

	t.Run("SHOULD marshal to expected json WHEN hosts and schemes are populated", func(t *testing.T) {
		expected := "{\"components\":{\"account\":{\"ingress\":{\"scheme\":\"http\",\"host\":\"account.horus.local\"}},\"analytic\":{\"ingress\":{\"scheme\":\"http\",\"host\":\"analytic.horus.local\"}},\"api\":{\"ingress\":{\"scheme\":\"http\",\"host\":\"api.horus.local\"}},\"auth\":{\"ingress\":{\"scheme\":\"http\",\"host\":\"auth.horus.local\"}},\"manager\":{\"ingress\":{\"scheme\":\"http\",\"host\":\"manager.horus.local\"}}}}"

		cfg := &core.Configuration{Manager: &core.Manager{
			APIEndpoint:      "http://api.horus.local",
			AnalyticEndpoint: "http://analytic.horus.local",
			AccountEndpoint:  "http://account.horus.local",
			AuthEndpoint:     "http://auth.horus.local",
			ManagerEndpoint:  "http://manager.horus.local",
		}}

		cr, err := ForConfiguration(cfg).ToCustomResource()
		require.NoError(t, err)

		b, err := json.Marshal(cr.Spec)
		require.NoError(t, err)

		assert.Equal(t, expected, string(b))
	})

	t.Run("SHOULD marshal to expected json WHEN keycloak configurations are populated", func(t *testing.T) {
		expected := "{\"global\":{\"keycloak\":{\"publicURL\":\"http://keycloak.iam/auth\",\"internalURL\":\"http://keycloak.iam.svc.cluster.local/auth\",\"realm\":\"zup\",\"clients\":{\"public\":{\"id\":\"horusec-frontend\"},\"confidential\":{\"id\":\"horusec-backend\",\"secret\":\"0548d0ba-0aea-4c76-b601-3d2dc5f30e6b\"}}}},\"components\":{\"auth\":{\"type\":\"keycloak\"}}}"

		otp := false
		cfg := &core.Configuration{Auth: &core.Auth{
			Type: "keycloak",
			Keycloak: &core.Keycloak{
				BasePath:     "http://keycloak.iam.svc.cluster.local/auth",
				ClientID:     "horusec-backend",
				ClientSecret: "0548d0ba-0aea-4c76-b601-3d2dc5f30e6b",
				Realm:        "zup",
				OTP:          &otp,
				KeycloakReactApp: &core.KeycloakReactApp{
					ClientID: "horusec-frontend",
					Realm:    "zup",
					BasePath: "http://keycloak.iam/auth",
				},
			},
		}}

		cr, err := ForConfiguration(cfg).ToCustomResource()
		require.NoError(t, err)

		b, err := json.Marshal(cr.Spec)
		require.NoError(t, err)

		assert.Equal(t, expected, string(b))
	})

	t.Run("SHOULD marshal to expected json WHEN just public keycloak configurations are populated", func(t *testing.T) {
		expected := "{\"global\":{\"keycloak\":{\"publicURL\":\"http://keycloak.iam/auth\",\"realm\":\"zup\",\"clients\":{\"public\":{\"id\":\"horusec-frontend\"}}}},\"components\":{\"auth\":{\"type\":\"keycloak\"}}}"

		cfg := &core.Configuration{Auth: &core.Auth{
			Type: "keycloak",
			Keycloak: &core.Keycloak{
				KeycloakReactApp: &core.KeycloakReactApp{
					ClientID: "horusec-frontend",
					Realm:    "zup",
					BasePath: "http://keycloak.iam/auth",
				},
			},
		}}

		cr, err := ForConfiguration(cfg).ToCustomResource()
		require.NoError(t, err)

		b, err := json.Marshal(cr.Spec)
		require.NoError(t, err)

		assert.Equal(t, expected, string(b))
	})
}
