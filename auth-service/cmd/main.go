package main

import (
	"github.com/EvgeniiAndronov/auth-service/internal/config"
	"github.com/EvgeniiAndronov/auth-service/pkg/database"
	"log"
)

func main() {
	dbCfg := config.LoadDBConfig()

	db, err := database.NewPostgresConnection(dbCfg)
	if err != nil {
		log.Fatalf("Failed to connection db! Error:\n%v", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to get db connection! Error:\n%v", err)
	}

	if err := sqlDB.Ping(); err != nil {
		log.Fatalf("Failed to ping db! Error:\n%v", err)
	}

	log.Println("Successfully connected to db!")
}
