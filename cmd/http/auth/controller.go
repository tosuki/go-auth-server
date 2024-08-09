package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthenticateRequestBody struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Authenticate(context *gin.Context) {
	var requestBody AuthenticateRequestBody

	if err := context.BindJSON(&requestBody); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"ok":      false,
			"message": "bad-request",
		})

		return
	}

	context.JSON(http.StatusOK, gin.H{
		"ok":   true,
		"data": requestBody,
	})
}
