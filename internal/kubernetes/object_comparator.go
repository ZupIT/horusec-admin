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
	"github.com/google/go-cmp/cmp"
	k8s "k8s.io/apimachinery/pkg/runtime"
)

var ignore = []string{
	"TypeMeta.APIVersion",
	"TypeMeta.Kind",
	"ObjectMeta.CreationTimestamp",
	"ObjectMeta.Finalizers",
	"ObjectMeta.Generation",
	"ObjectMeta.ManagedFields",
	"ObjectMeta.Namespace",
	"ObjectMeta.ResourceVersion",
	"ObjectMeta.SelfLink",
	"ObjectMeta.UID",
}

type ObjectComparator struct{}

func (c *ObjectComparator) filter(path cmp.Path) bool {
	for _, p := range ignore {
		if p == path.String() {
			return true
		}
	}
	return false
}

func (c *ObjectComparator) Diff(x, y k8s.Object) string {
	return cmp.Diff(x, y, cmp.FilterPath(c.filter, cmp.Ignore()))
}
