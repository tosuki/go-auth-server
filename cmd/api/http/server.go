package http

import (
	"go-auth-server/internal/auth/usecase"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitializeServer(authUsecase *usecase.AuthUsecase) {
	engine := gin.Default()

	engine.Use(cors.Default())
	AddRoutes(engine.Group("/api"), authUsecase)

	engine.Run()
}
