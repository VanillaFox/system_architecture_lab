package httphandlers

import _ "github.com/VanillaFox/system_architecture_lab/api-gateway/models"

// @Summary Создать конференцию
// @tags conferences
// @Accept json
// @Produce json
// @Param conference body models.ConferenceCreateModel true "conference body"
// @Param Authorization header string true "Auth"
// @Success 200
// @Failure 500
// @Router /api/v1/conferences [post]
func createConference() {
}

// @Summary Получение списка конференций
// @tags conferences
// @Accept json
// @Produce json
// @Param Authorization header string true "Auth"
// @Success 200 {object} models.Conferences
// @Failure 500
// @Router /api/v1/conferences [get]
func getConferences() {
}

// @Summary Обновление конференции
// @tags conferences
// @Accept json
// @Produce json
// @Param object_id path string true "ID конференции"
// @Param conference body models.ConferenceCreateModel true "conference body"
// @Param Authorization header string true "Auth"
// @Success 200
// @Failure 500
// @Router /api/v1/conferences/{object_id} [put]
func updateConference() {
}

// @Summary Удаление конференции
// @tags conferences
// @Accept json
// @Produce json
// @Param object_id path string true "ID конференции"
// @Param Authorization header string true "Auth"
// @Success 200
// @Failure 500
// @Router /api/v1/conferences/{object_id} [delete]
func deleteConference() {
}
