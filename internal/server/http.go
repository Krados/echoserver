package server

import (
	"echoserver/internal/controller"

	"github.com/gin-gonic/gin"
)

func NewHTTPServer(echoController *controller.EchoController) *gin.Engine {
	r := gin.Default()
	v1 := r.Group("api").Group("v1")
	{
		v1.GET("echo", echoController.Ping)
	}

	return r
}
