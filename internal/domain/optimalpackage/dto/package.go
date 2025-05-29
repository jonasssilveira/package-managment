package dto

import "order-package/internal/domain/optimalpackage/entity"

type PackageAmount struct {
	Amount int64 `json:"amount"`
}

type Package struct{
	Size int64 `json:"size"`
}

func (p Package) ToEntity()entity.PackDocument{
	return entity.PackDocument{Size : p.Size}
}

type PackCombination struct {
	Packs map[int64]int64
}
