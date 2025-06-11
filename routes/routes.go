package routes

import (
	"os"
	"vknow360/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupCORS() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}

	if frontendURL := os.Getenv("FRONTEND_URL"); frontendURL != "" {
		config.AllowOrigins = append(config.AllowOrigins, frontendURL)
	}
	config.AllowMethods = []string{"GET"}

	return cors.New(config)
}

func RegisterRoutes() *gin.Engine {
	r := gin.Default()
	r.Use(SetupCORS())
	r.GET("/status", handlers.Status)
	r.GET("/news", handlers.GetNews)
	return r
}
