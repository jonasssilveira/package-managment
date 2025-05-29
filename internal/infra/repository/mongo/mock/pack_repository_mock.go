
package mock

import (
	"context"
	"order-package/internal/domain/optimalpackage/entity"
)

type MockPackRepository struct {
	GetAvailableMock func(ctx context.Context) []int64
	AddPackMock    func(ctx context.Context, pack entity.PackDocument) error
	RemovePackMock func(ctx context.Context, pack entity.PackDocument) error
}

func (m *MockPackRepository) AddPack(ctx context.Context, packDoc entity.PackDocument) error {
	return m.AddPackMock(ctx, packDoc)
}

func (m *MockPackRepository) RemovePack(ctx context.Context, packDoc entity.PackDocument) error {
	return m.RemovePackMock(ctx, packDoc)
}

func (m *MockPackRepository) GetAvailablePacks(ctx context.Context) []int64 {
	return m.GetAvailableMock(ctx)
}
