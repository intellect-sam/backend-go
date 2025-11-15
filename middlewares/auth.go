package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/intellect-sam/backend-go/utils"
)

func Authenticate(context *gin.Context) {
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}

	// to check for the invalid token
	userId, err := utils.VerifyToken(token)

	if err != nil {
		context.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid token"})
		return
	}

	context.Set("userId", userId)

	context.Next()

}
