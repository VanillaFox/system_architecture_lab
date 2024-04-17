package v1

import (
	"net/http"

	"github.com/VanillaFox/system_architecture_lab/conferences/app/services"
	"github.com/VanillaFox/system_architecture_lab/conferences/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReportHandler struct {
	reportService *services.ReportService
}

func NewReportHandler(reportService *services.ReportService) *ReportHandler {
	return &ReportHandler{
		reportService: reportService,
	}
}

// @Summary Создать доклад
// @tags reports
// @Accept json
// @Produce json
// @Param conference body models.ReportCreateModel true "report body"
// @Param conference_object_id path string true "ID конференции"
// @Success 200
// @Failure 500
// @Router /api/v1/conferences/{conference_object_id}/reports [post]
func (h *ReportHandler) createReport(c *gin.Context) {
	ctx := c.Request.Context()

	objectIDParam := c.Param("conference_object_id")

	conferenceObjectID, err := primitive.ObjectIDFromHex(objectIDParam)
	if err != nil {
		_ = c.Error(err)

		return
	}

	var ReportCreateModel models.ReportCreateModel
	if err := c.BindJSON(&ReportCreateModel); err != nil {
		_ = c.Error(err)

		return
	}

	var report models.Report

	report.ConferenceObjectID = conferenceObjectID
	report.Title = ReportCreateModel.Title
	report.Description = ReportCreateModel.Description

	if err := h.reportService.CreateReport(ctx, &report); err != nil {
		_ = c.Error(err)

		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

// @Summary Получить доклады конференции
// @tags reports
// @Accept json
// @Produce json
// @Param conference_object_id path string true "ID конференции"
// @Success 200 {object} models.Reports
// @Failure 500
// @Router /api/v1/conferences/{conference_object_id}/reports [get]
func (h *ReportHandler) getReports(c *gin.Context) {
	ctx := c.Request.Context()

	objectIDParam := c.Param("conference_object_id")

	conferenceObjectID, err := primitive.ObjectIDFromHex(objectIDParam)
	if err != nil {
		_ = c.Error(err)

		return
	}
	result, err := h.reportService.GetReports(ctx, conferenceObjectID)

	if err != nil {
		_ = c.Error(err)

		return
	}

	c.JSON(http.StatusOK, result)
}

// @Summary Обновить доклад
// @tags reports
// @Accept json
// @Produce json
// @Param report_object_id path string true "ID доклада"
// @Param conference body models.ReportCreateModel true "report body"
// @Success 200
// @Failure 500
// @Router /api/v1/reports/{report_object_id} [put]
func (h *ReportHandler) updateReport(c *gin.Context) {
	ctx := c.Request.Context()

	objectIDParam := c.Param("report_object_id")

	reportObjectID, err := primitive.ObjectIDFromHex(objectIDParam)
	if err != nil {
		_ = c.Error(err)

		return
	}

	var ReportCreateModel models.ReportCreateModel
	if err := c.BindJSON(&ReportCreateModel); err != nil {
		_ = c.Error(err)

		return
	}

	var report models.Report
	report.Title = ReportCreateModel.Title
	report.Description = ReportCreateModel.Description

	err = h.reportService.UpdateReport(ctx, reportObjectID, &report)

	if err != nil {
		_ = c.Error(err)

		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

// @Summary Удалить доклад
// @tags reports
// @Accept json
// @Produce json
// @Param report_object_id path string true "ID доклада"
// @Success 200
// @Failure 500
// @Router /api/v1/reports/{report_object_id} [delete]
func (h *ReportHandler) deleteReport(c *gin.Context) {
	ctx := c.Request.Context()

	objectIDParam := c.Param("report_object_id")

	reportObjectID, err := primitive.ObjectIDFromHex(objectIDParam)
	if err != nil {
		_ = c.Error(err)

		return
	}

	err = h.reportService.DeleteReport(ctx, reportObjectID)

	if err != nil {
		_ = c.Error(err)

		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
