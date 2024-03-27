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
	apiV1.GET("username/:username", usersHandler.getUserByUsername)
	apiV1.GET("fullname/:full_name_prefix", usersHandler.getUserByFullName)
	apiV1.GET("", usersHandler.getUsers)
	apiV1.PUT(":username", usersHandler.updateUser)
	apiV1.DELETE(":username", usersHandler.deleteUser)
}
