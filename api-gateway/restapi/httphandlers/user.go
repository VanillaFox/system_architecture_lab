package httphandlers

import (
	_ "github.com/VanillaFox/system_architecture_lab/api-gateway/models"
)

// @Summary Создать пользователя
// @tags users
// @Accept json
// @Produce json
// @Param user body models.User true "user body"
// @Success 200
// @Failure 500
// @Router /api/v1/users [post]
func createUser() {
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
func getUserByUsername() {
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
func getUserByFullName() {
}

// @Summary Получить список всех пользователей
// @tags users
// @Accept json
// @Produce json
// @Param Authorization header string true "Auth"
// @Success 200 {object} models.Users
// @Failure 500
// @Router /api/v1/users/ [get]
func getUsers() {
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
func updateUser() {
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
func deleteUser() {
}
