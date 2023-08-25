package controller

import (
	"TinyTikTok/models"
	"TinyTikTok/service/impl"
	"TinyTikTok/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserListResponse struct {
	models.Response
	UserList []models.User `json:"user_list"`
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	action_type := c.Query("action_type")
	to_user_Id, _ := strconv.Atoi(c.Query("to_user_id"))

	_, err := impl.UserServiceImpl{}.GetUserById(int64(to_user_Id))
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.Response{
			StatusCode: 1,
			StatusMsg:  "user doesn't exist",
		})
	}

	userClaims, err := utils.ParseToken(c.Query("token"))
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.Response{
			StatusCode: 1,
			StatusMsg:  "Unauthorized",
		})
	}

	user_id := userClaims.JWTCommonEntity.Id

	if action_type == "1" {
		err3 := impl.FollowServiceimpl{}.AddFollowAction(user_id, int64(to_user_Id))
		if err3 == nil {
			c.JSON(http.StatusOK, models.Response{StatusCode: 0})
		} else {
			c.JSON(http.StatusOK, models.Response{StatusCode: 1, StatusMsg: "add fail"})
		}
	} else if action_type == "2" {
		err3 := impl.FollowServiceimpl{}.DelFollowAction(user_id, int64(to_user_Id))
		if err3 == nil {
			c.JSON(http.StatusOK, models.Response{StatusCode: 0})
		} else {
			c.JSON(http.StatusOK, models.Response{StatusCode: 1, StatusMsg: "delate fail"})
		}
	}
}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	id := c.Query("user_id")
	userId, _ := strconv.Atoi(id)
	users, err := impl.FollowServiceimpl{}.GetFollowList(int64(userId))
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.Response{
			StatusCode: 1,
			StatusMsg:  "user doesn't exist",
		})
	} else {
		c.JSON(http.StatusOK, UserListResponse{
			Response: models.Response{
				StatusCode: 0,
			},
			UserList: *users,
		})
	}
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	id := c.Query("user_id")
	userId, _ := strconv.Atoi(id)
	users, err := impl.FollowServiceimpl{}.GetFollowerList(int64(userId))
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.Response{
			StatusCode: 1,
			StatusMsg:  "user doesn't exist",
		})
	} else {
		c.JSON(http.StatusOK, UserListResponse{
			Response: models.Response{
				StatusCode: 0,
			},
			UserList: *users,
		})
	}
}

// FriendList all users have same friend list
func FriendList(c *gin.Context) {
	id := c.Query("user_id")
	userId, _ := strconv.Atoi(id)
	users, err := impl.FollowServiceimpl{}.GetFriendList(int64(userId))
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.Response{
			StatusCode: 1,
			StatusMsg:  "user doesn't exist",
		})
	} else {
		c.JSON(http.StatusOK, UserListResponse{
			Response: models.Response{
				StatusCode: 0,
			},
			UserList: *users,
		})
	}
}
