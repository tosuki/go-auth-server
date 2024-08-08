package auth

import (
    "net/http"
    "github.com/gin-gonic/gin"
    
    authUsecase "auth-server/internal/auth/usecase"
)

func Authenticate(context *gin.Context) {
    usecase.Authenticate(context.)

    context.JSON(http.StatusOK, gin.H{
        "ok": true,
    })
}

    
}
