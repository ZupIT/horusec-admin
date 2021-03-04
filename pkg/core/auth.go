package core

import "github.com/ZupIT/horusec-admin/pkg/api/install/v1alpha1"

type (
	Auth struct {
		Type      string `json:"horusec_auth_type,omitempty"`
		*Keycloak `json:",inline,omitempty"`
		*LDAP     `json:",inline,omitempty"`
	}
	Keycloak struct {
		BasePath          string `json:"horusec_keycloak_base_path,omitempty"`
		ClientID          string `json:"horusec_keycloak_client_id,omitempty"`
		ClientSecret      string `json:"horusec_keycloak_client_secret,omitempty"`
		Realm             string `json:"horusec_keycloak_realm,omitempty"`
		OTP               string `json:"horusec_keycloak_otp,omitempty"`
		*KeycloakReactApp `json:",inline,omitempty"`
	}
	KeycloakReactApp struct {
		ClientID string `json:"react_app_keycloak_client_id,omitempty"`
		Realm    string `json:"react_app_keycloak_realm,omitempty"`
		BasePath string `json:"react_app_keycloak_base_path,omitempty"`
	}
	LDAP struct {
		Base               string `json:"horusec_ldap_base,omitempty"`
		Host               string `json:"horusec_ldap_host,omitempty"`
		Port               string `json:"horusec_ldap_port,omitempty"`
		UseSSL             string `json:"horusec_ldap_usessl,omitempty"`
		SkipTLS            string `json:"horusec_ldap_skip_tls,omitempty"`
		InsecureSkipVerify string `json:"horusec_ldap_insecure_skip_verify,omitempty"`
		BindDN             string `json:"horusec_ldap_binddn,omitempty"`
		BindPassword       string `json:"horusec_ldap_bindpassword,omitempty"`
		UserFilter         string `json:"horusec_ldap_userfilter,omitempty"`
		GroupFilter        string `json:"horusec_ldap_groupfilter,omitempty"`
		AdminGroup         string `json:"horusec_ldap_admin_group,omitempty"`
	}
)

func newAuth(cr *v1alpha1.HorusecManager) *Auth {
	return &Auth{}
}
