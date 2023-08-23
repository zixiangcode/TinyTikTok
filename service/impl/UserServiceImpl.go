package impl

import (
	"TinyTikTok/db"
	"TinyTikTok/models"
	"TinyTikTok/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
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

func (userService UserServiceImpl) SaveUser(user models.User) (int64, error) {
	result := db.GetMysqlDB().Create(&user)
	if result.Error != nil {
		log.Printf("方法 SaveUser() 失败 %v", result.Error)
		return user.Id, result.Error
	}
	return user.Id, result.Error
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

	userID, err := userService.SaveUser(newUser) // 存到数据库中
	if err != nil {
		context.JSON(http.StatusInternalServerError, models.UserLoginResponse{
			Response: models.Response{StatusCode: 1, StatusMsg: "无法保存用户！"},
		})
	} else { //注册界面需要返回token,因为注册之后会直接登录，之后会直接使用token进行操作
		token, err := utils.GenerateToken(username, utils.JWTCommonEntity{Id: userID,
			CreateTime: newUser.CreateTime, IsDeleted: newUser.IsDeleted})
		if err != nil {
			context.JSON(http.StatusInternalServerError, models.UserRegisterResponse{
				StatusCode: 1,
				StatusMsg:  "无法保存用户！",
			})
		}
		context.JSON(http.StatusOK, models.UserRegisterResponse{
			StatusCode: 0,
			StatusMsg:  "Register successful!",
			UserID:     userID,
			Token:      token,
		})
		return nil
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
	// 生成token
	token, err := utils.GenerateToken(username, utils.JWTCommonEntity{Id: user.Id,
		CreateTime: user.CreateTime, IsDeleted: user.IsDeleted})
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

func (userService UserServiceImpl) GetUserByIds(userIDs []int64) (users []models.User, err error) {
	err = db.GetMysqlDB().Where("id in  (?)", userIDs).Find(&users).Error
	if err != nil {
		log.Printf("方法 GetUserById() 失败 %v", err)
		return
	}
	return
}

//UpdateFollowTotalCount 更新数据关注总数
func (userService UserServiceImpl) UpdateFollowTotalCount(db *gorm.DB, userID int64, count int) (err error) {
	err = db.Model(&models.User{}).Where("id = ?", userID).Update("follow_count", gorm.Expr("follow_count + ? ", count)).Error
	if err != nil {
		log.Printf("方法 UpdateFollowTotalCount() 失败 %v", err)
		return
	}
	return
}

//UpdateFollowerTotalCount 更新用户粉丝数
func (userService UserServiceImpl) UpdateFollowerTotalCount(db *gorm.DB, userID int64, count int) (err error) {
	err = db.Model(&models.User{}).Where("id = ?", userID).Update("follower_count", gorm.Expr("follower_count + ? ", count)).Error
	if err != nil {
		log.Printf("方法 UpdateFollowerTotalCount() 失败 %v", err)
		return
	}
	return
}
