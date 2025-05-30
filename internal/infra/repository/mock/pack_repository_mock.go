package mock

import (
	"context"
	"order-package/internal/domain/optimalpackage/entity"
)

type MockPackRepository struct {
	GetAvailableMock func(ctx context.Context) []int64
	AddPacksMock     func(ctx context.Context, packs []entity.PackDocument)
	RemovePackMock   func(ctx context.Context, packDoc entity.PackDocument)
}

func (m *MockPackRepository) AddPacks(ctx context.Context, packsDoc []entity.PackDocument) {
	m.AddPacksMock(ctx, packsDoc)
}

func (m *MockPackRepository) RemovePack(ctx context.Context, packsDoc entity.PackDocument) {
	m.RemovePackMock(ctx, packsDoc)
}

func (m *MockPackRepository) GetAvailablePacks(ctx context.Context) []int64 {
	return m.GetAvailableMock(ctx)
}
