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

func (pk PackCombo) Find(ctx context.Context, packs dto.PackageAmount) dto.PackCombination {
	packages := pk.packRepository.GetAvailablePacks(ctx)

	limit := packs.Amount + packages[len(packages)-1]

	dp := make([]*struct {
		packCount int
		packCombo dto.PackCombination
	}, limit+1)

	// base case: 0 items = 0 packs
	dp[0] = &struct {
		packCount int
		packCombo dto.PackCombination
	}{
		packCount: 0,
		packCombo: dto.PackCombination{Packs: []dto.Pack{}},
	}

	for i := int64(1); i <= limit; i++ {
		for _, size := range packages {
			if i-size >= 0 && dp[i-size] != nil {
				newCount := dp[i-size].packCount + 1

				if dp[i] == nil || newCount < dp[i].packCount {
					// deep copy the combination from i - size
					newCombo := deepCopyPacks(dp[i-size].packCombo.Packs)

					// update amount for the pack size or add it
					found := false
					for idx, p := range newCombo {
						if p.Size == size {
							newCombo[idx].Amount++
							found = true
							break
						}
					}
					if !found {
						newCombo = append(newCombo, dto.Pack{
							Size:   size,
							Amount: 1,
						})
					}

					// assign new combo
					dp[i] = &struct {
						packCount int
						packCombo dto.PackCombination
					}{
						packCount: newCount,
						packCombo: dto.PackCombination{Packs: newCombo},
					}
				}
			}
		}
	}

	for i := packs.Amount; i <= limit; i++ {
		if dp[i] != nil {
			return dp[i].packCombo
		}
	}

	return dto.PackCombination{}
}
func deepCopyPacks(src []dto.Pack) []dto.Pack {
	copied := make([]dto.Pack, len(src))
	copy(copied, src)
	return copied
}
func (pk PackCombo) Delete(ctx context.Context, packs dto.Package) error{
	return pk.packRepository.RemovePack(ctx, packs.ToEntity())
}

func (pk PackCombo) Add(ctx context.Context, packs dto.Package)error{
	return pk.packRepository.AddPack(ctx, packs.ToEntity())
}