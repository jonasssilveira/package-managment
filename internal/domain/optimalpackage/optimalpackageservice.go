package optimalpackage

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"order-package/internal/domain/optimalpackage/adapters"
	"order-package/internal/domain/optimalpackage/dto"
	"order-package/internal/infra/database"
	"order-package/internal/infra/repository/mongo"
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
	c := context.WithValue(ctx, database.CollectionConfig, mongo.PackCollection)
	var pack dto.Package
	err := ctx.BindJSON(&pack)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	packages := s.optimalPackageUseCase.Find(c, pack)
	ctx.JSON(http.StatusOK, packages)
}

func (s OptimalPackageService) Delete(context *gin.Context) {

}
