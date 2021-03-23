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
	"github.com/ZupIT/horusec-admin/internal/kubernetes"
	"github.com/ZupIT/horusec-admin/internal/logger"
	"github.com/ZupIT/horusec-admin/internal/tracing"
	api "github.com/ZupIT/horusec-admin/pkg/api/install/v1alpha1"
	client "github.com/ZupIT/horusec-admin/pkg/client/clientset/versioned/typed/install/v1alpha1"
	"github.com/ZupIT/horusec-admin/pkg/core"
	k8s "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ConfigService struct {
	client     client.HorusecManagerInterface
	comparator *kubernetes.ObjectComparator
}

func NewConfigService(c client.HorusecManagerInterface, cmp *kubernetes.ObjectComparator) *ConfigService {
	return &ConfigService{client: c, comparator: cmp}
}

func (s *ConfigService) GetConfig(ctx context.Context) (*core.Configuration, error) {
	cr, err := s.getOne(ctx)
	if err != nil {
		return nil, err
	}

	if cr == nil {
		return new(core.Configuration), nil
	}

	return adapter.ForCustomResource(cr).ToConfiguration()
}

func (s *ConfigService) CreateOrUpdate(ctx context.Context, cfg *core.Configuration) error {
	r, err := adapter.ForConfiguration(cfg).ToCustomResource()
	if err != nil {
		return err
	}

	err = s.apply(ctx, r)
	if err != nil {
		return err
	}

	return nil
}

func (s *ConfigService) getOne(ctx context.Context) (*api.HorusecManager, error) {
	log := logger.WithPrefix(ctx, "config_service")

	hm, err := s.list(ctx)
	if err != nil {
		return nil, err
	}

	if len(hm) > 1 {
		return nil, errors.New("more than one HorusecManager CR were found")
	} else if len(hm) == 0 {
		log.Debug("no HorusecManager was found")
		return nil, nil
	}

	return &hm[0], nil
}

func (s *ConfigService) apply(ctx context.Context, r *api.HorusecManager) error {
	o, err := s.getOne(ctx)
	if err != nil {
		return err
	}

	if o == nil {
		r.SetName("horusec")
		return s.create(ctx, r)
	}

	return s.updateIfNeeded(ctx, r, o)
}

func (s *ConfigService) list(ctx context.Context) ([]api.HorusecManager, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "internal/business.(*ConfigService).list")
	defer span.Finish()

	cfg, err := s.client.List(ctx, k8s.ListOptions{})
	if err != nil {
		span.SetError(err)
		return nil, fmt.Errorf("failed to list HorusecManager: %w", err)
	}
	return cfg.Items, nil
}

func (s *ConfigService) create(ctx context.Context, r *api.HorusecManager) error {
	span, ctx := tracing.StartSpanFromContext(ctx, "internal/business.(*ConfigService).create")
	defer span.Finish()

	_, err := s.client.Create(ctx, r, k8s.CreateOptions{})
	if err != nil {
		span.SetError(err)
		return fmt.Errorf("failed to create HorusecManager: %w", err)
	}

	logger.WithPrefix(ctx, "config_service").Debug("resource created")
	return nil
}

func (s *ConfigService) updateIfNeeded(ctx context.Context, newest, older *api.HorusecManager) error {
	log := logger.WithPrefix(ctx, "config_service")

	newest.SetName(older.GetName())
	newest.SetResourceVersion(older.GetResourceVersion())
	if diff := s.comparator.Diff(older, newest); diff != "" {
		log.WithField("diff", diff).Debug("resource changed")
		if err := s.update(ctx, newest); err != nil {
			return err
		}
	} else {
		log.Debug("resource unchanged")
	}
	return nil
}

func (s *ConfigService) update(ctx context.Context, r *api.HorusecManager) error {
	span, ctx := tracing.StartSpanFromContext(ctx, "internal/business.(*ConfigService).update")
	defer span.Finish()

	_, err := s.client.Update(ctx, r, k8s.UpdateOptions{})
	if err != nil {
		span.SetError(err)
		return fmt.Errorf("failed to update HorusecManager: %w", err)
	}

	logger.WithPrefix(ctx, "config_service").Debug("resource updated")
	return nil
}
