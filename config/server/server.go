package server

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ServerConfig struct {
	Server http.Server
	Router *gin.Engine
}

var serverConfig *ServerConfig

func GetConfig() *ServerConfig {
	if serverConfig != nil {
		return serverConfig
	}

	serverConfig = &ServerConfig{}
	serverConfig.Router = gin.Default()
	serverConfig.Server = http.Server{
		Addr:           ":8080",
		Handler:        serverConfig.Router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 10 << 20,
	}

	return serverConfig
}
