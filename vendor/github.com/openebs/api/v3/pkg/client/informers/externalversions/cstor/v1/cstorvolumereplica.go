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

package v1

import (
	"context"
	time "time"

	cstorv1 "github.com/openebs/api/v3/pkg/apis/cstor/v1"
	versioned "github.com/openebs/api/v3/pkg/client/clientset/versioned"
	internalinterfaces "github.com/openebs/api/v3/pkg/client/informers/externalversions/internalinterfaces"
	v1 "github.com/openebs/api/v3/pkg/client/listers/cstor/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// CStorVolumeReplicaInformer provides access to a shared informer and lister for
// CStorVolumeReplicas.
type CStorVolumeReplicaInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.CStorVolumeReplicaLister
}

type cStorVolumeReplicaInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewCStorVolumeReplicaInformer constructs a new informer for CStorVolumeReplica type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewCStorVolumeReplicaInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredCStorVolumeReplicaInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredCStorVolumeReplicaInformer constructs a new informer for CStorVolumeReplica type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredCStorVolumeReplicaInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CstorV1().CStorVolumeReplicas(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.CstorV1().CStorVolumeReplicas(namespace).Watch(context.TODO(), options)
			},
		},
		&cstorv1.CStorVolumeReplica{},
		resyncPeriod,
		indexers,
	)
}

func (f *cStorVolumeReplicaInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredCStorVolumeReplicaInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *cStorVolumeReplicaInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&cstorv1.CStorVolumeReplica{}, f.defaultInformer)
}

func (f *cStorVolumeReplicaInformer) Lister() v1.CStorVolumeReplicaLister {
	return v1.NewCStorVolumeReplicaLister(f.Informer().GetIndexer())
}
