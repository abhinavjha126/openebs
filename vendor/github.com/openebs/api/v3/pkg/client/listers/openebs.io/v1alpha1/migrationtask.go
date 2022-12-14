/*
Copyright 2021 The OpenEBS Authors

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
	v1alpha1 "github.com/openebs/api/v3/pkg/apis/openebs.io/v1alpha1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// MigrationTaskLister helps list MigrationTasks.
// All objects returned here must be treated as read-only.
type MigrationTaskLister interface {
	// List lists all MigrationTasks in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.MigrationTask, err error)
	// MigrationTasks returns an object that can list and get MigrationTasks.
	MigrationTasks(namespace string) MigrationTaskNamespaceLister
	MigrationTaskListerExpansion
}

// migrationTaskLister implements the MigrationTaskLister interface.
type migrationTaskLister struct {
	indexer cache.Indexer
}

// NewMigrationTaskLister returns a new MigrationTaskLister.
func NewMigrationTaskLister(indexer cache.Indexer) MigrationTaskLister {
	return &migrationTaskLister{indexer: indexer}
}

// List lists all MigrationTasks in the indexer.
func (s *migrationTaskLister) List(selector labels.Selector) (ret []*v1alpha1.MigrationTask, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MigrationTask))
	})
	return ret, err
}

// MigrationTasks returns an object that can list and get MigrationTasks.
func (s *migrationTaskLister) MigrationTasks(namespace string) MigrationTaskNamespaceLister {
	return migrationTaskNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// MigrationTaskNamespaceLister helps list and get MigrationTasks.
// All objects returned here must be treated as read-only.
type MigrationTaskNamespaceLister interface {
	// List lists all MigrationTasks in the indexer for a given namespace.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1alpha1.MigrationTask, err error)
	// Get retrieves the MigrationTask from the indexer for a given namespace and name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1alpha1.MigrationTask, error)
	MigrationTaskNamespaceListerExpansion
}

// migrationTaskNamespaceLister implements the MigrationTaskNamespaceLister
// interface.
type migrationTaskNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all MigrationTasks in the indexer for a given namespace.
func (s migrationTaskNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.MigrationTask, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.MigrationTask))
	})
	return ret, err
}

// Get retrieves the MigrationTask from the indexer for a given namespace and name.
func (s migrationTaskNamespaceLister) Get(name string) (*v1alpha1.MigrationTask, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("migrationtask"), name)
	}
	return obj.(*v1alpha1.MigrationTask), nil
}
