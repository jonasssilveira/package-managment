package adapters

import "context"

type PackRepository interface {
	GetAvailablePacks(ctx context.Context) []int64
}
