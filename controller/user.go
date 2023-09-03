package controller

import (
	"TinyTikTok/dao"
	"TinyTikTok/models"
	"TinyTikTok/service/impl"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin

var usersLoginInfo = map[string]models.User{
	"zhangleidouyin": {
		CommonEntity:  models.CommonEntity{Id: 1},
		Name:          "zhanglei",
		FollowCount:   10,
		FollowerCount: 5,
		IsFollow:      true,
	},
}

var userIdSequence = int64(1)

/*
	type UserLoginResponse struct {
		Response
		UserId int64  `json:"user_id,omitempty"`
		Token  string `json:"token"`
	}

	type UserResponse struct {
		Response
		User User `json:"user"`
	}
*/
func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	err := impl.UserServiceImpl{}.Register(username, password, c)
	if err != nil {
		log.Printf("Register Error!")

	}
	err = dao.UpdateUserFollowByUserName(username)
	if err != nil {
		log.Printf("更新 user 表的 is_follow 属性列失败")
	}
	/*
		token := username + password

		if _, exist := usersLoginInfo[token]; exist {
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 1, StatusMsg: "User already exist"},
			})
		} else {
			atomic.AddInt64(&userIdSequence, 1)
			newUser := User{
				Id:   userIdSequence,
				Name: username,
			}
			usersLoginInfo[token] = newUser
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 0},
				UserId:   userIdSequence,
				Token:    username + password,
			})
		}
	*/
}

func Login(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	err := impl.UserServiceImpl{}.Login(username, password, c)
	if err != nil {
		log.Printf("Login Error !")
	}
	err = dao.UpdateUserFollowByUserName(username)
	if err != nil {
		log.Printf("更新 user 表的 is_follow 属性列失败")
	}
	/*
		token := username + password

		if user, exist := usersLoginInfo[token]; exist {
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 0},
				UserId:   user.Id,
				Token:    token,
			})
		} else {
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
			})
		}
	*/
}

func UserInfo(c *gin.Context) {
	token := c.Query("token")
	userId := c.Query("user_id")

	userIdInt, _ := strconv.ParseInt(userId, 10, 64)
	user, err := impl.UserServiceImpl{}.UserInfo(userIdInt, token)

	if err != nil {
		log.Printf(err.Error())
		c.JSON(http.StatusOK, models.UserResponse{
			Response: models.Response{StatusCode: 1, StatusMsg: err.Error()},
		})
	} else {
		c.JSON(http.StatusOK, models.UserResponse{
			Response: models.Response{StatusCode: 0},
			User:     *user,
		})
	}
	/*
		if user, exist := usersLoginInfo[token]; exist {
			c.JSON(http.StatusOK, UserResponse{
				Response: Response{StatusCode: 0},
				User:     user,
			})
		} else {
			c.JSON(http.StatusOK, UserResponse{
				Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
			})
		}
	*/
}
