package models

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int    `json:"code"`    // 业务码
	Message string `json:"message"` // 响应消息
	Body    any    `json:"body"`    // 消息体
}

func ErrorResponse(c *gin.Context, code int, message string) {
	c.JSON(code, Response{
		Code:    code,
		Message: message,
		Body:    nil,
	})
}
func SuccessResponse(c *gin.Context, code int, data interface{}) {
	c.JSON(code, Response{
		Code:    code,
		Message: "success",
		Body:    data,
	})
}
