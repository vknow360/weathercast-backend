package main

import (
	"vknow360/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	r := routes.RegisterRoutes()
	r.Run(":8080")
}
