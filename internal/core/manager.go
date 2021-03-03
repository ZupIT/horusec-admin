package core

type Manager struct {
	EndpointAPI      string `json:"react_app_horusec_endpoint_api,omitempty"`
	EndpointAnalytic string `json:"react_app_horusec_endpoint_analytic,omitempty"`
	EndpointAccount  string `json:"react_app_horusec_endpoint_account,omitempty"`
	EndpointAuth     string `json:"react_app_horusec_endpoint_auth,omitempty"`
	ManagerPath      string `json:"react_app_horusec_manager_path,omitempty"`
}
