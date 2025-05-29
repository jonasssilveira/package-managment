package mock

import (
	"context"
	"order-package/internal/domain/optimalpackage/dto"
)

type MockPackUseCase struct {
	FindMock func(ctx context.Context, packs dto.PackageAmount) dto.PackCombination
	DeleteMock func(ctx context.Context, packs dto.Package) error
	AddMock func(ctx context.Context, packs dto.Package) error
}

func (m *MockPackUseCase) Add(ctx context.Context, packs dto.Package) error {
	return m.AddMock(ctx, packs)
}

func (m *MockPackUseCase) Delete(ctx context.Context, packs dto.Package) error {
	return m.DeleteMock(ctx, packs)
}

func (m *MockPackUseCase) Find(ctx context.Context, packs dto.PackageAmount) dto.PackCombination {
	return m.FindMock(ctx, packs)
}
