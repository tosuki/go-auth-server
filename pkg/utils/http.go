package utils

import "github.com/gin-gonic/gin"

func AnswerHTTPRequestWithData(ctx *gin.Context, status int, data any) {
	ctx.JSON(status, gin.H{
		"ok":   true,
		"data": data,
	})
}

func RefuseHTTPRequestWithMessage(ctx *gin.Context, status int, message string) {
	ctx.JSON(status, gin.H{
		"ok":      false,
		"message": message,
	})
}

func RefuseHTTPRequestWithError(ctx *gin.Context, status int, err error) {
	RefuseHTTPRequestWithMessage(ctx, status, err.Error())
}
