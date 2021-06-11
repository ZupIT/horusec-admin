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

// HorusecPlatform defines the installation properties that will be applied.
//
// <!-- go code generation tags
// +genclient
// +k8s:deepcopy-gen=true
// -->
type HorusecPlatform struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec defines the implementation of this definition.
	// +optional
	Spec HorusecPlatformSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`

	Status Status `json:"status"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// HorusecManagerList is a collection of HorusecManagers.
type HorusecManagerList struct {
	v1.TypeMeta `json:",inline"`
	// +optional
	v1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`
	Items       []HorusecPlatform `json:"items" protobuf:"bytes,2,rep,name=items"`
}

type Status struct{}

// HorusecPlatformSpec defines the installation properties that will be applied.
//
// <!-- go code generation tags
// +k8s:deepcopy-gen=true
// -->
type HorusecPlatformSpec struct {
	Components Components `json:"components"`
	Global     Global     `json:"global"`
}

type Components struct {
	Analytic      Analytic `json:"analytic"`
	API           Analytic `json:"api"`
	Auth          Auth     `json:"auth"`
	Core          Analytic `json:"core"`
	Manager       Analytic `json:"manager"`
	Messages      Analytic `json:"messages"`
	Vulnerability Analytic `json:"vulnerability"`
	Webhook       Analytic `json:"webhook"`
}

type Analytic struct {
	Container    Container     `json:"container"`
	Database     *Database     `json:"database,omitempty"`
	ExtraEnv     []interface{} `json:"extraEnv"`
	Ingress      *Ingress      `json:"ingress,omitempty"`
	Name         string        `json:"name"`
	Pod          Pod           `json:"pod"`
	Port         AnalyticPort  `json:"port"`
	ReplicaCount int64         `json:"replicaCount"`
	Enabled      *bool         `json:"enabled,omitempty"`
	MailServer   *Broker       `json:"mailServer,omitempty"`
}

type Container struct {
	Image           Image                    `json:"image"`
	LivenessProbe   interface{}              `json:"livenessProbe"`
	ReadinessProbe  interface{}              `json:"readinessProbe"`
	Resources       interface{}              `json:"resources"`
	SecurityContext ContainerSecurityContext `json:"securityContext"`
}

type Image struct {
	PullPolicy  PullPolicy    `json:"pullPolicy"`
	PullSecrets []interface{} `json:"pullSecrets"`
	Registry    Registry      `json:"registry"`
	Repository  string        `json:"repository"`
	Tag         Tag           `json:"tag"`
}

type ContainerSecurityContext struct {
	Enabled      bool  `json:"enabled"`
	RunAsNonRoot bool  `json:"runAsNonRoot"`
	RunAsUser    int64 `json:"runAsUser"`
}

type Database struct {
	Host      string    `json:"host"`
	LogMode   bool      `json:"logMode"`
	Migration Migration `json:"migration"`
	Name      string    `json:"name"`
	Password  Jwt       `json:"password"`
	Port      int64     `json:"port"`
	SSLMode   bool      `json:"sslMode"`
	User      Jwt       `json:"user"`
}

type Migration struct {
	Image Image `json:"image"`
}

type Jwt struct {
	SecretKeyRef SecretKeyRef `json:"secretKeyRef"`
}

type SecretKeyRef struct {
	Key  string `json:"key"`
	Name string `json:"name"`
}

type Ingress struct {
	Enabled bool        `json:"enabled"`
	Host    string      `json:"host"`
	Path    string      `json:"path"`
	TLS     interface{} `json:"tls"`
}

type Broker struct {
	Host     string `json:"host"`
	Password Jwt    `json:"password"`
	Port     int64  `json:"port"`
	User     Jwt    `json:"user"`
}

type Pod struct {
	Autoscaling     Autoscaling        `json:"autoscaling"`
	SecurityContext PodSecurityContext `json:"securityContext"`
}

type Autoscaling struct {
	Enabled      bool  `json:"enabled"`
	MaxReplicas  int64 `json:"maxReplicas"`
	MinReplicas  int64 `json:"minReplicas"`
	TargetCPU    int64 `json:"targetCPU"`
	TargetMemory int64 `json:"targetMemory"`
}

type PodSecurityContext struct {
	Enabled bool  `json:"enabled"`
	FSGroup int64 `json:"fsGroup"`
}

type AnalyticPort struct {
	HTTP int64 `json:"http"`
}

type Auth struct {
	Container    Container     `json:"container"`
	ExtraEnv     []interface{} `json:"extraEnv"`
	Ingress      Ingress       `json:"ingress"`
	Name         string        `json:"name"`
	Pod          Pod           `json:"pod"`
	Port         AuthPort      `json:"port"`
	ReplicaCount int64         `json:"replicaCount"`
	Type         string        `json:"type"`
}

type AuthPort struct {
	Grpc int64 `json:"grpc"`
	HTTP int64 `json:"http"`
}

type Global struct {
	Administrator Administrator `json:"administrator"`
	Broker        Broker        `json:"broker"`
	Database      Database      `json:"database"`
	Jwt           Jwt           `json:"jwt"`
	Keycloak      Keycloak      `json:"keycloak"`
}

type Administrator struct {
	Email    string `json:"email"`
	Enabled  bool   `json:"enabled"`
	Password Jwt    `json:"password"`
	User     Jwt    `json:"user"`
}

type Keycloak struct {
	Clients     Clients `json:"clients"`
	InternalURL string  `json:"internalURL"`
	Otp         bool    `json:"otp"`
	PublicURL   string  `json:"publicURL"`
	Realm       string  `json:"realm"`
}

type Clients struct {
	Confidential Confidential `json:"confidential"`
	Public       Public       `json:"public"`
}

type Confidential struct {
	ID           string       `json:"id"`
	SecretKeyRef SecretKeyRef `json:"secretKeyRef"`
}

type Public struct {
	ID string `json:"id"`
}

type PullPolicy string
const (
	IfNotPresent PullPolicy = "IfNotPresent"
)

type Registry string
const (
	DockerIoHoruszup Registry = "docker.io/horuszup"
)

type Tag string
const (
	V2121 Tag = "v2.13.1"
)
