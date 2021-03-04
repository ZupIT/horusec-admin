package core

type Manager struct {
	APIEndpoint      string `json:"react_app_horusec_endpoint_api"`
	AnalyticEndpoint string `json:"react_app_horusec_endpoint_analytic"`
	AccountEndpoint  string `json:"react_app_horusec_endpoint_account"`
	AuthEndpoint     string `json:"react_app_horusec_endpoint_auth"`
	ManagerPath      string `json:"react_app_horusec_manager_path"`
}
