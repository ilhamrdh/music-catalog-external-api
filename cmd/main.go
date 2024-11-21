package cmd

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/ilhamrdh/music-catalog-external-api/internal/configs"
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

	r := gin.Default()

	r.Run(cfg.Service.Port)
}
