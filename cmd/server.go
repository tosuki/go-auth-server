package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
    
    "auth-server/cmd/auth"
)

func RunServer() {
    r := gin.Default()

    r.GET("/ok", func(context *gin.Context) {
        context.JSON(http.StatusOK, gin.H{
            "ok": true,
            "message": "hello world",
        })
    })

    auth.AddAuthRoutes(r.Group("/auth"))

    r.Run()
}
