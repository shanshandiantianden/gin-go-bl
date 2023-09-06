package middlewares

import (
	"gin-go-bl/internal/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UnifiedResponseMiddleware 是处理统一HTTP响应格式的中间件
// 该中间件将在将响应发送给客户端之前拦截响应，并根据你指定的格式进行格式化。

// 返回值：
//   gin.HandlerFunc：Gin中间件处理函数

func UnifiedResponseMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		// 检查是否在处理请求时发生了错误
		// 如果发生了错误，通过ErrorResponse函数创建一个错误响应，并返回给客户端
		if len(c.Errors) > 0 {
			err := c.Errors.Last()
			Models.ErrorResponse(c, http.StatusInternalServerError, err.Error())
			return
		}

		// 检查是否设置了响应状态码
		// 如果未设置响应状态码，默认将状态码设置为200（OK）
		if c.Writer.Status() == 0 {
			c.Writer.WriteHeader(http.StatusOK)
		}

		// 如果没有错误，则格式化响应
		// 检查是否设置了"response_data"键的值，如果有，则调用SuccessResponse函数创建一个成功响应，并返回给客户端
		if c.Writer.Status() >= http.StatusOK && c.Writer.Status() < http.StatusMultipleChoices {
			data, exists := c.Get("response_data")
			if exists {
				Models.SuccessResponse(c, c.Writer.Status(), data)
				return
			}
		}
	}
}
