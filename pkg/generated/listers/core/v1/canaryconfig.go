/*
Copyright The Fission Authors.

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

package v1

import (
	v1 "github.com/fission/fission/pkg/apis/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CanaryConfigLister helps list CanaryConfigs.
// All objects returned here must be treated as read-only.
type CanaryConfigLister interface {
	// List lists all CanaryConfigs in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.CanaryConfig, err error)
	// CanaryConfigs returns an object that can list and get CanaryConfigs.
	CanaryConfigs(namespace string) CanaryConfigNamespaceLister
	CanaryConfigListerExpansion
}

// _canaryConfigLister implements the CanaryConfigLister interface.
type _canaryConfigLister struct {
	indexer cache.Indexer
}

// NewCanaryConfigLister returns a new CanaryConfigLister.
func NewCanaryConfigLister(indexer cache.Indexer) CanaryConfigLister {
	return &_canaryConfigLister{indexer: indexer}
}

// List lists all CanaryConfigs in the indexer.
func (s *_canaryConfigLister) List(selector labels.Selector) (ret []*v1.CanaryConfig, err error) {
	err = cache.ListAll(s.indexer, selector, func(m any) {
		ret = append(ret, m.(*v1.CanaryConfig))
	})
	return ret, err
}

// CanaryConfigs returns an object that can list and get CanaryConfigs.
func (s *_canaryConfigLister) CanaryConfigs(namespace string) CanaryConfigNamespaceLister {
	return _canaryConfigNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CanaryConfigNamespaceLister helps list and get CanaryConfigs.
// All objects returned here must be treated as read-only.
type CanaryConfigNamespaceLister interface {
	// List lists all CanaryConfigs in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.CanaryConfig, err error)
	// Get retrieves the CanaryConfig from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.CanaryConfig, error)
	CanaryConfigNamespaceListerExpansion
}

// _canaryConfigNamespaceLister implements the CanaryConfigNamespaceLister
// interface.
type _canaryConfigNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CanaryConfigs in the indexer for a given namespace.
func (s _canaryConfigNamespaceLister) List(selector labels.Selector) (ret []*v1.CanaryConfig, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m any) {
		ret = append(ret, m.(*v1.CanaryConfig))
	})
	return ret, err
}

// Get retrieves the CanaryConfig from the indexer for a given namespace and name.
func (s _canaryConfigNamespaceLister) Get(name string) (*v1.CanaryConfig, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("canaryconfig"), name)
	}
	return obj.(*v1.CanaryConfig), nil
}
