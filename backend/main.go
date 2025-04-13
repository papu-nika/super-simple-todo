package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/papu-nika/super-simple-todo/backend/auth"
	"golang.org/x/exp/slog"
)

func main() {
	// Initialize slog logger
	logger := slog.New(slog.NewJSONHandler(os.Stdout))
	r := gin.Default()

	authHandler, err := auth.NewAuthHandler()
	if err != nil {
		logger.Error("Failed to initialize auth handler", slog.String("error", err.Error()))
		log.Fatalf("Failed to initialize auth handler: %v", err)
	}

	r.GET("/auth/google", authHandler.GoogleLogin)
	r.POST("/auth/google", authHandler.GoogleCallback)

	logger.Info("Starting server on :8080")
	if err := r.Run(); err != nil {
		logger.Error("Failed to start server", slog.String("error", err.Error()))
		log.Fatalf("Failed to start server: %v", err)
	}
}
