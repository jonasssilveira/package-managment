package main

import (
	"order-package/internal/domain/optimalpackage"
	"order-package/internal/infra"
	"order-package/internal/infra/database"
	"order-package/internal/infra/repository"
)

func main() {
	server := infra.NewServer()
	collection := database.NewInMemoryPackRepository()
	repository := repository.NewMongoPackRepository(collection)
	optimalUseCase := optimalpackage.NewPackageUseCase(repository)
	optimalService := optimalpackage.NewOptimalPackageService(optimalUseCase)

	server.Start()
	handler := infra.NewHandle(*server)
	handler.Static("/", "./static")
	handler.Post("/packs-find", optimalService.Find)
	handler.Post("/packs-create", optimalService.Create)
	handler.Delete("/packs/:size", optimalService.Delete)

	handler.Run()
}
