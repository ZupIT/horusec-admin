package kubernetes

import (
	"errors"
	"fmt"
	"os"

	clientset "github.com/ZupIT/horusec-admin/pkg/client/clientset/versioned"
	client "github.com/ZupIT/horusec-admin/pkg/client/clientset/versioned/typed/install/v1alpha1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
)

func NewRestConfig() (*rest.Config, error) {
	kubecfg := os.Getenv("KUBECONFIG")

	if len(kubecfg) == 0 {
		return nil, errors.New("environment variable KUBECONFIG need to be set")
	}

	restcfg, err := clientcmd.BuildConfigFromFlags("", kubecfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create k8s rest client: %w", err)
	}

	return restcfg, nil
}

func NewHorusecManagerClient(restConfig *rest.Config) (client.HorusecManagerInterface, error) {
	namespace := os.Getenv("NAMESPACE")

	if len(namespace) == 0 {
		return nil, errors.New("environment variable NAMESPACE need to be set")
	}

	c, err := clientset.NewForConfig(restConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create a new HorusecManager client: %w", err)
	}

	return c.InstallV1alpha1().HorusecManagers(namespace), nil
}
