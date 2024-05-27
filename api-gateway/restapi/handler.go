package restapi

import (
	"github.com/VanillaFox/system_architecture_lab/api-gateway/restapi/middlewares"
	"github.com/VanillaFox/system_architecture_lab/api-gateway/restapi/proxy"
	"github.com/gin-gonic/gin"
)

type RouteHandler struct {
	conferencesHandler *proxy.ConferencesHandler
	usersHandler       *proxy.UsersHandler
}

func NewRouter(
	conerencesHandler *proxy.ConferencesHandler,
	usersHandler *proxy.UsersHandler,
) RouteHandler {
	return RouteHandler{
		conferencesHandler: conerencesHandler,
		usersHandler:       usersHandler,
	}
}

func (rh RouteHandler) SetRoutes(router *gin.Engine) {
	// conference group
	conferenceGroup := router.Group("/api/v1/conferences", middlewares.HanleErrors)

	conferenceGroup.POST("", rh.conferencesHandler.Handle)
	conferenceGroup.GET("", rh.conferencesHandler.Handle)
	conferenceGroup.PUT(":object_id", rh.conferencesHandler.Handle)
	conferenceGroup.DELETE(":object_id", rh.conferencesHandler.Handle)

	// report grout
	reportGroup := router.Group("/api/v1", middlewares.HanleErrors)

	reportGroup.POST("conferences/:conference_object_id/reports", rh.conferencesHandler.Handle)
	reportGroup.GET("conferences/:conference_object_id/reports", rh.conferencesHandler.Handle)
	reportGroup.PUT("reports/:report_object_id", rh.conferencesHandler.Handle)
	reportGroup.DELETE("reports/:report_object_id", rh.conferencesHandler.Handle)

	// auth
	authGroup := router.Group("/api/auth", middlewares.HanleErrors)

	authGroup.GET("", rh.usersHandler.Handle)

	// user group without auth
	userGroup := router.Group("/api/v1/users", middlewares.HanleErrors)

	userGroup.POST("", rh.usersHandler.Handle)

	// user group with auth
	userSecuredGroup := router.Group("/api/v1/users", middlewares.HanleErrors)

	userSecuredGroup.GET("username/:username", rh.usersHandler.Handle)
	userSecuredGroup.GET("fullname/:full_name_prefix", rh.usersHandler.Handle)
	userSecuredGroup.GET("", rh.usersHandler.Handle)
	userSecuredGroup.PUT(":username", rh.usersHandler.Handle)
	userSecuredGroup.DELETE(":username", rh.usersHandler.Handle)
}
