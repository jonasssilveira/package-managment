package main

import (
	"order-package/internal/domain/optimalpackage"
	"order-package/internal/infra"
	"order-package/internal/infra/config"
	"order-package/internal/infra/database"
	"order-package/internal/infra/repository/mongo"
)

func main() {
	server := infra.NewServer()
	cfg, err := config.NewConfig()
	if err != nil {
		panic(err)
	}
	collection := database.NewMongoCollection(cfg.Database)
	repository := mongo.NewMongoPackRepository(collection)
	optimalUseCase := optimalpackage.NewPackCombo(repository)
	optimalService := optimalpackage.NewOptimalPackageService(optimalUseCase)

	server.Start()
	handler := infra.NewHandle(*server)

	handler.Post("/packs-find", optimalService.Find)
	handler.Post("/packs-create", optimalService.Create)
	handler.Delete("/packs/:size", optimalService.Delete)

	handler.Run()
}
