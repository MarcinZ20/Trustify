package server

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
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
	serverConfig.Router.Use(
		cors.New(cors.Config{
			AllowOrigins:  []string{"http://localhost:5173"},
			AllowMethods:  []string{"GET", "POST"},
			AllowHeaders:  []string{"Origin", "Content-Type", "Accept"},
			ExposeHeaders: []string{"Content-Length"},
			MaxAge:        12 * time.Hour,
		}))

	serverConfig.Server = http.Server{
		Addr:           ":8080",
		Handler:        serverConfig.Router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 10 << 20,
	}

	return serverConfig
}
