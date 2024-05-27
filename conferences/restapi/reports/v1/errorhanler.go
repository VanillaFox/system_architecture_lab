package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HanleErrors(c *gin.Context) {
	c.Next()

	err := c.Errors.Last()
	if err == nil {
		return
	}

	c.JSON(http.StatusInternalServerError, err)
}
