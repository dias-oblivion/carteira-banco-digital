package api

import (
	"net/http"
	"strings"

	"github.com/dias-oblivion/carteira-banco-digital/api/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var bearerToken string

		if authorizationHeader := ctx.Request.Header.Get("Authorization"); authorizationHeader == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "authentication header not found"})
			ctx.Abort()
			return
		}

		if strings.Contains(ctx.Request.Header.Get("Authorization"), "Bearer") {
			bearerToken = strings.Split(ctx.Request.Header.Get("Authorization"), " ")[1]
			err := utils.VerifyJWTToken(bearerToken)
			if err != nil {
				ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
				ctx.Abort()
				return
			}
		} else {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "invalid authentication header"})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
