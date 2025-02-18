/*
Copyright 2022 The Kubernetes Authors.

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
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1 "k8s.io/cloud-provider-gcp/crd/apis/network/v1"
)

// GKENetworkParamsLister helps list GKENetworkParamses.
// All objects returned here must be treated as read-only.
type GKENetworkParamsLister interface {
	// List lists all GKENetworkParamses in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.GKENetworkParams, err error)
	// GKENetworkParamses returns an object that can list and get GKENetworkParamses.
	GKENetworkParamses(namespace string) GKENetworkParamsNamespaceLister
	GKENetworkParamsListerExpansion
}

// gKENetworkParamsLister implements the GKENetworkParamsLister interface.
type gKENetworkParamsLister struct {
	indexer cache.Indexer
}

// NewGKENetworkParamsLister returns a new GKENetworkParamsLister.
func NewGKENetworkParamsLister(indexer cache.Indexer) GKENetworkParamsLister {
	return &gKENetworkParamsLister{indexer: indexer}
}

// List lists all GKENetworkParamses in the indexer.
func (s *gKENetworkParamsLister) List(selector labels.Selector) (ret []*v1.GKENetworkParams, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.GKENetworkParams))
	})
	return ret, err
}

// GKENetworkParamses returns an object that can list and get GKENetworkParamses.
func (s *gKENetworkParamsLister) GKENetworkParamses(namespace string) GKENetworkParamsNamespaceLister {
	return gKENetworkParamsNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// GKENetworkParamsNamespaceLister helps list and get GKENetworkParamses.
// All objects returned here must be treated as read-only.
type GKENetworkParamsNamespaceLister interface {
	// List lists all GKENetworkParamses in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.GKENetworkParams, err error)
	// Get retrieves the GKENetworkParams from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.GKENetworkParams, error)
	GKENetworkParamsNamespaceListerExpansion
}

// gKENetworkParamsNamespaceLister implements the GKENetworkParamsNamespaceLister
// interface.
type gKENetworkParamsNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all GKENetworkParamses in the indexer for a given namespace.
func (s gKENetworkParamsNamespaceLister) List(selector labels.Selector) (ret []*v1.GKENetworkParams, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.GKENetworkParams))
	})
	return ret, err
}

// Get retrieves the GKENetworkParams from the indexer for a given namespace and name.
func (s gKENetworkParamsNamespaceLister) Get(name string) (*v1.GKENetworkParams, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("gkenetworkparams"), name)
	}
	return obj.(*v1.GKENetworkParams), nil
}
