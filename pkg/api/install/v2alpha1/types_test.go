package v2alpha1

import (
	"testing"

	"github.com/stretchr/testify/assert"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
	"k8s.io/apimachinery/pkg/util/json"
)

func Test(test *testing.T) {
	var f = false
	var t = true
	entity := HorusecPlatformSpec{
		Components: Components{
			Analytic: Analytic{
				ExposableComponent: ExposableComponent{
					Component: Component{
						Name: "analytic",
						Port: Ports{
							HTTP: 8005,
						},
						ExtraEnv:     nil,
						ReplicaCount: 0,
						Pod: Pod{
							Autoscaling: Autoscaling{
								Enabled:      false,
								MaxReplicas:  0,
								MinReplicas:  nil,
								TargetCPU:    nil,
								TargetMemory: nil,
							},
							SecurityContext: PodSecurityContext{
								Enabled:            false,
								PodSecurityContext: corev1.PodSecurityContext{},
							},
						},
						Container: Container{
							Image: Image{
								Registry: "docker.io/horuszup/horusec-analytic:v2.13.1-alpha.1",
							},
							LivenessProbe: corev1.Probe{
								Handler: corev1.Handler{
									HTTPGet: &corev1.HTTPGetAction{
										Path: "/analytic/health",
										Port: intstr.IntOrString{Type: intstr.String, StrVal: "http"},
									},
								},
							},
							ReadinessProbe: corev1.Probe{
								Handler: corev1.Handler{
									HTTPGet: &corev1.HTTPGetAction{
										Path: "/analytic/health",
										Port: intstr.IntOrString{Type: intstr.String, StrVal: "http"},
									},
								},
							},
							Resources:       corev1.ResourceRequirements{},
							SecurityContext: ContainerSecurityContext{},
						},
					},
					Ingress: Ingress{
						Host: "analytic.local",
						Path: "",
					},
				},
				Database: Database{
					Host:    "postgresql",
					LogMode: false,
					Name:    "horusec_analytic_db",
					Port:    5432,
					SslMode: &f,
					Credentials: Credentials{
						User: SecretRef{
							KeyRef: &corev1.SecretKeySelector{
								LocalObjectReference: corev1.LocalObjectReference{
									Name: "username",
								},
								Key:      "horusec-analytic-database",
								Optional: &f,
							},
						},
						Password: SecretRef{
							KeyRef: &corev1.SecretKeySelector{
								LocalObjectReference: corev1.LocalObjectReference{
									Name: "password",
								},
								Key:      "horusec-analytic-database",
								Optional: &f,
							},
						},
					},
					Migration: Migration{},
				},
			},
			API: ExposableComponent{
				Component: Component{
					Name: "api",
					Port: Ports{
						HTTP: 8000,
					},
					ExtraEnv:     nil,
					ReplicaCount: 0,
					Pod: Pod{
						Autoscaling: Autoscaling{
							Enabled:      false,
							MaxReplicas:  0,
							MinReplicas:  nil,
							TargetCPU:    nil,
							TargetMemory: nil,
						},
						SecurityContext: PodSecurityContext{
							Enabled:            false,
							PodSecurityContext: corev1.PodSecurityContext{},
						},
					},
					Container: Container{
						Image: Image{
							Registry: "docker.io/horuszup/horusec-api:v2.13.1-alpha.1",
						},
						LivenessProbe: corev1.Probe{
							Handler: corev1.Handler{
								HTTPGet: &corev1.HTTPGetAction{
									Path: "/api/health",
									Port: intstr.IntOrString{Type: intstr.String, StrVal: "http"},
								},
							},
						},
						ReadinessProbe: corev1.Probe{
							Handler: corev1.Handler{
								HTTPGet: &corev1.HTTPGetAction{
									Path: "/api/health",
									Port: intstr.IntOrString{Type: intstr.String, StrVal: "http"},
								},
							},
						},
						Resources:       corev1.ResourceRequirements{},
						SecurityContext: ContainerSecurityContext{},
					},
				},
				Ingress: Ingress{
					Host: "api.local",
					Path: "",
				},
			},
			Auth: Auth{
				Type: "horusec",
				User: User{
					Administrator: UserInfo{
						Email:   "admin-user@horusec.io",
						Enabled: false,
						Credentials: Credentials{
							User: SecretRef{
								KeyRef: &corev1.SecretKeySelector{
									LocalObjectReference: corev1.LocalObjectReference{
										Name: "username",
									},
									Key:      "horusec-platform-auth-user-administrator",
									Optional: &t,
								},
							},
							Password: SecretRef{
								KeyRef: &corev1.SecretKeySelector{
									LocalObjectReference: corev1.LocalObjectReference{
										Name: "password",
									},
									Key:      "horusec-platform-auth-user-administrator",
									Optional: &t,
								},
							},
						},
					},
					Default: UserInfo{
						Email:   "user@horusec.io",
						Enabled: true,
						Credentials: Credentials{
							User: SecretRef{
								KeyRef: &corev1.SecretKeySelector{
									LocalObjectReference: corev1.LocalObjectReference{
										Name: "username",
									},
									Key:      "horusec-platform-auth-user-default",
									Optional: &f,
								},
							},
							Password: SecretRef{
								KeyRef: &corev1.SecretKeySelector{
									LocalObjectReference: corev1.LocalObjectReference{
										Name: "password",
									},
									Key:      "horusec-platform-auth-user-default",
									Optional: &f,
								},
							},
						},
					},
				},
				ExposableComponent: ExposableComponent{
					Component: Component{
						Name: "auth",
						Port: Ports{
							HTTP: 8006,
							GRPC: 8007,
						},
						ExtraEnv:     nil,
						ReplicaCount: 0,
						Pod: Pod{
							Autoscaling: Autoscaling{
								Enabled:      false,
								MaxReplicas:  0,
								MinReplicas:  nil,
								TargetCPU:    nil,
								TargetMemory: nil,
							},
							SecurityContext: PodSecurityContext{
								Enabled:            false,
								PodSecurityContext: corev1.PodSecurityContext{},
							},
						},
						Container: Container{
							Image: Image{
								Registry: "docker.io/horuszup/horusec-auth:v2.13.1-alpha.1",
							},
							LivenessProbe: corev1.Probe{
								Handler: corev1.Handler{
									HTTPGet: &corev1.HTTPGetAction{
										Path: "/auth/health",
										Port: intstr.IntOrString{Type: intstr.String, StrVal: "http"},
									},
								},
							},
							ReadinessProbe: corev1.Probe{
								Handler: corev1.Handler{
									HTTPGet: &corev1.HTTPGetAction{
										Path: "/auth/health",
										Port: intstr.IntOrString{Type: intstr.String, StrVal: "http"},
									},
								},
							},
							Resources:       corev1.ResourceRequirements{},
							SecurityContext: ContainerSecurityContext{},
						},
					},
					Ingress: Ingress{
						Host: "auth.local",
						Path: "",
					},
				},
			},
			Core: ExposableComponent{
				Component: Component{
					Name: "core",
					Port: Ports{
						HTTP: 8003,
					},
					ExtraEnv:     nil,
					ReplicaCount: 0,
					Pod: Pod{
						Autoscaling: Autoscaling{
							Enabled:      false,
							MaxReplicas:  0,
							MinReplicas:  nil,
							TargetCPU:    nil,
							TargetMemory: nil,
						},
						SecurityContext: PodSecurityContext{
							Enabled:            false,
							PodSecurityContext: corev1.PodSecurityContext{},
						},
					},
					Container: Container{
						Image: Image{
							Registry: "docker.io/horuszup/horusec-core:v2.13.1-alpha.1",
						},
						LivenessProbe: corev1.Probe{
							Handler: corev1.Handler{
								HTTPGet: &corev1.HTTPGetAction{
									Path: "/core/health",
									Port: intstr.IntOrString{Type: intstr.String, StrVal: "http"},
								},
							},
						},
						ReadinessProbe: corev1.Probe{
							Handler: corev1.Handler{
								HTTPGet: &corev1.HTTPGetAction{
									Path: "/core/health",
									Port: intstr.IntOrString{Type: intstr.String, StrVal: "http"},
								},
							},
						},
						Resources:       corev1.ResourceRequirements{},
						SecurityContext: ContainerSecurityContext{},
					},
				},
				Ingress: Ingress{
					Host: "core.local",
					Path: "",
				},
			},
			Manager: ExposableComponent{
				Component: Component{
					Name: "manager",
					Port: Ports{
						HTTP: 8043,
					},
					ExtraEnv:     nil,
					ReplicaCount: 0,
					Pod: Pod{
						Autoscaling: Autoscaling{
							Enabled:      false,
							MaxReplicas:  0,
							MinReplicas:  nil,
							TargetCPU:    nil,
							TargetMemory: nil,
						},
						SecurityContext: PodSecurityContext{
							Enabled:            false,
							PodSecurityContext: corev1.PodSecurityContext{},
						},
					},
					Container: Container{
						Image: Image{
							Registry: "docker.io/horuszup/horusec-manager:v2.13.1-alpha.1",
						},
						LivenessProbe: corev1.Probe{
							Handler: corev1.Handler{
								HTTPGet: &corev1.HTTPGetAction{
									Path: "/manager/health",
									Port: intstr.IntOrString{Type: intstr.String, StrVal: "http"},
								},
							},
						},
						ReadinessProbe: corev1.Probe{
							Handler: corev1.Handler{
								HTTPGet: &corev1.HTTPGetAction{
									Path: "/manager/health",
									Port: intstr.IntOrString{Type: intstr.String, StrVal: "http"},
								},
							},
						},
						Resources:       corev1.ResourceRequirements{},
						SecurityContext: ContainerSecurityContext{},
					},
				},
				Ingress: Ingress{
					Host: "manager.local",
					Path: "",
				},
			},
			Messages: Messages{
				Enabled: false,
				MailServer: MailServer{
					Host: "smtp.mailtrap.io",
					Port: 443,
					Credentials: Credentials{
						User: SecretRef{
							KeyRef: &corev1.SecretKeySelector{
								LocalObjectReference: corev1.LocalObjectReference{
									Name: "username",
								},
								Key:      "horusec-platform-smtp",
								Optional: &f,
							},
						},
						Password: SecretRef{
							KeyRef: &corev1.SecretKeySelector{
								LocalObjectReference: corev1.LocalObjectReference{
									Name: "password",
								},
								Key:      "horusec-platform-smtp",
								Optional: &f,
							},
						},
					},
				},
				ExposableComponent: ExposableComponent{
					Component: Component{
						Name: "messages",
						Port: Ports{
							HTTP: 8002,
						},
						ExtraEnv:     nil,
						ReplicaCount: 0,
						Pod: Pod{
							Autoscaling: Autoscaling{
								Enabled:      false,
								MaxReplicas:  0,
								MinReplicas:  nil,
								TargetCPU:    nil,
								TargetMemory: nil,
							},
							SecurityContext: PodSecurityContext{
								Enabled:            false,
								PodSecurityContext: corev1.PodSecurityContext{},
							},
						},
						Container: Container{
							Image: Image{
								Registry: "docker.io/horuszup/horusec-messages:v2.13.1-alpha.1",
							},
							LivenessProbe: corev1.Probe{
								Handler: corev1.Handler{
									HTTPGet: &corev1.HTTPGetAction{
										Path: "/messages/health",
										Port: intstr.IntOrString{Type: intstr.String, StrVal: "http"},
									},
								},
							},
							ReadinessProbe: corev1.Probe{
								Handler: corev1.Handler{
									HTTPGet: &corev1.HTTPGetAction{
										Path: "/messages/health",
										Port: intstr.IntOrString{Type: intstr.String, StrVal: "http"},
									},
								},
							},
							Resources:       corev1.ResourceRequirements{},
							SecurityContext: ContainerSecurityContext{},
						},
					},
					Ingress: Ingress{
						Host: "messages.local",
						Path: "",
					},
				},
			},
			Vulnerability: ExposableComponent{
				Component: Component{
					Name: "vulnerability",
					Port: Ports{
						HTTP: 8001,
					},
					ExtraEnv:     nil,
					ReplicaCount: 0,
					Pod: Pod{
						Autoscaling: Autoscaling{
							Enabled:      false,
							MaxReplicas:  0,
							MinReplicas:  nil,
							TargetCPU:    nil,
							TargetMemory: nil,
						},
						SecurityContext: PodSecurityContext{
							Enabled:            false,
							PodSecurityContext: corev1.PodSecurityContext{},
						},
					},
					Container: Container{
						Image: Image{
							Registry: "docker.io/horuszup/horusec-vulnerability:v2.13.1-alpha.1",
						},
						LivenessProbe: corev1.Probe{
							Handler: corev1.Handler{
								HTTPGet: &corev1.HTTPGetAction{
									Path: "/vulnerability/health",
									Port: intstr.IntOrString{Type: intstr.String, StrVal: "http"},
								},
							},
						},
						ReadinessProbe: corev1.Probe{
							Handler: corev1.Handler{
								HTTPGet: &corev1.HTTPGetAction{
									Path: "/vulnerability/health",
									Port: intstr.IntOrString{Type: intstr.String, StrVal: "http"},
								},
							},
						},
						Resources:       corev1.ResourceRequirements{},
						SecurityContext: ContainerSecurityContext{},
					},
				},
				Ingress: Ingress{
					Host: "vulnerability.local",
					Path: "",
				},
			},
			Webhook: Webhook{
				Timeout: 60,
				ExposableComponent: ExposableComponent{
					Component: Component{
						Name: "webhook",
						Port: Ports{
							HTTP: 8004,
						},
						ExtraEnv:     nil,
						ReplicaCount: 0,
						Pod: Pod{
							Autoscaling: Autoscaling{
								Enabled:      false,
								MaxReplicas:  0,
								MinReplicas:  nil,
								TargetCPU:    nil,
								TargetMemory: nil,
							},
							SecurityContext: PodSecurityContext{
								Enabled:            false,
								PodSecurityContext: corev1.PodSecurityContext{},
							},
						},
						Container: Container{
							Image: Image{
								Registry: "docker.io/horuszup/horusec-webhook:v2.13.1-alpha.1",
							},
							LivenessProbe: corev1.Probe{
								Handler: corev1.Handler{
									HTTPGet: &corev1.HTTPGetAction{
										Path: "/webhook/health",
										Port: intstr.IntOrString{Type: intstr.String, StrVal: "http"},
									},
								},
							},
							ReadinessProbe: corev1.Probe{
								Handler: corev1.Handler{
									HTTPGet: &corev1.HTTPGetAction{
										Path: "/webhook/health",
										Port: intstr.IntOrString{Type: intstr.String, StrVal: "http"},
									},
								},
							},
							Resources:       corev1.ResourceRequirements{},
							SecurityContext: ContainerSecurityContext{},
						},
					},
					Ingress: Ingress{
						Host: "webhook.local",
						Path: "",
					},
				},
			},
		},
		Global: Global{
			Broker: Broker{
				Host: "rabbitmq",
				Port: 5672,
				Credentials: Credentials{
					User: SecretRef{
						KeyRef: &corev1.SecretKeySelector{
							LocalObjectReference: corev1.LocalObjectReference{
								Name: "username",
							},
							Key:      "horusec-platform-broker",
							Optional: &f,
						},
					},
					Password: SecretRef{
						KeyRef: &corev1.SecretKeySelector{
							LocalObjectReference: corev1.LocalObjectReference{
								Name: "password",
							},
							Key:      "horusec-platform-broker",
							Optional: &f,
						},
					},
				},
			},
			Database: Database{
				Host:      "postgresql",
				LogMode:   false,
				Name:      "horusec_db",
				Port:      5432,
				SslMode:   &f,
				Migration: Migration{},
				Credentials: Credentials{
					User: SecretRef{
						KeyRef: &corev1.SecretKeySelector{
							LocalObjectReference: corev1.LocalObjectReference{
								Name: "username",
							},
							Key:      "horusec-platform-database",
							Optional: &f,
						},
					},
					Password: SecretRef{
						KeyRef: &corev1.SecretKeySelector{
							LocalObjectReference: corev1.LocalObjectReference{
								Name: "password",
							},
							Key:      "horusec-platform-database",
							Optional: &f,
						},
					},
				},
			},
			JWT: JWT{
				SecretKeyRef: &corev1.SecretKeySelector{
					LocalObjectReference: corev1.LocalObjectReference{
						Name: "secret",
					},
					Key:      "horusec-platform-jwt",
					Optional: &f,
				},
			},
			Keycloak: Keycloak{
				Clients: Clients{
					Confidential: Confidential{
						ID: "",
						SecretKeyRef: &corev1.SecretKeySelector{
							LocalObjectReference: corev1.LocalObjectReference{
								Name: "secret",
							},
							Key:      "horusec-platform-keycloak",
							Optional: &t,
						},
					},
					Public: Public{
						ID: GroupName,
					},
				},
				InternalURL: "",
				Otp:         false,
				PublicURL:   "",
				Realm:       "",
			},
			Ldap: Ldap{
				Base:               "",
				Host:               "",
				Port:               0,
				UseSSL:             false,
				SkipTLS:            false,
				InsecureSkipVerify: false,
				BindDN:             "",
				BindPassword: LdapBindPassword{
					SecretKeyRef: &corev1.SecretKeySelector{
						LocalObjectReference: corev1.LocalObjectReference{
							Name: "secret",
						},
						Key:      "horusec-platform-jwt",
						Optional: &f,
					},
				},
				UserFilter: "",
				AdminGroup: "",
			},
			GrpcUseCerts: false,
		},
	}

	bytes, _ := json.Marshal(entity)

	print(string(bytes))
	assert.NotEmpty(test, bytes)
}
