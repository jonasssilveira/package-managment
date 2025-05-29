package optimalpackage

import (
	"context"
	"order-package/internal/domain/optimalpackage/adapters"
	"order-package/internal/domain/optimalpackage/dto"
)

type PackCombo struct {
	packRepository adapters.PackRepository
}

func NewPackCombo(packRepository adapters.PackRepository) *PackCombo {
	return &PackCombo{
		packRepository: packRepository,
	}
}

func (pk PackCombo) Find(ctx context.Context, packs dto.Package) dto.PackCombination {
	packages := pk.packRepository.GetAvailablePacks(ctx)

	limit := packs.Amount + packages[len(packages)]

	dp := make([]*struct {
		packCount int
		packCombo dto.PackCombination
	}, limit+1)

	dp[0] = &struct {
		packCount int
		packCombo dto.PackCombination
	}{0, dto.PackCombination{}}

	for i := int64(1); i <= limit; i++ {
		for _, size := range packages {
			if i-size >= 0 && dp[i-size] != nil {
				newCount := dp[i-size].packCount + 1
				if dp[i] == nil || newCount < dp[i].packCount {
					var newCombo dto.PackCombination
					for k, v := range dp[i-size].packCombo.Packs {
						newCombo.Packs[k] = v
					}
					newCombo.Packs[size]++
					dp[i] = &struct {
						packCount int
						packCombo dto.PackCombination
					}{newCount, newCombo}
				}
			}
		}
	}

	for i := packs.Amount; i <= limit; i++ {
		if dp[i] != nil {
			return dp[i].packCombo
		}
	}

	return dto.PackCombination{Packs: make(map[int64]int64)}
}
