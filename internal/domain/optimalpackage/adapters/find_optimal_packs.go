package adapters

import (
	"context"
	"order-package/internal/domain/optimalpackage/dto"
)

type FindOptimalPacks interface {
	// Find findOptimalPackCombo calculates the optimal combination of pack sizes to fulfill a given order amount.
	//
	// The function satisfies the following rules:
	// 1. Only whole packs are allowed — packs cannot be broken.
	// 2. The total number of items shipped must be the smallest possible that is greater than or equal to the requested amount.
	// 3. Among all such combinations, the one with the fewest total number of packs is preferred.
	//
	// Parameters:
	//   - amount: The total number of items ordered.
	//
	// Returns:
	//
	//	A map[int]int where the key is the pack size and the value is the number of packs used.
	//	The map represents the combination of packs that best fulfills the order based on the constraints.
	//
	// Example:
	//
	//	input: amount = 12001, sizes = []int{250, 500, 1000, 2000, 5000}
	//	output: map[int]int{250: 1, 2000: 1, 5000: 2}
	//
	// Note:
	//   - The algorithm uses dynamic programming and iterates from 1 to the given amount.
	//   - It keeps track of the minimal total item count that satisfies the order, and among them, the fewest packs.
	//   - Time complexity is approximately O(n * m), where n = amount and m = number of pack sizes.
	Find(ctx context.Context, packs dto.PackageAmount) dto.PackCombination
	Delete(ctx context.Context, packs dto.Package)
	Add(ctx context.Context, packs dto.Packages)
}
