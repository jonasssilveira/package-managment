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
	packages := r.collection.Find()
	var sizes []int64
	for _, pack := range packages {
		sizes = append(sizes, pack.Size)
	}
	slices.Sort(sizes)
	return sizes
}
func (r *MongoPackRepository) RemovePack(ctx context.Context, packDoc entity.PackDocument) {
	r.collection.Delete(packDoc.Size)
}
func (r *MongoPackRepository) AddPacks(ctx context.Context, packDoc []entity.PackDocument) {
	r.collection.CreateMany(packDoc)
}
