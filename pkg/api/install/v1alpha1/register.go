package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const (
	// Package-wide consts from generator "register".
	GroupName = "install.horusec.io"
)

var (
	// Package-wide variables from generator "register".
	SchemeGroupVersion = schema.GroupVersion{Group: GroupName, Version: "v1alpha1"}
	SchemeBuilder      = runtime.NewSchemeBuilder(addKnownTypes)
	localSchemeBuilder = &SchemeBuilder
	AddToScheme        = localSchemeBuilder.AddToScheme
)

func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&HorusecManager{},
		&HorusecManagerList{},
	)

	scheme.AddKnownTypes(
		SchemeGroupVersion,
		&v1.Status{},
	)

	v1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
