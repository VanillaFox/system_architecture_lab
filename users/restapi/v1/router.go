package v1

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	router *gin.Engine,
	usersHandler *UserHandler,
) {
	apiV1 := router.Group("/api/v1/users", HanleErrors)

	apiV1.POST("", usersHandler.createUser)

	apiV1Secured := router.Group("/api/v1/users", HanleErrors)
	apiV1Secured.Use(AuthMiddleware)

	apiV1Secured.GET("username/:username", usersHandler.getUserByUsername)
	apiV1Secured.GET("fullname/:full_name_prefix", usersHandler.getUserByFullName)
	apiV1Secured.GET("", usersHandler.getUsers)
	apiV1Secured.PUT(":username", usersHandler.updateUser)
	apiV1Secured.DELETE(":username", usersHandler.deleteUser)
}
