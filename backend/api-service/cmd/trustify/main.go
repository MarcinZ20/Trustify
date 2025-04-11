package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/MarcinZ20/Trustify/backend/api-service/api/routes"
	"github.com/MarcinZ20/Trustify/backend/api-service/config/server"
)

func main() {
	apiServer := server.GetConfig()
	routes.ServerRoutes(apiServer.Router)

	go func() {
		if err := apiServer.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)

	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Server shutdown ...")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	if err := apiServer.Server.Shutdown(ctx); err != nil {
		log.Printf("Server shutdown due to an error: %s\n", err)
	}

	<-ctx.Done()
	log.Println("Server down")
}
