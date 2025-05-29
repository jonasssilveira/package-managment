package optimalpackage

import (
	"context"
	"net/http"
	"order-package/internal/domain/optimalpackage/adapters"
	"order-package/internal/domain/optimalpackage/dto"
	"order-package/internal/infra/database"
	"order-package/internal/infra/repository/mongo"
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
	c := context.WithValue(ctx, database.CollectionName, mongo.PackCollection)
	var pack dto.PackageAmount
	err := ctx.BindJSON(&pack)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	packages := s.optimalPackageUseCase.Find(c, pack)
	ctx.JSON(http.StatusOK, packages)
}

func (s OptimalPackageService) Delete(ctx *gin.Context) {
	c := context.WithValue(ctx, database.CollectionName, mongo.PackCollection)
	size, err:= strconv.ParseInt(ctx.Param("size"), 10, 64)
	if err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	pack := dto.Package{Size:size} 
	packages := s.optimalPackageUseCase.Delete(c, pack)
	ctx.JSON(http.StatusOK, packages)
}

func (s OptimalPackageService) Create(ctx *gin.Context) {
	c := context.WithValue(ctx, database.CollectionName, mongo.PackCollection)
	var pack dto.Package
	err := ctx.BindJSON(&pack)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := s.optimalPackageUseCase.Add(c, pack); err!=nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, nil)
}
