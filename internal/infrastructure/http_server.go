package infrastructure

import (
	"log"

	"github.com/gin-gonic/gin"
	"go-micro-services-template/internal/property"
	"go-micro-services-template/internal/router"
)

func StartHTTPServer(cfg *property.Config, r router.Router) error {
	engine := gin.New()
	engine.Use(gin.Recovery())

	r.Setup(engine)

	log.Println("service started on :" + cfg.Port)
	return engine.Run(":" + cfg.Port)
}
