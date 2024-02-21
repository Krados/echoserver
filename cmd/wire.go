//go:build wireinject
// +build wireinject

package main

import (
	"echoserver/internal/conf"
	"echoserver/internal/controller"
	"echoserver/internal/server"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
	"go.uber.org/zap"
)

func initApp(cfg *conf.Config, sugar *zap.SugaredLogger) (*gin.Engine, error) {
	panic(wire.Build(controller.ProviderSet, server.ProviderSet))
}
