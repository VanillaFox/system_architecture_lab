package auth

import (
	"fmt"
	"net/http"

	"github.com/VanillaFox/system_architecture_lab/users/app/services"
	"github.com/VanillaFox/system_architecture_lab/users/models"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authService *services.AuthService
}

var (
	ErrAuthFail = fmt.Errorf("authorization fail")
	ErrInternal = fmt.Errorf("internal server error")
)

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

// @Summary Аутентификация
// @tags auth
// @Accept json
// @Produce json
// @Param Authorization header string true "Auth"
// @Success 200
// @Failure 500
// @Router /api/auth [get]
func (h *AuthHandler) auth(c *gin.Context) {
	ctx := c.Request.Context()

	username, password, ok := c.Request.BasicAuth()

	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": ErrAuthFail.Error()})

		return
	}

	creds := models.Creds{
		Username: username,
		Password: models.NewPassword(password),
	}

	result, err := h.authService.Auth(ctx, &creds)

	if err != nil {
		if err == services.ErrInvalidUsernameOrPassword {
			c.JSON(http.StatusUnauthorized, gin.H{"error": ErrAuthFail.Error()})
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": ErrInternal.Error()})

		return
	}

	c.JSON(http.StatusOK, result)
}
