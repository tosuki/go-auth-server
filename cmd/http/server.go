package http

import (
	"auth-server/cmd/http/auth"
	"auth-server/internal/auth/usecase"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RunServer(authUsecase *usecase.AuthUsecase) {
	r := gin.Default()

	r.Use(cors.Default())
	auth.AddAuthenticationRoutes(authUsecase, r.Group("/auth"))

	r.Run()
}
