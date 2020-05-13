package ginmiddleware

import (
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func RecoveryWithZap(logger *zap.Logger, stack bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				if brokenPipe {
					httpRequest, _ := httputil.DumpRequest(c.Request, false)
					logger.Error(c.Request.RequestURI,
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
					)
					_ = c.Error(err.(error))
					c.Abort()
					return
				}

				if stack {
					logger.Error(fmt.Sprintf("[Recovery from panic] %s", err),
						zap.String("req", c.Request.Method+" "+c.Request.RequestURI),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					logger.Error(fmt.Sprintf("[Recovery from panic] %s", err),
						zap.String("req", c.Request.Method+" "+c.Request.RequestURI),
					)
				}

				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()

		c.Next()
	}
}
