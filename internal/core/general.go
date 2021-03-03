package core

type General struct {
	EnableApplicationAdmin string `json:"horusec_enable_application_admin,omitempty"`
	JwtSecretKey           string `json:"horusec_jwt_secret_key,omitempty"`
	ApplicationAdminData   string `json:"horusec_application_admin_data,omitempty"`
}
