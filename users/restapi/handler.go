package restapi

import (
	"github.com/VanillaFox/system_architecture_lab/users/restapi/auth"
	v1 "github.com/VanillaFox/system_architecture_lab/users/restapi/v1"
	"github.com/gin-gonic/gin"
)

type RouteHandler struct {
	userHandler *v1.UserHandler
	authHandler *auth.AuthHandler
}

func NewRouter(
	userHandler *v1.UserHandler,
	authHandler *auth.AuthHandler,
) RouteHandler {
	return RouteHandler{
		userHandler: userHandler,
		authHandler: authHandler,
	}
}

func (rh RouteHandler) SetRoutes(r *gin.Engine) {
	v1.RegisterRoutes(r, rh.userHandler)
	auth.RegisterRoutes(r, rh.authHandler)
}
