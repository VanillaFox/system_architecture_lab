package v1

import (
	"github.com/VanillaFox/system_architecture_lab/conferences/restapi/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	router *gin.Engine,
	reportHandler *ReportHandler,
) {
	apiV1 := router.Group("/api/v1", HanleErrors)
	apiV1.Use(middlewares.AuthMiddleware)

	apiV1.POST("conferences/:conference_object_id/reports", reportHandler.createReport)
	apiV1.GET("conferences/:conference_object_id/reports", reportHandler.getReports)

	apiV1.PUT("reports/:report_object_id", reportHandler.updateReport)
	apiV1.DELETE("reports/:report_object_id", reportHandler.deleteReport)
}
