package middlewares

import (
	"fmt"
	"golang-e-wallet-rest-api/internal/pkgs/utils/logger"
	"time"

	"github.com/gin-gonic/gin"
)

func LoggerMiddleware(c *gin.Context) {
	log := logger.Log
	startTime := time.Now()
	c.Next()
	endTime := time.Now()

	latencyTime := endTime.Sub(startTime)
	reqMethod := c.Request.Method
	reqUri := c.Request.RequestURI
	statusCode := c.Writer.Status()
	clientIP := c.ClientIP()

	fmt.Printf("ERRORS: %+v\n", c.Errors)

	if lastErr := c.Errors.Last(); lastErr != nil {
		log.WithFields(map[string]any{
			"METHOD":    reqMethod,
			"URI":       reqUri,
			"STATUS":    statusCode,
			"LATENCY":   latencyTime,
			"CLIENT_IP": clientIP,
		}).Error(c.Errors)
		return
	}

	log.WithFields(map[string]any{
		"METHOD":    reqMethod,
		"URI":       reqUri,
		"STATUS":    statusCode,
		"LATENCY":   latencyTime,
		"CLIENT_IP": clientIP,
	}).Infof("REQUEST %s %s SUCCESS", reqMethod, reqUri)
}
