package dto

import "order-package/internal/domain/optimalpackage/entity"

type PackageAmount struct {
	Amount int64 `json:"amount"`
}

type Package struct {
	Size int64 `json:"size"`
}

type Packages struct {
	Packages []Package `json:"packages"`
}

func (p Package) ToEntity() entity.PackDocument {
	return entity.PackDocument{Size: p.Size}
}

func (p Packages) ToEntity() []entity.PackDocument {
	var packages []entity.PackDocument
	for i := 0; i < len(p.Packages); i++ {
		packages = append(packages, p.Packages[i].ToEntity())
	}
	return packages
}

type Pack struct {
	Size   int64 `json:"size"`
	Amount int64 `json:"amount"`
}

type PackCombination struct {
	Packs []Pack `json:"packs"`
}
