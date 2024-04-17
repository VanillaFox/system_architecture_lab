package restapi

import (
	conferenceV1 "github.com/VanillaFox/system_architecture_lab/conferences/restapi/conferneces/v1"
	reportV1 "github.com/VanillaFox/system_architecture_lab/conferences/restapi/reports/v1"
	"github.com/gin-gonic/gin"
)

type RouteHandler struct {
	conferenceHandler *conferenceV1.ConferenceHandler
	reportHandler     *reportV1.ReportHandler
}

func NewRouter(
	conferenceHandler *conferenceV1.ConferenceHandler,
	reportHandler *reportV1.ReportHandler,
) RouteHandler {
	return RouteHandler{
		conferenceHandler: conferenceHandler,
		reportHandler:     reportHandler,
	}
}

func (rh RouteHandler) SetRoutes(r *gin.Engine) {
	conferenceV1.RegisterRoutes(r, rh.conferenceHandler)
	reportV1.RegisterRoutes(r, rh.reportHandler)
}
