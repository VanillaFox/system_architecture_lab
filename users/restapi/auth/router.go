package auth

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	router *gin.Engine,
	authHandler *AuthHandler,
) {
	authGroup := router.Group("/api/auth", HanleErrors)
	authGroup.GET("", authHandler.auth)
}
