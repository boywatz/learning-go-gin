package middleware

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] %s %s %d %s \n",
			params.ClientIP,
			params.TimeStamp,
			params.Method,
			params.Path,
			params.StatusCode,
			strconv.FormatFloat(params.Latency.Seconds(), 'f', 2, 64)+"sec")
	})
}
