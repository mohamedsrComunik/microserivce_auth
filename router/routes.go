package router

import (
	"github.com/mohamedsrComunik/microserivce_auth/handler"

	"github.com/gin-gonic/gin"
)

type router struct {
	handler handler.Handler
}

// New to create new instance of router struct with deps
func New(handler handler.Handler) router {
	return router{
		handler,
	}
}

func (rt router) SetRoutes() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/auth-srv/api/v1/")
	{
		v1.POST("/auth", rt.handler.Auth)
		v1.POST("/user", rt.handler.Create)
		v1.GET("/valid", rt.handler.ValidateToken)
	}
	return r
}
