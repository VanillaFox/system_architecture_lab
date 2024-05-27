package proxy

import (
	"io"
	"net/http"
	"net/http/httputil"

	"github.com/VanillaFox/system_architecture_lab/api-gateway/pkg/circuitbreaker"
	"github.com/gin-gonic/gin"
)

func NewUsersHandler(proxy *httputil.ReverseProxy, circuitBreaker *circuitbreaker.CircuitBreaker) *UsersHandler {
	return &UsersHandler{
		proxy:          proxy,
		circuitBreaker: circuitBreaker,
	}
}

type UsersHandler struct {
	proxy          *httputil.ReverseProxy
	circuitBreaker *circuitbreaker.CircuitBreaker
}

func (uh *UsersHandler) Handle(c *gin.Context) {
	uh.proxy.Director(c.Request)

	resp, err := uh.circuitBreaker.RoundTrip(c.Request)
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": "service unavailable"})
		return
	}

	if resp == nil {
		c.Status(http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	b, err := io.ReadAll(resp.Body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
	}

	c.Status(resp.StatusCode)
	c.Writer.Write(b)
}
