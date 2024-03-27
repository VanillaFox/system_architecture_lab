package restapi

import (
	v1 "github.com/VanillaFox/system_architecture_lab/users/restapi/v1"
	"github.com/gin-gonic/gin"
)

type RouteHandler struct {
	userHandler *v1.UserHandler
}

func NewRouter(
	userHandler *v1.UserHandler,
) RouteHandler {
	return RouteHandler{
		userHandler: userHandler,
	}
}

func (rh RouteHandler) SetRoutes(r *gin.Engine) {
	v1.RegisterRoutes(r, rh.userHandler)
}
