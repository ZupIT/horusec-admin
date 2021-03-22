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
