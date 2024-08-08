package auth

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func AddAuthRoutes(routerGroup *gin.RouterGroup) {
    routerGroup.GET("/ok", func(context *gin.Context) {
        context.JSON(http.StatusOK, gin.H{
            "ok": true,
        })
    })
}
