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
	api "github.com/ZupIT/horusec-admin/pkg/api/install/v1alpha1"
	"github.com/stretchr/testify/require"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCustomResource_ToConfiguration(t *testing.T) {
	t.Run("SHOULD marshal to expected json WHEN hosts and schemes are populated", func(t *testing.T) {
		expected := "{\"react_app_horusec_endpoint_api\":\"http://api.horus.local\",\"react_app_horusec_endpoint_analytic\":\"http://analytic.horus.local\",\"react_app_horusec_endpoint_account\":\"http://account.horus.local\",\"react_app_horusec_endpoint_auth\":\"http://auth.horus.local\",\"react_app_horusec_manager_path\":\"\"}"

		hm := &api.HorusecManager{Spec: api.HorusecManagerSpec{Components: &api.Components{
			Account:  &api.Account{Ingress: &api.Ingress{Host: "account.horus.local", Scheme: "http"}},
			Analytic: &api.Analytic{Ingress: &api.Ingress{Host: "analytic.horus.local", Scheme: "http"}},
			API:      &api.API{Ingress: &api.Ingress{Host: "api.horus.local", Scheme: "http"}},
			Auth:     &api.Auth{Ingress: &api.Ingress{Host: "auth.horus.local", Scheme: "http"}},
			Manager:  &api.Manager{Ingress: &api.Ingress{Host: "manager.horus.local", Scheme: "http"}}},
		}}

		cfg, err := ForCustomResource(hm).ToConfiguration()
		require.NoError(t, err)

		b, err := json.Marshal(cfg)
		require.NoError(t, err)

		assert.Equal(t, expected, string(b))
	})
}
