package ginmiddleware

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func setupLogging(duration time.Duration, logger *zap.Logger) {
	go func() {
		for range time.Tick(duration) {
			_ = logger.Sync()
		}
	}()
}

func AccessLogger(duration time.Duration, logger *zap.Logger) gin.HandlerFunc {
	setupLogging(duration, logger)

	return func(c *gin.Context) {
		startTime := time.Now()

		c.Next()

		latency := time.Since(startTime)
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()

		// log_format main '$remote_addr - [$time_local] "$request" '
		//                '$status $request_time "$http_referer"'
		//                '"$http_user_agent" "$http_x_forwarded_for" $body_bytes_sent';
		logger.Info(fmt.Sprintf("%s - [%s] \"%s %s %s\" %d %.3fms \"%s\" \"%s\" \"%s\" %d",
			clientIP,
			startTime.Format(time.RFC3339),
			method,
			c.Request.RequestURI,
			c.Request.Proto,
			statusCode,
			latency.Seconds()*1000,
			c.Request.Referer(),
			c.Request.UserAgent(),
			c.GetHeader("X-Forwarded-For"),
			c.Writer.Size(),
		))
	}

}
