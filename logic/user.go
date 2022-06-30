package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/snowflake"
)

func Signup(p *models.ParamSignUp) {
	// 判断用户是否已存在
	mysql.Exit()
	// 生成 UserId
	snowflake.GenID()
	// 生成 User 实例
	// 持久化
	mysql.InsterUser()
}
