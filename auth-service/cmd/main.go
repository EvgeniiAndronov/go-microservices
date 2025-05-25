package main

import (
	"github.com/EvgeniiAndronov/auth-service/internal/api"
	"github.com/EvgeniiAndronov/auth-service/internal/config"
	"github.com/EvgeniiAndronov/auth-service/internal/models"
	"github.com/EvgeniiAndronov/auth-service/internal/repository"
	"github.com/EvgeniiAndronov/auth-service/pkg/database"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	dbCfg := config.LoadDBConfig()

	db, err := database.NewPostgresConnection(dbCfg)
	if err != nil {
		log.Fatalf("Failed to connection db! Error:\n%v", err)
	}
	log.Println("Successfully connected to db!")

	repository.InitDB(db)
	if err := db.AutoMigrate(&models.User{}); err != nil {
		log.Fatalf("Failed to migrate table! Error:\n%v", err)
	}

	router := api.SetupRouter()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	log.Println("Starting server on 8080")
	log.Fatal(router.Run("0.0.0.0:8080"))
}
