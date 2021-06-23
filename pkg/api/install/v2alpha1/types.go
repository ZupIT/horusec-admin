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

package v2alpha1

import (
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/json"

	"github.com/ZupIT/horusec-admin/pkg/api/install/v2alpha1/state"
)

//
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// HorusecPlatform defines the installation properties that will be applied.
//
// <!-- go code generation tags
// +genclient
// +k8s:deepcopy-gen=true
// -->
type HorusecPlatform struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ObjectMeta `json:"metadata" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the implementation of this definition.
	// +optional
	Spec HorusecPlatformSpec `json:"spec" protobuf:"bytes,2,opt,name=spec"`

	Status HorusecPlatformStatus `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// HorusecPlatformList is a collection of HorusecManagers.
type HorusecPlatformList struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ListMeta `json:"metadata" protobuf:"bytes,1,opt,name=metadata"`
	Items       []HorusecPlatform `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// HorusecPlatformStatus defines the observed state of HorusecPlatform
type HorusecPlatformStatus struct {
	Conditions []v1.Condition `json:"conditions"`
	State      state.Type     `json:"state"`
}

// HorusecPlatformSpec defines the installation properties that will be applied.
//
// <!-- go code generation tags
// +k8s:deepcopy-gen=true
// -->
type HorusecPlatformSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Components Components `json:"components"`
	Global     Global     `json:"global"`
}

func (h *HorusecPlatform) ToBytes() []byte {
	bytes, _ := json.Marshal(h)
	return bytes
}

type Global struct {
	Broker       Broker   `json:"broker"`
	Database     Database `json:"database"`
	JWT          JWT      `json:"jwt"`
	Keycloak     Keycloak `json:"keycloak"`
	Ldap         Ldap     `json:"ldap"`
	GrpcUseCerts bool     `json:"grpcUseCerts"`
}

type Keycloak struct {
	Clients     Clients `json:"clients"`
	InternalURL string  `json:"internalURL"`
	Otp         bool    `json:"otp"`
	PublicURL   string  `json:"publicURL"`
	Realm       string  `json:"realm"`
}

type Ldap struct {
	Base               string           `json:"base"`
	Host               string           `json:"host"`
	Port               int              `json:"port"`
	UseSSL             bool             `json:"useSsl"`
	SkipTLS            bool             `json:"skipTls"`
	InsecureSkipVerify bool             `json:"insecureSkipVerify"`
	BindDN             string           `json:"bindDn"`
	BindPassword       LdapBindPassword `json:"bindPassword"`
	UserFilter         string           `json:"userFilter"`
	AdminGroup         string           `json:"adminGroup"`
}

type LdapBindPassword struct {
	SecretKeyRef *corev1.SecretKeySelector `json:"secretKeyRef" protobuf:"bytes,4,opt,name=secretKeyRef"`
}

type Clients struct {
	Confidential Confidential `json:"confidential"`
	Public       Public       `json:"public"`
}

type Confidential struct {
	ID           string                    `json:"id"`
	SecretKeyRef *corev1.SecretKeySelector `json:"secretKeyRef" protobuf:"bytes,4,opt,name=secretKeyRef"`
}

type Public struct {
	ID string `json:"id"`
}

type JWT struct {
	SecretKeyRef *corev1.SecretKeySelector `json:"secretKeyRef" protobuf:"bytes,4,opt,name=secretKeyRef"`
}

type Broker struct {
	Host        string `json:"host"`
	Port        int    `json:"port"`
	Credentials `json:",inline"`
}

type UserInfo struct {
	Email       string `json:"email"`
	Enabled     bool   `json:"enabled"`
	Credentials `json:",inline"`
}

//nolint:golint, stylecheck // no need to be API in uppercase
type Components struct {
	Analytic      Analytic           `json:"analytic"`
	API           ExposableComponent `json:"api"`
	Auth          Auth               `json:"auth"`
	Core          ExposableComponent `json:"core"`
	Manager       ExposableComponent `json:"manager"`
	Messages      Messages           `json:"messages"`
	Vulnerability ExposableComponent `json:"vulnerability"`
	Webhook       Webhook            `json:"webhook"`
}

type Webhook struct {
	Timeout            int `json:"timeout"`
	ExposableComponent `json:",inline"`
}

type Analytic struct {
	ExposableComponent `json:",inline"`
	Database           Database `json:"database"`
}

type Auth struct {
	Type               AuthType `json:"type"`
	User               User     `json:"user"`
	ExposableComponent `json:",inline"`
}

type User struct {
	Administrator UserInfo `json:"administrator"`
	Default       UserInfo `json:"default"`
}

type AuthType string

type Messages struct {
	Enabled            bool       `json:"enabled"`
	MailServer         MailServer `json:"mailServer"`
	EmailFrom          string     `json:"emailFrom"`
	ExposableComponent `json:",inline"`
}

type Container struct {
	Image           Image                       `json:"image"`
	LivenessProbe   corev1.Probe                `json:"livenessProbe"`
	ReadinessProbe  corev1.Probe                `json:"readinessProbe"`
	Resources       corev1.ResourceRequirements `json:"resources"`
	SecurityContext ContainerSecurityContext    `json:"securityContext"`
}

type Image struct {
	PullPolicy  string   `json:"pullPolicy"`
	PullSecrets []string `json:"pullSecrets"`
	Registry    string   `json:"registry"`
	Repository  string   `json:"repository"`
	Tag         string   `json:"tag"`
}

type ContainerSecurityContext struct {
	Enabled                bool `json:"enabled"`
	corev1.SecurityContext `json:",inline"`
}

type PodSecurityContext struct {
	Enabled                   bool `json:"enabled"`
	corev1.PodSecurityContext `json:",inline"`
}

type Database struct {
	Host        string    `json:"host"`
	LogMode     bool      `json:"logMode"`
	Name        string    `json:"name"`
	Port        int       `json:"port"`
	SslMode     *bool     `json:"sslMode"`
	Migration   Migration `json:"migration"`
	Credentials `json:",inline"`
}

type Migration struct {
	Image Image `json:"image"`
}

type Ingress struct {
	Enabled *bool  `json:"enabled"`
	Host    string `json:"host"`
	Path    string `json:"path"`
	TLS     TLS    `json:"tls"`
}

type TLS struct {
	SecretName string `json:"secretName"`
}

type Pod struct {
	Autoscaling     Autoscaling        `json:"autoscaling"`
	SecurityContext PodSecurityContext `json:"securityContext"`
}

type Autoscaling struct {
	Enabled      bool   `json:"enabled"`
	MaxReplicas  int32  `json:"maxReplicas"`
	MinReplicas  *int32 `json:"minReplicas"`
	TargetCPU    *int32 `json:"targetCPU"`
	TargetMemory *int32 `json:"targetMemory"`
}

type Ports struct {
	HTTP int `json:"http"`
	GRPC int `json:"grpc"`
}

type MailServer struct {
	Host        string `json:"host"`
	Port        int    `json:"port"`
	Credentials `json:",inline"`
}

type Credentials struct {
	User     SecretRef `json:"user"`
	Password SecretRef `json:"password"`
}

type SecretRef struct {
	KeyRef *corev1.SecretKeySelector `json:"secretKeyRef" protobuf:"bytes,4,opt,name=secretKeyRef"`
}

type Component struct {
	Name         string          `json:"name"`
	Port         Ports           `json:"port"`
	ExtraEnv     []corev1.EnvVar `json:"extraEnv"`
	ReplicaCount int32           `json:"replicaCount"`
	Pod          Pod             `json:"pod"`
	Container    Container       `json:"container"`
}

type ExposableComponent struct {
	Component `json:",inline"`
	Ingress   Ingress `json:"ingress"`
}
