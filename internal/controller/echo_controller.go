package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type EchoController struct{}

func NewEchoController() *EchoController {
	return &EchoController{}
}

func (s *EchoController) Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"msg":    "pong",
	})
}
