package server

import (
	"time"

	"github.com/Krados/echoserver/internal/controller"
	ginzap "github.com/gin-contrib/zap"
	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

func NewHTTPServer(echoController *controller.EchoController, logger *zap.Logger) *gin.Engine {
	r := gin.New()
	gin.SetMode(gin.ReleaseMode)
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	// Logs all panic to error log
	//   - stack means whether output the stack info.
	r.Use(ginzap.RecoveryWithZap(logger, true))
	v1 := r.Group("api").Group("v1")
	{
		v1.GET("echo", echoController.Ping)
	}

	return r
}
