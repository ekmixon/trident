// Copyright 2021 NetApp, Inc. All Rights Reserved.

// Code generated by informer-gen. DO NOT EDIT.

package v1

import (
	"context"
	time "time"

	netappv1 "github.com/netapp/trident/persistent_store/crd/apis/netapp/v1"
	versioned "github.com/netapp/trident/persistent_store/crd/client/clientset/versioned"
	internalinterfaces "github.com/netapp/trident/persistent_store/crd/client/informers/externalversions/internalinterfaces"
	v1 "github.com/netapp/trident/persistent_store/crd/client/listers/netapp/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// TridentSnapshotInformer provides access to a shared informer and lister for
// TridentSnapshots.
type TridentSnapshotInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1.TridentSnapshotLister
}

type tridentSnapshotInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewTridentSnapshotInformer constructs a new informer for TridentSnapshot type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewTridentSnapshotInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredTridentSnapshotInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredTridentSnapshotInformer constructs a new informer for TridentSnapshot type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredTridentSnapshotInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options metav1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TridentV1().TridentSnapshots(namespace).List(context.TODO(), options)
			},
			WatchFunc: func(options metav1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.TridentV1().TridentSnapshots(namespace).Watch(context.TODO(), options)
			},
		},
		&netappv1.TridentSnapshot{},
		resyncPeriod,
		indexers,
	)
}

func (f *tridentSnapshotInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredTridentSnapshotInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *tridentSnapshotInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&netappv1.TridentSnapshot{}, f.defaultInformer)
}

func (f *tridentSnapshotInformer) Lister() v1.TridentSnapshotLister {
	return v1.NewTridentSnapshotLister(f.Informer().GetIndexer())
}
