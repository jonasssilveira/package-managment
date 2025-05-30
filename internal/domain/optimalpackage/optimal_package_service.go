package optimalpackage

import (
	"net/http"
	"order-package/internal/domain/optimalpackage/adapters"
	"order-package/internal/domain/optimalpackage/dto"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OptimalPackage interface {
	Find(context *gin.Context) map[int]int
}

type OptimalPackageService struct {
	optimalPackageUseCase adapters.FindOptimalPacks
}

func NewOptimalPackageService(optPackageUseCase adapters.FindOptimalPacks) OptimalPackageService {
	return OptimalPackageService{
		optimalPackageUseCase: optPackageUseCase,
	}
}

func (s OptimalPackageService) Find(ctx *gin.Context) {
	var pack dto.PackageAmount
	err := ctx.BindJSON(&pack)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	packages := s.optimalPackageUseCase.Find(ctx, pack)
	ctx.JSON(http.StatusOK, packages)
}

func (s OptimalPackageService) Delete(ctx *gin.Context) {
	size, err := strconv.ParseInt(ctx.Param("size"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	pack := dto.Package{Size: size}
	s.optimalPackageUseCase.Delete(ctx, pack)
	ctx.JSON(http.StatusOK, nil)
}

func (s OptimalPackageService) Create(ctx *gin.Context) {
	var pack dto.Packages
	err := ctx.BindJSON(&pack)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	s.optimalPackageUseCase.Add(ctx, pack)
	ctx.JSON(http.StatusCreated, nil)
}
