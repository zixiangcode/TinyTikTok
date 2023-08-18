package impl

import (
	"TinyTikTok/db"
	"TinyTikTok/models"
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

type UserServiceImpl struct {
}

func (userService UserServiceImpl) GetUserById(Id int64) (models.User, error) {
	var user models.User

	err := db.GetMysqlDB().Where("id = ? AND is_deleted != ?", Id, 1).First(&user).Error
	if err != nil {
		log.Printf("方法 GetUserById() 失败 %v", err)
		return user, err
	}
	return user, nil
}

func (userService UserServiceImpl) GetUserByName(name string) (models.User, error) {
	var user models.User

	err := db.GetMysqlDB().Where("name = ? AND is_deleted != ?", name, 1).First(&user).Error
	if err != nil {
		log.Printf("方法 GetUserByName() 失败 %v", err)
		return user, err
	}
	return user, nil
}

func (userService UserServiceImpl) SaveUser(user models.User) error {
	err := db.GetMysqlDB().Create(&user).Error
	if err != nil {
		log.Printf("方法 SaveUser() 失败 %v", err)
		return err
	}
	return nil
}

func (userService UserServiceImpl) Register(username string, password string, context *gin.Context) error {
	_, errName := userService.GetUserByName(username)
	if errName == nil { // 为空说明查询没报错，也就是查到了，因此用户名重复
		context.JSON(http.StatusBadRequest, models.UserLoginResponse{
			Response: models.Response{StatusCode: 1, StatusMsg: "用户名重复"},
		})
		return nil
	}

	// 密码加密
	encrypt, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	password = string(encrypt)

	newUser := models.User{ // 新增一个用户的实例对象，保存基础信息
		CommonEntity: models.NewCommonEntity(),
		Name:         username,
		Password:     password,
	}

	err := userService.SaveUser(newUser) // 存到数据库中
	if err != nil {
		context.JSON(http.StatusInternalServerError, models.UserLoginResponse{
			Response: models.Response{StatusCode: 1, StatusMsg: "无法保存用户！"},
		})
	} else {
		context.JSON(http.StatusOK, models.UserLoginResponse{
			Response: models.Response{StatusCode: 0},
			UserId:   newUser.Id,
		})
	}
	return nil
}

func (userService UserServiceImpl) Login(username string, password string, context *gin.Context) error {
	// 首先查询是否存在此用户
	user, err := userService.GetUserByName(username)
	if err != nil {
		context.JSON(http.StatusBadRequest, models.UserLoginResponse{
			Response: models.Response{StatusCode: 1, StatusMsg: "用户不存在，请注册!"},
		})
		return nil
	}
	// 对密码解密，并与从数据库查询到的信息做对比
	pwdErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if pwdErr != nil {
		context.JSON(http.StatusOK, models.UserLoginResponse{
			Response: models.Response{StatusCode: 1, StatusMsg: "密码错误！"},
		})
		return pwdErr
	}
	// TODO: 生成和解析 token 的函数待实现，放在 utils 层中
	token := ""
	// 登录成功
	context.JSON(http.StatusOK, models.UserLoginResponse{
		Response: models.Response{StatusCode: 0, StatusMsg: "登录成功！"},
		UserId:   user.Id,
		Token:    token,
	})
	return nil
}

func (userService UserServiceImpl) UserInfo(userId int64, token string) (*models.User, error) {
	user, err := userService.GetUserById(userId)
	if err != nil {
		return nil, errors.New("用户不存在！")
	}
	return &user, nil
}
