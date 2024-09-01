package http

import (
	"go-auth-server/cmd/api/http/session/controller"
	"go-auth-server/internal/auth/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddRoutes(rg *gin.RouterGroup, authUsecase *usecase.AuthUsecase) error {
	sessionController := &controller.SessionController{
		AuthUsecase: authUsecase,
	}

	rg.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"ok":      true,
			"message": "hello world",
		})
	})

	rg.POST("/session/signup", sessionController.SignUp)
	rg.POST("/session/signin", sessionController.SignIn)
	rg.GET(
		"/session/rewoke",
		sessionController.Rewoke,
	)

	return nil
}
