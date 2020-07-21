package mocks

import (
	"context"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	toolscache "sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type FakeCache struct {
}

func (FakeCache) Get(ctx context.Context, key client.ObjectKey, obj runtime.Object) error {
	return errors.NewNotFound(schema.GroupResource{}, "")
}

func (FakeCache) List(ctx context.Context, list runtime.Object, opts ...client.ListOption) error {
	panic("implement me")
}

func (FakeCache) GetInformer(ctx context.Context, obj runtime.Object) (toolscache.Informer, error) {
	panic("implement me")
}

func (FakeCache) GetInformerForKind(gctx context.Context, vk schema.GroupVersionKind) (toolscache.Informer, error) {
	panic("implement me")
}

func (FakeCache) Start(stopCh <-chan struct{}) error {
	panic("implement me")
}

func (FakeCache) WaitForCacheSync(stop <-chan struct{}) bool {
	panic("implement me")
}

func (FakeCache) IndexField(ctx context.Context, obj runtime.Object, field string, extractValue client.IndexerFunc) error {
	panic("implement me")
}
