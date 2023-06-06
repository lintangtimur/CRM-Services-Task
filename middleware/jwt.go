package middleware

import (
	"CRM-Services-Task/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := utils.TokenValid(c)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func SuperadminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := utils.ExtractToken(c)
		claims, b := utils.ExtractClaims(token)

		if !b {
			c.JSON(400, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}
		if claims["role_id"] != 1 {
			//bukan superadmin
			fmt.Println(claims["role_id"])
			c.JSON(http.StatusUnauthorized, gin.H{"error": "kamu bukan superadmin"})
			c.Abort()
			return
		}

		c.Next()
	}
}
