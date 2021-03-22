// Copyright 2020 ZUP IT SERVICOS EM TECNOLOGIA E INOVACAO SA
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package authz

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewToken(t *testing.T) {
	t.Run("should successfully create a new authz instance", func(t *testing.T) {
		authz := New()
		assert.NotNil(t, authz)
	})

	t.Run("token and createAt values should be setted", func(t *testing.T) {
		authz := New()
		assert.NotEmpty(t, authz.token)
		assert.NotEmpty(t, authz.createdAt)
	})
}

func TestGetToken(t *testing.T) {
	t.Run("should return the token value", func(t *testing.T) {
		authz := New()
		assert.Equal(t, authz.token, authz.GetToken())
	})
}
