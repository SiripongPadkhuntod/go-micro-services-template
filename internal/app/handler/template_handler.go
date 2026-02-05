package handler

import (
	"context"
	"fmt"
	"log"

	"go-micro-services-template/pkg/dto"

	"github.com/gin-gonic/gin"
)



func (h *Handler) Health(_ *gin.Context) func(ctx context.Context, _ dto.EmptyStruct) (*dto.HealthResponse, error) {
	return func(ctx context.Context, _ dto.EmptyStruct) (*dto.HealthResponse, error) {
		log.Println("Health check invoked")
		if err := h.svc.Health(ctx); err != nil {
			fmt.Println("Health check failed:", err)
			return nil, err
		}

		// cfg := property.Get()

		return &dto.HealthResponse{
			Status: dto.HealthStatusUp,
			// Service: cfg.Server.ServiceName,
			// Version: cfg.Server.ServiceVersion,
		}, nil
	}
}


