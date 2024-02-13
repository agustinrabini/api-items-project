package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// HealthCheckerHandler struct provides the handler for a health check endpoint.
type HealthCheckerHandler struct{}

func NewHealthCheckerHandler() HealthCheckerHandler {
	return HealthCheckerHandler{}
}

// Ping is the handler of test app
// @Summary Ping
// @Description test if the router works correctly
// @Tags ping
// @Produce  json
// @Success 200
// @Router /ping [get]
func (h HealthCheckerHandler) Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
