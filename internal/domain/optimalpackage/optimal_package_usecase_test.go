package optimalpackage

import (
	"context"
	"errors"
	"order-package/internal/domain/optimalpackage/dto"
	"order-package/internal/domain/optimalpackage/entity"
	"order-package/internal/infra/repository/mongo/mock"
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
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
			expectedOutput: dto.PackCombination{Packs: []dto.Pack{{Size:250, Amount: 1}}},
		},{
			name:           "Best fit with fewer items and fewer packs",
			mockRepo: &mock.MockPackRepository{
				GetAvailableMock: func(ctx context.Context) []int64 { 
					return []int64{250, 500, 1000, 2000, 5000}
				},
			},
			amount:         251,
			expectedOutput: dto.PackCombination{Packs: []dto.Pack{{Size:500, Amount: 1}}},
		},
		{
			name:      "Large number with multiple packs",
			mockRepo: &mock.MockPackRepository{
				GetAvailableMock: func(ctx context.Context) []int64 { 
					return []int64{250, 500, 1000, 2000, 5000}
				},
			},
			amount:    12001,
			expectedOutput: dto.PackCombination{Packs: []dto.Pack{{Size:5000, Amount: 2}, {Size:2000, Amount: 1}, {Size:250, Amount: 1}}},
		},
		{
			name:           "Fallback to largest pack only",
			mockRepo: &mock.MockPackRepository{
				GetAvailableMock: func(ctx context.Context) []int64 { 
					return []int64{23, 31, 53}
				},
			},
			amount:         500000,
			expectedOutput: dto.PackCombination{Packs: []dto.Pack{{Size:53, Amount: 9429}, {Size:31, Amount: 7}, {Size:23, Amount: 2}, }},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			packCombo := NewPackCombo(tt.mockRepo)
			input := dto.PackageAmount{Amount: tt.amount}
			result := packCombo.Find(context.Background(), input)

			if !reflect.DeepEqual(result, tt.expectedOutput) {
				t.Errorf("expected %v, got %v", tt.expectedOutput, result.Packs)
			}
		})
	}
}

func TestPackCombo_Add(t *testing.T) {
	tests := []struct {
		name      string
		mockRepo  *mock.MockPackRepository
		input     dto.Package
		mockError error
	}{
		{
			name: "successfully adds a pack",
			input: dto.Package{
				Size:   500,
			},
			mockRepo: &mock.MockPackRepository{
				AddPackMock: func(ctx context.Context, packEntity entity.PackDocument) error {
					return nil
				},
			},
			mockError: nil,
		},
		{
			name: "repository returns an error",
			input: dto.Package{
				Size:   1000,
			},
			mockRepo: &mock.MockPackRepository{
				AddPackMock: func(ctx context.Context, packEntity entity.PackDocument) error {
					return errors.New("db error")
				},
			},
			mockError: errors.New("db error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			usecase := NewPackCombo(tt.mockRepo)
			err := usecase.Add(context.Background(), tt.input)

			if tt.mockError != nil {
				require.Error(t, err)
				require.EqualError(t, err, tt.mockError.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestPackCombo_Delete(t *testing.T) {
	tests := []struct {
		name      string
		mockRepo  *mock.MockPackRepository
		input     dto.Package
		mockError error
	}{
		{
			name: "successfully deletes a pack",
			input: dto.Package{
				Size:   250,
			},
			mockRepo: &mock.MockPackRepository{
				RemovePackMock: func(ctx context.Context, packEntity entity.PackDocument) error {
					return nil
				},
			},
			mockError: nil,
		},
		{
			name: "repository returns an error on delete",
			input: dto.Package{
				Size:   500,
			},
			mockRepo: &mock.MockPackRepository{
				RemovePackMock: func(ctx context.Context, packEntity entity.PackDocument) error {
					return  errors.New("db error")
				},
			},
			mockError: errors.New("db error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			usecase := NewPackCombo(tt.mockRepo)
			err := usecase.Delete(context.Background(), tt.input)

			if tt.mockError != nil {
				require.Error(t, err)
				require.EqualError(t, err, tt.mockError.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}
