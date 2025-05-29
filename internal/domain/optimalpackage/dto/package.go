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

type Pack struct {
	Size   int64
	Amount int64
}

type PackCombination struct {
	Packs []Pack
}