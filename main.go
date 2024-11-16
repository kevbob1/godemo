package main

import (
	"log/slog"
	"net/http"
	"os"

	sloggin "github.com/samber/slog-gin"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	router := gin.New()
	router.Use(sloggin.New(logger))
	router.Use(gin.Recovery())

	router.SetTrustedProxies(nil)
	router.GET("/health", healthCheck)
	router.GET("/hello", helloWorld)

	logger.Info("Server started", "port", 3000)
	router.Run(":3000")
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "healthy",
	})
}

func helloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, World!",
	})
}
