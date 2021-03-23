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
	"github.com/ZupIT/horusec-admin/pkg/core"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestConfiguration_ToCustomResource(t *testing.T) {
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
}
