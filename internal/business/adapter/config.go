package adapter

import (
	api "github.com/ZupIT/horusec-admin/pkg/api/install/v1alpha1"
	"github.com/ZupIT/horusec-admin/pkg/core"
)

type Configuration core.Configuration

func NewConfiguration(cr *api.HorusecManager) *Configuration {
	return &Configuration{
		General: new(core.General),
		Auth:    new(core.Auth),
		Manager: newManager(cr),
	}
}

func (c *Configuration) CR() (*api.HorusecManager, error) {
	cr, err := NewCustomResource(c)
	if err != nil {
		return nil, err
	}

	return (*api.HorusecManager)(cr), nil
}
