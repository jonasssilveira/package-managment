package optimalpackage

import (
	"context"
	"order-package/internal/domain/optimalpackage/dto"
	"order-package/internal/infra/repository/mock"
	"reflect"
	"testing"
)

func TestFind(t *testing.T) {
	tests := []struct {
		name           string
		mockRepo       *mock.MockPackRepository
		amount         int64
		expectedOutput dto.PackCombination
	}{
		{
			name: "Exact match single pack",
			mockRepo: &mock.MockPackRepository{
				GetAvailableMock: func(ctx context.Context) []int64 {
					return []int64{250, 500, 1000, 2000, 5000}
				},
			},
			amount:         250,
			expectedOutput: dto.PackCombination{Packs: []dto.Pack{{Size: 250, Amount: 1}}},
		}, {
			name: "Best fit with fewer items and fewer packs",
			mockRepo: &mock.MockPackRepository{
				GetAvailableMock: func(ctx context.Context) []int64 {
					return []int64{250, 500, 1000, 2000, 5000}
				},
			},
			amount:         251,
			expectedOutput: dto.PackCombination{Packs: []dto.Pack{{Size: 500, Amount: 1}}},
		},
		{
			name: "Large number with multiple packs",
			mockRepo: &mock.MockPackRepository{
				GetAvailableMock: func(ctx context.Context) []int64 {
					return []int64{250, 500, 1000, 2000, 5000}
				},
			},
			amount:         12001,
			expectedOutput: dto.PackCombination{Packs: []dto.Pack{{Size: 5000, Amount: 2}, {Size: 2000, Amount: 1}, {Size: 250, Amount: 1}}},
		},
		{
			name: "Fallback to largest pack only",
			mockRepo: &mock.MockPackRepository{
				GetAvailableMock: func(ctx context.Context) []int64 {
					return []int64{23, 31, 53}
				},
			},
			amount:         500000,
			expectedOutput: dto.PackCombination{Packs: []dto.Pack{{Size: 53, Amount: 9429}, {Size: 31, Amount: 7}, {Size: 23, Amount: 2}}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			packCombo := NewPackageUseCase(tt.mockRepo)
			input := dto.PackageAmount{Amount: tt.amount}
			result := packCombo.Find(context.Background(), input)

			if !reflect.DeepEqual(result, tt.expectedOutput) {
				t.Errorf("expected %v, got %v", tt.expectedOutput, result.Packs)
			}
		})
	}
}
