package core

import (
	"github.com/ZupIT/horusec-admin/pkg/api/install/v1alpha1"
)

type Configuration struct {
	*General `json:",inline,omitempty"`
	*Auth    `json:",inline,omitempty"`
	*Manager `json:",inline,omitempty"`
}

func newConfiguration(cr *v1alpha1.HorusecManager) *Configuration {
	return &Configuration{
		General: newGeneral(cr),
		Auth:    newAuth(cr),
		Manager: newManager(cr),
	}
}

func (c *Configuration) toCR() (*v1alpha1.HorusecManager, error) {
	cr, err := newCR(c)
	if err != nil {
		return nil, err
	}

	return (*v1alpha1.HorusecManager)(cr), nil
}
