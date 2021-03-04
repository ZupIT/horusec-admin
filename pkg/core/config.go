package core

import api "github.com/ZupIT/horusec-admin/pkg/api/install/v1alpha1"

type Configuration struct {
	*General `json:",inline,omitempty"`
	*Auth    `json:",inline,omitempty"`
	*Manager `json:",inline,omitempty"`
}

func NewConfiguration(cr *api.HorusecManager) *Configuration {
	return &Configuration{
		General: newGeneral(cr),
		Auth:    newAuth(cr),
		Manager: newManager(cr),
	}
}

func (c *Configuration) CR() (*api.HorusecManager, error) {
	cr, err := newCR(c)
	if err != nil {
		return nil, err
	}

	return (*api.HorusecManager)(cr), nil
}
