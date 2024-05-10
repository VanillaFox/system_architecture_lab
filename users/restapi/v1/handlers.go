package v1

import (
	"net/http"

	"github.com/VanillaFox/system_architecture_lab/users/app/services"
	"github.com/VanillaFox/system_architecture_lab/users/models"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// @Summary Создать пользователя
// @tags users
// @Accept json
// @Produce json
// @Param user body models.User true "user body"
// @Success 200
// @Failure 500
// @Router /api/v1/users [post]
func (h *UserHandler) createUser(c *gin.Context) {
	ctx := c.Request.Context()

	var user *models.User
	if err := c.BindJSON(&user); err != nil {
		_ = c.Error(err)

		return
	}

	if err := h.userService.CreateUser(ctx, user); err != nil {
		_ = c.Error(err)

		return
	}

	c.JSON(http.StatusOK, gin.H{})
}

// @Summary Получить пользователя по логину
// @tags users
// @Accept json
// @Produce json
// @Param username path string true "Логин"
// @Param Authorization header string true "Auth"
// @Success 200 {object} models.User
// @Failure 500
// @Router /api/v1/users/username/{username} [get]
func (h *UserHandler) getUserByUsername(c *gin.Context) {
	ctx := c.Request.Context()

	username := c.Param("username")

	result, err := h.userService.GetUserByUsername(ctx, username)
	if err != nil {
		_ = c.Error(err)

		return
	}

	c.JSON(http.StatusOK, result)
}

// @Summary Получить пользователя по маске ФИ
// @tags users
// @Accept json
// @Produce json
// @Param full_name_prefix path string true "Имя пользователя"
// @Param Authorization header string true "Auth"
// @Success 200 {object} models.Users
// @Failure 500
// @Router /api/v1/users/fullname/{full_name_prefix} [get]
func (h *UserHandler) getUserByFullName(c *gin.Context) {
	ctx := c.Request.Context()

	fullNamePrefix := c.Param("full_name_prefix")

	result, err := h.userService.GetUserByFullnamePrefix(ctx, fullNamePrefix)
	if err != nil {
		_ = c.Error(err)

		return
	}

	c.JSON(http.StatusOK, result)
}

// @Summary Получить список всех пользователей
// @tags users
// @Accept json
// @Produce json
// @Param Authorization header string true "Auth"
// @Success 200 {object} models.Users
// @Failure 500
// @Router /api/v1/users/ [get]
func (h *UserHandler) getUsers(c *gin.Context) {
	ctx := c.Request.Context()

	result, err := h.userService.GetUsers(ctx)
	if err != nil {
		_ = c.Error(err)

		return
	}

	c.JSON(http.StatusOK, result)
}

// @Summary Обновить пользователя
// @tags users
// @Accept json
// @Produce json
// @Param username path string true "Логин"
// @Param user body models.User true "user body"
// @Param Authorization header string true "Auth"
// @Success 200
// @Failure 500
// @Router /api/v1/users/{username} [put]
func (h *UserHandler) updateUser(c *gin.Context) {
	ctx := c.Request.Context()

	username := c.Param("username")

	var user *models.User
	if err := c.BindJSON(&user); err != nil {
		_ = c.Error(err)

		return
	}

	result, err := h.userService.UpdateUser(ctx, username, user)
	if err != nil {
		_ = c.Error(err)

		return
	}

	c.JSON(http.StatusOK, result)
}

// @Summary Удалить пользователя
// @tags users
// @Accept json
// @Produce json
// @Param username path string true "Логин"
// @Param Authorization header string true "Auth"
// @Success 200
// @Failure 500
// @Router /api/v1/users/{username} [delete]
func (h *UserHandler) deleteUser(c *gin.Context) {
	ctx := c.Request.Context()

	username := c.Param("username")

	err := h.userService.DeleteUser(ctx, username)
	if err != nil {
		_ = c.Error(err)

		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
