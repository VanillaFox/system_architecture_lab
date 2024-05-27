package httphandlers

import (
	_ "github.com/VanillaFox/system_architecture_lab/api-gateway/models"
)

// @Summary Создать доклад
// @tags reports
// @Accept json
// @Produce json
// @Param conference body models.ReportCreateModel true "report body"
// @Param conference_object_id path string true "ID конференции"
// @Param Authorization header string true "Auth"
// @Success 200
// @Failure 500
// @Router /api/v1/conferences/{conference_object_id}/reports [post]
func createReport() {
}

// @Summary Получить доклады конференции
// @tags reports
// @Accept json
// @Produce json
// @Param conference_object_id path string true "ID конференции"
// @Param Authorization header string true "Auth"
// @Success 200 {object} models.Reports
// @Failure 500
// @Router /api/v1/conferences/{conference_object_id}/reports [get]
func getReports() {
}

// @Summary Обновить доклад
// @tags reports
// @Accept json
// @Produce json
// @Param report_object_id path string true "ID доклада"
// @Param conference body models.ReportCreateModel true "report body"
// @Param Authorization header string true "Auth"
// @Success 200
// @Failure 500
// @Router /api/v1/reports/{report_object_id} [put]
func updateReport() {
}

// @Summary Удалить доклад
// @tags reports
// @Accept json
// @Produce json
// @Param report_object_id path string true "ID доклада"
// @Param Authorization header string true "Auth"
// @Success 200
// @Failure 500
// @Router /api/v1/reports/{report_object_id} [delete]
func deleteReport() {
}
