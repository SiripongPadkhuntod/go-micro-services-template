package router

import (
	"go-micro-services-template/internal/app/port"

	"github.com/gin-gonic/gin"
)

type Router interface {
	Setup(*gin.Engine)
}

type router struct {
	template port.Handler
}

func NewRouter(h port.Handler) Router {
	return &router{template: h}
}

// func (r *router) Setup(g *gin.Engine) {
// 	v1 := g.Group("/template-service/v1")
// 	{
// 		v1.GET("/health", BindReqJson200Resp(func(c *gin.Context) (any, error) {
// 			handler := r.template.Health(c)
// 			return handler(c.Request.Context(), dto.EmptyStruct{})
// 		}))
// 	}
// }

func (r *router) Setup(g *gin.Engine) {
	v1 := g.Group("/v1")
	{
		v1.GET("/health",
			BindReqJson200Resp(r.template.HealthHTTP),
		)
	}
}
