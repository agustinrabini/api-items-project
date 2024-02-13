package handlers

import (
	"github.com/agustinrabini/api-items-project/cmd/api/config"

	"github.com/agustinrabini/go-toolkit/goutils/logger"
	"github.com/gin-gonic/gin"
)

func LoggerHandler(requestName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		reqLogger := logger.NewRequestLogger(c, requestName, config.LogRatio, config.LogBodyRatio)
		c.Next()
		reqLogger.LogResponse(c)
	}
}
