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
	"github.com/ZupIT/horusec-admin/internal/kubernetes"
	"github.com/ZupIT/horusec-admin/internal/logger"
	"github.com/ZupIT/horusec-admin/internal/tracing"
	api "github.com/ZupIT/horusec-admin/pkg/api/install/v2alpha1"
	client "github.com/ZupIT/horusec-admin/pkg/client/clientset/versioned/typed/install/v2alpha1"
	k8s "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/json"
)

type ConfigService struct {
	client     client.HorusecPlatformInterface
	comparator *kubernetes.ObjectComparator
}

func NewConfigService(c client.HorusecPlatformInterface, cmp *kubernetes.ObjectComparator) *ConfigService {
	return &ConfigService{client: c, comparator: cmp}
}

func (s *ConfigService) GetConfig(ctx context.Context) (*api.HorusecPlatform, error) {
	log := logger.WithPrefix(ctx, "config_service")

	hm, err := s.list(ctx)
	if err != nil {
		return nil, err
	}

	if len(hm) > 1 {
		return nil, errors.New("more than one HorusecPlatform CR were found")
	} else if len(hm) == 0 {
		log.Debug("no HorusecPlatform was found")
		return nil, nil
	}

	return &hm[0], nil
}

func (s *ConfigService) CreateOrUpdate(ctx context.Context, raw []byte) error {
	if raw == nil {
		return errors.New("not accept raw empty")
	}
	newEntity := &api.HorusecPlatform{}
	if err := json.Unmarshal(raw, newEntity); err != nil {
		return err
	}
	older, err := s.GetConfig(ctx)
	if err != nil {
		return err
	}
	if older == nil {
		return s.createResource(ctx, newEntity)
	}
	return s.updatePartially(ctx, older, newEntity)
}

func (s *ConfigService) updatePartially(ctx context.Context, older, newest *api.HorusecPlatform) error {
	return s.updateResourceIfNeeded(ctx, older, newest)
}

func (s *ConfigService) list(ctx context.Context) ([]api.HorusecPlatform, error) {
	span, ctx := tracing.StartSpanFromContext(ctx, "internal/business.(*ConfigService).list")
	defer span.Finish()

	cfg, err := s.client.List(ctx, k8s.ListOptions{})
	if err != nil {
		span.SetError(err)
		return nil, fmt.Errorf("failed to list HorusecPlatform: %w", err)
	}
	return cfg.Items, nil
}

func (s *ConfigService) createResource(ctx context.Context, r *api.HorusecPlatform) error {
	span, ctx := tracing.StartSpanFromContext(ctx, "internal/business.(*ConfigService).createResource")
	defer span.Finish()

	r.SetName("horusec")
	_, err := s.client.Create(ctx, r, k8s.CreateOptions{})
	if err != nil {
		span.SetError(err)
		return fmt.Errorf("failed to create HorusecPlatform: %w", err)
	}

	logger.WithPrefix(ctx, "config_service").Debug("resource created")
	return nil
}

func (s *ConfigService) updateResourceIfNeeded(ctx context.Context, older, newest *api.HorusecPlatform) error {
	log := logger.WithPrefix(ctx, "config_service")

	newest.SetName(older.GetName())
	newest.SetResourceVersion(older.GetResourceVersion())
	if diff := s.comparator.Diff(newest, older); diff != "" {
		log.WithField("diff", diff).Debug("resource changed")
		if err := s.updateResource(ctx, newest); err != nil {
			return err
		}
	} else {
		log.Debug("resource unchanged")
	}
	return nil
}

func (s *ConfigService) updateResource(ctx context.Context, r *api.HorusecPlatform) error {
	span, ctx := tracing.StartSpanFromContext(ctx, "internal/business.(*ConfigService).updateResource")
	defer span.Finish()

	_, err := s.client.Update(ctx, r, k8s.UpdateOptions{})
	if err != nil {
		span.SetError(err)
		return fmt.Errorf("failed to update HorusecPlatform: %w", err)
	}

	logger.WithPrefix(ctx, "config_service").Debug("resource updated")
	return nil
}
