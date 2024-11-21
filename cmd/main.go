package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ilhamrdh/music-catalog-external-api/internal/configs"
	membershipsHandler "github.com/ilhamrdh/music-catalog-external-api/internal/handler/memberships"
	"github.com/ilhamrdh/music-catalog-external-api/internal/models/memberships"
	membershipsRepo "github.com/ilhamrdh/music-catalog-external-api/internal/repositories/memberships"
	membershipsSvc "github.com/ilhamrdh/music-catalog-external-api/internal/services/memberships"
	"github.com/ilhamrdh/music-catalog-external-api/pkg/internalsql"
)

func main() {
	var cfg *configs.Config

	err := configs.Init(
		configs.WithConfigFolder([]string{"./internal/configs"}),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)
	if err != nil {
		log.Fatal("Gagal membaca config")
	}

	cfg = configs.Get()

	db, err := internalsql.Connect(cfg.Database.DatabaseSourceName)
	if err != nil {
		log.Fatal("Gagal inisial database", err)
	}
	db.AutoMigrate(&memberships.User{})
	r := gin.Default()

	membershipRepo := membershipsRepo.NewReporitory(db)
	mempershipService := membershipsSvc.NewService(cfg, membershipRepo)
	membershipHandler := membershipsHandler.NewHandler(r, mempershipService)
	membershipHandler.RegisterRoute()

	r.Run(cfg.Service.Port)
}
