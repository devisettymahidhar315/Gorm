package router

import (
	"app/api"

	"github.com/gin-gonic/gin"
)

type Router struct {
	rou *api.Api
}

func Init(a *api.Api) *Router {
	return &Router{rou: a}

}

func (ro *Router) Routes(c *gin.Engine) {
	c.POST("/:id/:name", ro.rou.Post)
	c.DELETE("/:id", ro.rou.Del_key)
	c.DELETE("/all", ro.rou.Del_all)
	c.GET("/:key", ro.rou.Get_key)
	c.GET("/print", ro.rou.Get_Print)

}
