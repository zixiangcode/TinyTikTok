package service

import (
	"TinyTikTok/models"
	"github.com/gin-gonic/gin"
)

type UserService interface {
	GetUserById(Id int64) (models.User, error) // 通过 Id 查询用户

	GetUserByName(name string) (models.User, error) // 通过 Name 查询用户

	SaveUser(user models.User) error // 将用户存储在数据库中

	Register(username string, password string, context *gin.Context) error // 注册

	Login(username string, password string, context *gin.Context) error // 登录

	UserInfo(userId int64, token string) (*models.User, error) // 通过用户 ID 获取用户信息
}
