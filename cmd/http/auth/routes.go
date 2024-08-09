package auth

import (
	"auth-server/internal/auth/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddAuthenticationRoutes(authUsecase *usecase.AuthUsecase, routerGroup *gin.RouterGroup) {
	authHttpController := AuthHttpController{
		authUsecase: authUsecase,
	}

	routerGroup.POST("/authenticate", authHttpController.Authenticate)
	routerGroup.GET("/ok", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"ok":      true,
			"message": "ok!",
		})
	})
}
