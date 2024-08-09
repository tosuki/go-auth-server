package auth

import (
	"auth-server/internal/auth"
	"auth-server/internal/auth/usecase"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authUsecase *usecase.AuthUsecase
}

type RequestBody interface {
	Validate() bool
}

type AuthenticateRequestBody struct {
	Email    string
	Password string
}

func (requestBody *AuthenticateRequestBody) Validate() bool {
	return requestBody.Email != "" && requestBody.Password != ""
}

func (controller *AuthController) Authenticate(context *gin.Context) {
	var requestBody AuthenticateRequestBody

	if err := context.BindJSON(&requestBody); err != nil && !requestBody.Validate() {
		context.JSON(http.StatusBadRequest, gin.H{
			"ok":      false,
			"message": "bad-request",
		})
		return
	}

	token, authenticationError := controller.authUsecase.Authenticate(requestBody.Email, []byte(requestBody.Password))

	if authenticationError != nil {
		switch {
			case errors.Is(authenticationError, auth.ErrorInvalidEmail):
			case errors.Is(authenticationError, auth.ErrorInvalidPassword):
				context.JSON(http.StatusUnauthorized, gin.H{
					"ok":      false,
					"message": authenticationError.Error(),
				})
				return
			default:
				fmt.Printf("An unsupported error occurred: %s", authenticationError)
				context.JSON(http.StatusInternalServerError, gin.H{
					"ok":      false,
					"message": "internal-server-error",
				})
				return
		}
	}

	context.JSON(http.StatusAccepted, gin.H{
		"ok":      true,
		"message": "accepted",
		"data":    token,
	})
}
