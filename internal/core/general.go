package core

import "github.com/ZupIT/horusec-admin/pkg/api/install/v1alpha1"

type General struct {
	EnableApplicationAdmin bool   `json:"horusec_enable_application_admin,omitempty"`
	JwtSecretKey           string `json:"horusec_jwt_secret_key,omitempty"`
	ApplicationAdminData   string `json:"horusec_application_admin_data,omitempty"`
}

func newGeneral(cr *v1alpha1.HorusecManager) *General {
	return &General{}
}
