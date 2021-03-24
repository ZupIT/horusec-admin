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
	api "github.com/ZupIT/horusec-admin/pkg/api/install/v1alpha1"
	"github.com/ZupIT/horusec-admin/pkg/core"
)

type CustomResource api.HorusecManager

func ForCustomResource(hm *api.HorusecManager) *CustomResource {
	return (*CustomResource)(hm)
}

func (cr *CustomResource) ToConfiguration() *core.Configuration {
	general := cr.toGeneral()
	auth := cr.toAuth()
	manager := cr.toManager()

	if general == nil && auth == nil && manager == nil {
		return nil
	}

	return &core.Configuration{General: general, Auth: auth, Manager: manager}
}
