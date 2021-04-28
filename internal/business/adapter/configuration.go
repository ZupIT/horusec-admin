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
	"github.com/ZupIT/horusec-admin/pkg/core"
	jsonpatch "github.com/evanphx/json-patch"
)

type Configuration core.Configuration

func ForConfiguration(configuration *core.Configuration) *Configuration {
	return (*Configuration)(configuration)
}

func ForConfigurationRaw(raw []byte) (*Configuration, error) {
	var c *Configuration
	if err := json.Unmarshal(raw, &c); err != nil {
		return nil, err
	}
	return c, nil
}

func (c Configuration) ToCustomResource() (*api.HorusecManager, error) {
	components, err := c.toComponents()
	if err != nil {
		return nil, err
	}

	return &api.HorusecManager{Spec: api.HorusecManagerSpec{Global: c.toGlobal(), Components: components}}, nil
}

func (c *Configuration) MergePatch(patch []byte) (*Configuration, error) {
	v1, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}

	return mergeConfigurations(v1, patch)
}

func mergeConfigurations(v1, v2 []byte) (*Configuration, error) {
	merged, err := jsonpatch.MergePatch(v1, v2)
	if err != nil {
		return nil, err
	}

	var result *Configuration
	if err := json.Unmarshal(merged, &result); err != nil {
		return nil, err
	}

	return result, nil
}
