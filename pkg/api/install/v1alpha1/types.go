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

package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// HorusecManager defines the installation properties that will be applied.
//
// <!-- go code generation tags
// +genclient
// +k8s:deepcopy-gen=true
// -->
type HorusecManager struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the implementation of this definition.
	// +optional
	Spec HorusecManagerSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`

	Status Status `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// HorusecManagerList is a collection of HorusecManagers.
type HorusecManagerList struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items       []HorusecManager `json:"items" protobuf:"bytes,2,rep,name=items"`
}

type Status struct{}

// HorusecManagerSpec defines the installation properties that will be applied.
//
// <!-- go code generation tags
// +k8s:deepcopy-gen=true
// -->
type HorusecManagerSpec struct {
	Global     *Global     `json:"global,omitempty"`
	Components *Components `json:"components,omitempty"`
}

type JWT struct {
	SecretName string `json:"secretName,omitempty"`
}

type Broker struct {
	Enabled    bool   `json:"enabled,omitempty"`
	Host       string `json:"host,omitempty"`
	Port       int    `json:"port,omitempty"`
	SecretName string `json:"secretName,omitempty"`
}

type Database struct {
	Host       string `json:"host,omitempty"`
	Port       int    `json:"port,omitempty"`
	Dialect    string `json:"dialect,omitempty"`
	SslMode    bool   `json:"sslMode,omitempty"`
	LogMode    bool   `json:"logMode,omitempty"`
	SecretName string `json:"secretName,omitempty"`
}

type Keycloak struct {
	PublicURL   string   `json:"publicURL,omitempty"`
	InternalURL string   `json:"internalURL,omitempty"`
	Realm       string   `json:"realm,omitempty"`
	Otp         bool     `json:"otp,omitempty"`
	Clients     *Clients `json:"clients,omitempty"`
}

type Clients struct {
	Public       *ClientCredentials `json:"public,omitempty"`
	Confidential *ClientCredentials `json:"confidential,omitempty"`
}

type ClientCredentials struct {
	ID     string `json:"id,omitempty"`
	Secret string `json:"secret,omitempty"`
}

type Global struct {
	EnableAdmin bool      `json:"enableAdmin,omitempty"`
	JWT         *JWT      `json:"jwt,omitempty"`
	Broker      *Broker   `json:"broker,omitempty"`
	Database    *Database `json:"database,omitempty"`
	Keycloak    *Keycloak `json:"keycloak,omitempty"`
}

type Ingress struct {
	Enabled   bool     `json:"enabled,omitempty"`
	Scheme    string   `json:"scheme,omitempty"`
	Host      string   `json:"host,omitempty"`
	Protocols []string `json:"protocols,omitempty"`
}

type Account struct {
	Name    string   `json:"name,omitempty"`
	Enabled bool     `json:"enabled,omitempty"`
	Port    *Port    `json:"port,omitempty"`
	Ingress *Ingress `json:"ingress,omitempty"`
}

type Analytic struct {
	Name    string   `json:"name,omitempty"`
	Enabled bool     `json:"enabled,omitempty"`
	Port    *Port    `json:"port,omitempty"`
	Ingress *Ingress `json:"ingress,omitempty"`
}

type API struct {
	Name    string   `json:"name,omitempty"`
	Enabled bool     `json:"enabled,omitempty"`
	Port    *Port    `json:"port,omitempty"`
	Ingress *Ingress `json:"ingress,omitempty"`
}

type Port struct {
	HTTP int `json:"http,omitempty"`
	GRPC int `json:"grpc,omitempty"`
}

type Auth struct {
	Name    string   `json:"name,omitempty"`
	Enabled bool     `json:"enabled,omitempty"`
	Port    *Port    `json:"port,omitempty"`
	Ingress *Ingress `json:"ingress,omitempty"`
}

type Manager struct {
	Name    string   `json:"name,omitempty"`
	Enabled bool     `json:"enabled,omitempty"`
	Port    *Port    `json:"port,omitempty"`
	Ingress *Ingress `json:"ingress,omitempty"`
}

type Components struct {
	Account  *Account  `json:"account,omitempty"`
	Analytic *Analytic `json:"analytic,omitempty"`
	API      *API      `json:"api,omitempty"`
	Auth     *Auth     `json:"auth,omitempty"`
	Manager  *Manager  `json:"manager,omitempty"`
}
