package business

import (
	"context"
	"errors"
	"fmt"

	"github.com/ZupIT/horusec-admin/internal/logger"
	api "github.com/ZupIT/horusec-admin/pkg/api/install/v1alpha1"
	clientset "github.com/ZupIT/horusec-admin/pkg/client/clientset/versioned/typed/install/v1alpha1"
	"github.com/ZupIT/horusec-admin/pkg/core"
	"github.com/google/go-cmp/cmp"
	kerrors "k8s.io/apimachinery/pkg/api/errors"
	k8s "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ConfigService struct {
	client      clientset.HorusecManagerInterface
	compareOpts cmp.Option
}

func NewConfigService(client clientset.HorusecManagerInterface) *ConfigService {
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

	return core.NewConfiguration(cr), nil
}

func (s *ConfigService) Update(cfg *core.Configuration) error {
	r2, err := cfg.CR()
	if err != nil {
		return err
	}

	r1, err := s.getOne()
	if err != nil {
		return err
	}

	if r1 != nil {
		r2.Name = r1.Name
	} else {
		r2.Name = "horusec"
	}

	err = s.apply(r2)
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

	o, err := s.client.Get(context.TODO(), r.Name, k8s.GetOptions{})
	if kerrors.IsNotFound(err) {
		_, err = s.client.Create(context.TODO(), r, k8s.CreateOptions{})
		if err != nil {
			return err
		}
		log.Debug("resource created")
		return nil
	} else if err != nil {
		return err
	}

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
