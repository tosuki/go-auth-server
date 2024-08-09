package http

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"auth-server/cmd/http/auth"
)

func RunServer() {
	r := gin.Default()

	r.Use(cors.Default())
	auth.AddAuthRoutes(r.Group("/auth"))

	r.Run()
}
