package optimalpackage

import (
	"context"
	"order-package/internal/domain/optimalpackage/dto"
	"order-package/internal/domain/optimalpackage/mock"
	"reflect"
	"testing"
)

func TestFind(t *testing.T) {
	tests := []struct {
		name           string
		mockRepo       *mock.MockPackRepository
		amount         int64
		expectedOutput map[int64]int64
	}{
		{
			name: "Exact match single pack",
			mockRepo: &mock.MockPackRepository{
				GetAvailable: func(ctx context.Context) []int64 {
					return []int64{250, 500, 1000, 2000, 5000}
				},
			},
			amount:         250,
			expectedOutput: map[int64]int64{250: 1},
		},{
			name:           "Best fit with fewer items and fewer packs",
			mockRepo: &mock.MockPackRepository{
				GetAvailable: func(ctx context.Context) []int64 { 
					return []int64{250, 500, 1000, 2000, 5000}
				},
			},
			amount:         251,
			expectedOutput: map[int64]int64{500: 1},
		},
		{
			name:      "Large number with multiple packs",
			mockRepo: &mock.MockPackRepository{
				GetAvailable: func(ctx context.Context) []int64 { 
					return []int64{250, 500, 1000, 2000, 5000}
				},
			},
			amount:    12001,
			expectedOutput: map[int64]int64{
				5000: 2,
				2000: 1,
				250:  1,
			},
		},
		{
			name:           "Fallback to largest pack only",
			mockRepo: &mock.MockPackRepository{
				GetAvailable: func(ctx context.Context) []int64 { 
					return []int64{23, 31, 53}
				},
			},
			amount:         500000,
			expectedOutput: map[int64]int64{23: 2, 31: 7, 53: 9429},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			packCombo := NewPackCombo(tt.mockRepo)
			input := dto.Package{Amount: tt.amount}
			result := packCombo.Find(context.Background(), input)

			if !reflect.DeepEqual(result.Packs, tt.expectedOutput) {
				t.Errorf("expected %v, got %v", tt.expectedOutput, result.Packs)
			}
		})
	}
}