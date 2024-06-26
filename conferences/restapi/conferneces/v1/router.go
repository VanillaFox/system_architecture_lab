package v1

import (
	"github.com/VanillaFox/system_architecture_lab/conferences/restapi/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	router *gin.Engine,
	conferenceHandler *ConferenceHandler,
) {
	apiV1 := router.Group("/api/v1/conferences", HanleErrors)
	apiV1.Use(middlewares.AuthMiddleware)

	apiV1.POST("", conferenceHandler.createConference)
	apiV1.GET("", conferenceHandler.getConferences)
	apiV1.PUT(":object_id", conferenceHandler.updateConference)
	apiV1.DELETE(":object_id", conferenceHandler.deleteConference)
}
