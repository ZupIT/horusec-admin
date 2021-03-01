/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "github.com/tiagoangelozup/horusec-admin/pkg/api/install/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// HorusecManagerLister helps list HorusecManagers.
// All objects returned here must be treated as read-only.
type HorusecManagerLister interface {
	// List lists all HorusecManagers in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.HorusecManager, err error)
	// HorusecManagers returns an object that can list and get HorusecManagers.
	HorusecManagers(namespace string) HorusecManagerNamespaceLister
	HorusecManagerListerExpansion
}

// horusecManagerLister implements the HorusecManagerLister interface.
type horusecManagerLister struct {
	indexer cache.Indexer
}

// NewHorusecManagerLister returns a new HorusecManagerLister.
func NewHorusecManagerLister(indexer cache.Indexer) HorusecManagerLister {
	return &horusecManagerLister{indexer: indexer}
}

// List lists all HorusecManagers in the indexer.
func (s *horusecManagerLister) List(selector labels.Selector) (ret []*v1alpha1.HorusecManager, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.HorusecManager))
	})
	return ret, err
}

// HorusecManagers returns an object that can list and get HorusecManagers.
func (s *horusecManagerLister) HorusecManagers(namespace string) HorusecManagerNamespaceLister {
	return horusecManagerNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// HorusecManagerNamespaceLister helps list and get HorusecManagers.
// All objects returned here must be treated as read-only.
type HorusecManagerNamespaceLister interface {
	// List lists all HorusecManagers in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.HorusecManager, err error)
	// Get retrieves the HorusecManager from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.HorusecManager, error)
	HorusecManagerNamespaceListerExpansion
}

// horusecManagerNamespaceLister implements the HorusecManagerNamespaceLister
// interface.
type horusecManagerNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all HorusecManagers in the indexer for a given namespace.
func (s horusecManagerNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.HorusecManager, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.HorusecManager))
	})
	return ret, err
}

// Get retrieves the HorusecManager from the indexer for a given namespace and name.
func (s horusecManagerNamespaceLister) Get(name string) (*v1alpha1.HorusecManager, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("horusecmanager"), name)
	}
	return obj.(*v1alpha1.HorusecManager), nil
}
