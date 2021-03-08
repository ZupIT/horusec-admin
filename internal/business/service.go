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

package business

import (
	"context"
	"errors"
	"fmt"

	"github.com/ZupIT/horusec-admin/internal/business/adapter"
	"github.com/ZupIT/horusec-admin/internal/logger"
	api "github.com/ZupIT/horusec-admin/pkg/api/install/v1alpha1"
	client "github.com/ZupIT/horusec-admin/pkg/client/clientset/versioned/typed/install/v1alpha1"
	"github.com/ZupIT/horusec-admin/pkg/core"
	"github.com/google/go-cmp/cmp"
	k8s "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ConfigService struct {
	client      client.HorusecManagerInterface
	compareOpts cmp.Option
}

func NewConfigService(client client.HorusecManagerInterface) *ConfigService {
	ignore := [...]string{
		"ObjectMeta.CreationTimestamp", "ObjectMeta.Finalizers", "ObjectMeta.Generation",
		"ObjectMeta.ManagedFields", "ObjectMeta.Namespace", "ObjectMeta.ResourceVersion", "ObjectMeta.SelfLink",
		"ObjectMeta.UID", "TypeMeta.APIVersion",
	}
	return &ConfigService{
		client: client,
		compareOpts: cmp.FilterPath(func(path cmp.Path) bool {
			for _, p := range ignore {
				if p == path.String() {
					return true
				}
			}
			return false
		}, cmp.Ignore()),
	}
}

func (s *ConfigService) GetConfig() (*core.Configuration, error) {
	cr, err := s.getOne()
	if err != nil {
		return nil, err
	}

	if cr == nil {
		return new(core.Configuration), nil
	}

	cfg := adapter.NewConfiguration(cr)
	return (*core.Configuration)(cfg), nil
}

func (s *ConfigService) CreateOrUpdate(cfg *core.Configuration) error {
	r, err := (*adapter.Configuration)(cfg).CR()
	if err != nil {
		return err
	}

	err = s.apply(r)
	if err != nil {
		return err
	}

	return nil
}

func (s *ConfigService) getOne() (*api.HorusecManager, error) {
	cfg, err := s.client.List(context.TODO(), k8s.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to get configuration: %w", err)
	}

	size := len(cfg.Items)
	if size > 1 {
		return nil, errors.New("more than one HorusecManager CR were found")
	}

	if size == 0 {
		return nil, nil
	}

	return &cfg.Items[0], nil
}

func (s *ConfigService) apply(r *api.HorusecManager) error {
	log := logger.WithPrefix("service")

	o, err := s.getOne()
	if err != nil {
		return err
	}
	if o == nil {
		r.SetName("horusec")
		_, err = s.client.Create(context.TODO(), r, k8s.CreateOptions{})
		if err != nil {
			return err
		}
		log.Debug("resource created")
		return nil
	}

	r.SetName(o.GetName())
	r.SetResourceVersion(o.GetResourceVersion())
	diff := cmp.Diff(o, r, s.compareOpts)
	if diff != "" {
		log.Debug("resource changes:\n" + diff)
		_, err = s.client.Update(context.TODO(), r, k8s.UpdateOptions{})
		if err != nil {
			return err
		}
		log.Debug("resource updated")
	} else {
		log.Debug("resource unchanged")
	}
	return nil
}
