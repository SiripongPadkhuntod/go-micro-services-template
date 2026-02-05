package port

import (
	"github.com/gin-gonic/gin"
)

type Handler interface {
	// Health(c *gin.Context) func(ctx context.Context, _ dto.EmptyStruct) (*dto.HealthResponse, error)
	HealthHTTP(c *gin.Context) (any, error)
}
