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
	if kubecfg != "" {
		cfg, err := clientcmd.BuildConfigFromFlags("", kubecfg)
		if err != nil {
			return nil, fmt.Errorf("failed to configure the rest client from kubeconfig filepath: %w", err)
		}
		return cfg, nil
	}

	cfg, err := rest.InClusterConfig()
	if err != nil {
		return nil, fmt.Errorf("failed to configure in cluster rest client: %w", err)
	}

	return cfg, nil
}

func NewHorusecManagerClient(restConfig *rest.Config) (client.HorusecPlatformInterface, error) {
	namespace := os.Getenv("NAMESPACE")

	if namespace == "" {
		return nil, errors.New("environment variable NAMESPACE need to be set")
	}

	c, err := clientset.NewForConfig(restConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create a new HorusecPlatform client: %w", err)
	}

	return c.InstallV1alpha1().HorusecPlatforms(namespace), nil
}
