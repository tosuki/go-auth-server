package middlware

import (
	"go-auth-server/pkg/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewSessionMiddlware(context *gin.Context) {
	authorizationToken := context.GetHeader("authorization")

	if authorizationToken == "" {
		utils.RefuseHTTPRequestWithMessage(context, http.StatusUnauthorized, "unauthorized")
		return
	}

	context.Next()
}
