package adapters

import (
	"context"
	"order-package/internal/domain/optimalpackage/entity"
)

type PackRepository interface {
	GetAvailablePacks(ctx context.Context) []int64
	RemovePack(ctx context.Context, packDoc entity.PackDocument) error
	AddPack(ctx context.Context, document entity.PackDocument) error
} 
