package auth

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func RunServer() {
    r := gin.Default()

    r.GET("/ok", func(context *gin.Context) {
        context.JSON(http.StatusOK, gin.H{
            "ok": true,
            "message": "hello world",
        })
    })

    r.Run()
}
