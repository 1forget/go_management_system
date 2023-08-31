package bean

import (
	"github.com/gin-gonic/gin"
)

type CommonResult struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

// 因为没有泛型 存入token必须新加一个Response
type WithTokenResult struct {
	CommonResult
	Token string `json:"token"`
}

func ResponseError(c *gin.Context, code int, message string) {
	c.JSON(code, CommonResult{Success: false, Message: message})
}

func ResponseSuccess(c *gin.Context, message string) {
	c.JSON(200, CommonResult{Success: true, Message: message})
}

func ResponseWithToken(c *gin.Context, message string, token string) {
	c.JSON(201, WithTokenResult{CommonResult{Success: true, Message: message}, token})
}
