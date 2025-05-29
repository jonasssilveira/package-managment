package mongo

import (
	"context"
	"order-package/internal/domain/optimalpackage/entity"
	"order-package/internal/infra/database"
	"slices"
)

const (
	PackCollection = "pack_sizes"
)

type MongoPackRepository struct {
	collection database.Collection
}

func NewMongoPackRepository(collection database.Collection) *MongoPackRepository {
	return &MongoPackRepository{collection: collection}
}

func (r *MongoPackRepository) GetAvailablePacks(ctx context.Context) []int64 {
	packages := r.collection.Find(ctx)
	var sizes []int64
	for i := 0; i < len(packages); i++ {
		sizes = append(sizes, packages[i].(*entity.PackDocument).Size)
	}
	slices.Sort(sizes)
	return sizes
}
func (r *MongoPackRepository) RemovePack(ctx context.Context, packDoc entity.PackDocument)error{
	return r.collection.Delete(ctx, packDoc)
}
func (r *MongoPackRepository) AddPack(ctx context.Context, packDoc entity.PackDocument)error{
	return r.collection.Create(ctx, packDoc)
}
