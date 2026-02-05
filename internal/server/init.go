package server

import (
	"go-micro-services-template/internal/app/port"
	"go-micro-services-template/internal/app/service"
	"go-micro-services-template/internal/router"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initInfrastructure(ctx any) *port.Adapter {
	// init infrastructure here
	return &port.Adapter{}
}

func initRepository(ctx any) *port.Repository {
	// init repositories here
	return &port.Repository{}
}

func initService(adpt port.Adapter, repo port.Repository) port.Service {
	return service.New(adpt, repo)
}

func initServer(h port.Handler) *gin.Engine {
	r := gin.Default()

	// setup router
	rt := router.NewRouter(h)
	rt.Setup(r)

	// swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}

