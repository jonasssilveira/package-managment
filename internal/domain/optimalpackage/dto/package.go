package dto

type Package struct {
	Amount int64 `json:"amount"`
}

type PackCombination struct {
	Packs map[int64]int64
}
