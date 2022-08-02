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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	time "time"

	openebsiov1alpha1 "github.com/openebs/api/v3/pkg/apis/openebs.io/v1alpha1"
	versioned "github.com/openebs/api/v3/pkg/client/clientset/versioned"
	internalinterfaces "github.com/openebs/api/v3/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/openebs/api/v3/pkg/client/listers/openebs.io/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// CStorCompletedBackupInformer provides access to a shared informer and lister for
// CStorCompletedBackups.
type CStorCompletedBackupInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.CStorCompletedBackupLister
}

type cStorCompletedBackupInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewCStorCompletedBackupInformer constructs a new informer for CStorCompletedBackup type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCStorCompletedBackupInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCStorCompletedBackupInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredCStorCompletedBackupInformer constructs a new informer for CStorCompletedBackup type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCStorCompletedBackupInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OpenebsV1alpha1().CStorCompletedBackups(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.OpenebsV1alpha1().CStorCompletedBackups(namespace).Watch(context.TODO(), options)
			},
		},
		&openebsiov1alpha1.CStorCompletedBackup{},
		resyncPeriod,
		indexers,
	)
}

func (f *cStorCompletedBackupInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCStorCompletedBackupInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *cStorCompletedBackupInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&openebsiov1alpha1.CStorCompletedBackup{}, f.defaultInformer)
}

func (f *cStorCompletedBackupInformer) Lister() v1alpha1.CStorCompletedBackupLister {
	return v1alpha1.NewCStorCompletedBackupLister(f.Informer().GetIndexer())
}
