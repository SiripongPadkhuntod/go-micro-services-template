package handler

import (
	"go-micro-services-template/pkg/dto"

	"github.com/gin-gonic/gin"
)

// HealthHTTP godoc
// @Summary Health Check
// @Description Returns service status
// @Tags Health
// @Produce json
// @Success 200 {object} dto.HealthResponse
// @Router /v1/health [get]
func (h *Handler) HealthHTTP(c *gin.Context) (any, error) {
	if err := h.svc.Health(c.Request.Context()); err != nil {
		return nil, err
	}

	return &dto.HealthResponse{
		Status: dto.HealthStatusUp,
	}, nil
}
