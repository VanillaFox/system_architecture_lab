package v1

import (
	"net/http"

	"github.com/VanillaFox/system_architecture_lab/conferences/app/services"
	"github.com/VanillaFox/system_architecture_lab/conferences/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ConferenceHandler struct {
	conferenceService *services.ConferenceService
}

func NewConferenceHandler(conferenceService *services.ConferenceService) *ConferenceHandler {
	return &ConferenceHandler{
		conferenceService: conferenceService,
	}
}

// @Summary Создать конференцию
// @tags conferences
// @Accept json
// @Produce json
// @Param conference body models.ConferenceCreateModel true "conference body"
// @Success 200
// @Failure 500
// @Router /api/v1/conferences [post]
func (h *ConferenceHandler) createConference(c *gin.Context) {
	ctx := c.Request.Context()

	var conference *models.ConferenceCreateModel
	if err := c.BindJSON(&conference); err != nil {
		_ = c.Error(err)

		return
	}

	if err := h.conferenceService.CreateConference(ctx, conference); err != nil {
		_ = c.Error(err)

		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

// @Summary Получение списка конференций
// @tags conferences
// @Accept json
// @Produce json
// @Success 200 {object} models.Conferences
// @Failure 500
// @Router /api/v1/conferences [get]
func (h *ConferenceHandler) getConferences(c *gin.Context) {
	ctx := c.Request.Context()

	result, err := h.conferenceService.GetConferences(ctx)
	if err != nil {
		_ = c.Error(err)

		return
	}

	c.JSON(http.StatusOK, result)
}

// @Summary Обновление конференции
// @tags conferences
// @Accept json
// @Produce json
// @Param object_id path string true "ID конференции"
// @Param conference body models.ConferenceCreateModel true "conference body"
// @Success 200
// @Failure 500
// @Router /api/v1/conferences/{object_id} [put]
func (h *ConferenceHandler) updateConference(c *gin.Context) {
	ctx := c.Request.Context()

	objectIDParam := c.Param("object_id")

	objectID, err := primitive.ObjectIDFromHex(objectIDParam)
	if err != nil {
		_ = c.Error(err)

		return
	}

	var conference *models.ConferenceCreateModel
	if err := c.BindJSON(&conference); err != nil {
		_ = c.Error(err)

		return
	}

	err = h.conferenceService.UpdateConference(ctx, objectID, conference)
	if err != nil {
		_ = c.Error(err)

		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

// @Summary Удаление конференции
// @tags conferences
// @Accept json
// @Produce json
// @Param object_id path string true "ID конференции"
// @Success 200
// @Failure 500
// @Router /api/v1/conferences/{object_id} [delete]
func (h *ConferenceHandler) deleteConference(c *gin.Context) {
	ctx := c.Request.Context()

	objectIDParam := c.Param("object_id")

	objectID, err := primitive.ObjectIDFromHex(objectIDParam)
	if err != nil {
		_ = c.Error(err)

		return
	}

	err = h.conferenceService.DeleteConference(ctx, objectID)
	if err != nil {
		_ = c.Error(err)

		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
