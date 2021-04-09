package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func CustomerLogger() gin.HandlerFunc  {
	return gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] %",
				params.ClientIP,
				params.TimeStamp.Format(time.RubyDate),
				params.Method,
				params.Path,
				params.StatusCode,
				params.Latency,
			)
	})
}