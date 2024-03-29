// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/Krados/echoserver/internal/conf"
	"github.com/Krados/echoserver/internal/controller"
	"github.com/Krados/echoserver/internal/server"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Injectors from wire.go:

func initApp(cfg *conf.Config, sugar *zap.SugaredLogger) (*gin.Engine, error) {
	echoController := controller.NewEchoController()
	engine := server.NewHTTPServer(echoController)
	return engine, nil
}
