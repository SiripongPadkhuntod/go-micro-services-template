package main

import (
	"log"


	_ "go-micro-services-template/docs"


	"go-micro-services-template/internal/property"
	"go-micro-services-template/internal/server"
)

// @title Auth Service API
// @version 1.0
// @description Authentication microservice
// @host localhost:8080
// @BasePath /
func main() {
	cfg := property.Load()
	log.Printf("Configuration loaded: %+v\n", cfg)

	if err := server.RunServerWithConfig(cfg); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
