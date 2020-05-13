package ginmiddleware

import (
	"math/rand"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-una/una/tools"
)

func ReqId(c *gin.Context) {
	c.Set("reqId", tools.MD5(string(time.Now().UnixNano()) + string(rand.Intn(99999)))[:10])

	c.Next()
}
