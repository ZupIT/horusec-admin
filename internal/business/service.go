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
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	tlog "github.com/opentracing/opentracing-go/log"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	k8s "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ConfigService struct {
	client      client.HorusecManagerInterface
	compareOpts cmp.Option
}

func NewConfigService(client client.HorusecManagerInterface) *ConfigService {
	ignore := [...]string{
		"TypeMeta.APIVersion",
		"TypeMeta.Kind",
		"ObjectMeta.CreationTimestamp",
		"ObjectMeta.Finalizers",
		"ObjectMeta.Generation",
		"ObjectMeta.ManagedFields",
		"ObjectMeta.Namespace",
		"ObjectMeta.ResourceVersion",
		"ObjectMeta.SelfLink",
		"ObjectMeta.UID",
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

func (s *ConfigService) GetConfig(ctx context.Context) (*core.Configuration, error) {
	cr, err := s.getOne(ctx)
	if err != nil {
		return nil, err
	}

	if cr == nil {
		return new(core.Configuration), nil
	}

	cfg := adapter.NewConfiguration(cr)
	return (*core.Configuration)(cfg), nil
}

func (s *ConfigService) CreateOrUpdate(ctx context.Context, cfg *core.Configuration) error {
	r, err := (*adapter.Configuration)(cfg).CR()
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
	log := logger.WithPrefix(ctx, "service")

	hm, err := s.list(ctx)
	if err != nil {
		return nil, err
	}

	if len(hm) > 1 {
		return nil, errors.New("more than one HorusecManager CR were found")
	}

	if len(hm) == 0 {
		log.Debug("no HorusecManager was found")
		return nil, nil
	}

	res := &hm[0]
	log.WithField("name", res.Name).WithField("namespace", res.Namespace).Debug("a HorusecManager was found")
	return res, nil
}

func (s *ConfigService) apply(ctx context.Context, r *api.HorusecManager) error {
	o, err := s.getOne(ctx)
	if err != nil {
		return err
	}

	log := logger.WithPrefix(ctx, "service")
	if o == nil {
		r.SetName("horusec")
		if err = s.create(ctx, r); err != nil {
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
		if err = s.update(ctx, r); err != nil {
			return err
		}
		log.Debug("resource updated")
	} else {
		log.Debug("resource unchanged")
	}
	return nil
}

func (s *ConfigService) list(ctx context.Context) ([]api.HorusecManager, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "internal/business.(*ConfigService).list")
	defer span.Finish()

	cfg, err := s.client.List(ctx, k8s.ListOptions{})
	if err != nil {
		ext.Error.Set(span, true)
		fields := make([]tlog.Field, 0)
		fields = append(fields, tlog.String("event", "error"), tlog.String("message", err.Error()))
		if serr, ok := err.(*kerrors.StatusError); ok {
			status := serr.Status()
			fields = append(fields, tlog.Int32("code", status.Code), tlog.String("reason", string(status.Reason)))
		}
		span.LogFields(fields...)
		return nil, fmt.Errorf("failed to list HorusecManager: %w", err)
	}
	return cfg.Items, nil
}

func (s *ConfigService) create(ctx context.Context, r *api.HorusecManager) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "internal/business.(*ConfigService).create")
	defer span.Finish()

	_, err := s.client.Create(ctx, r, k8s.CreateOptions{})
	if err != nil {
		ext.Error.Set(span, true)
		fields := make([]tlog.Field, 0)
		fields = append(fields, tlog.String("event", "error"), tlog.String("message", err.Error()))
		if serr, ok := err.(*kerrors.StatusError); ok {
			status := serr.Status()
			fields = append(fields, tlog.Int32("code", status.Code), tlog.String("reason", string(status.Reason)))
		}
		span.LogFields(fields...)
		return fmt.Errorf("failed to create HorusecManager: %w", err)
	}
	return nil
}

func (s *ConfigService) update(ctx context.Context, r *api.HorusecManager) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "internal/business.(*ConfigService).update")
	defer span.Finish()

	_, err := s.client.Update(ctx, r, k8s.UpdateOptions{})
	if err != nil {
		ext.Error.Set(span, true)
		fields := make([]tlog.Field, 0)
		fields = append(fields, tlog.String("event", "error"), tlog.String("message", err.Error()))
		if serr, ok := err.(*kerrors.StatusError); ok {
			status := serr.Status()
			fields = append(fields, tlog.Int32("code", status.Code), tlog.String("reason", string(status.Reason)))
		}
		span.LogFields(fields...)
		return fmt.Errorf("failed to update HorusecManager: %w", err)
	}
	return nil
}
