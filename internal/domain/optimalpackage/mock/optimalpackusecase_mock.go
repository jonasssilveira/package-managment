package mock

import "context"


type MockPackRepository struct {
	GetAvailable func(ctx context.Context) []int64
}

func (m *MockPackRepository) GetAvailablePacks(ctx context.Context) []int64 {
	return m.GetAvailable(ctx)
}