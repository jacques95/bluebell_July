package controller

import (
	"bluebell/logic"
	"bluebell/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

func SignUpHandler(c *gin.Context) {
	// 参数接收并校验
	var p models.ParamSignUp
	if err := c.ShouldBindJSON(&p); err != nil {
		zap.L().Error("SignUp with invalid param", zap.Error(err))
		c.JSONP(http.StatusOK, gin.H{
			"msg": "invalid param.",
		})
		return
	}
	// 业务处理
	logic.Signup()
	// 返回响应
	c.JSONP(http.StatusOK, "ok")
}
