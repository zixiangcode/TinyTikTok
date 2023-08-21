package controller

import (
	"TinyTikTok/dao"
	"TinyTikTok/models"
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
	token := c.Query("token")
	user2Id, _ := strconv.Atoi(c.Query("to_user_id"))
	user1, err := dao.FindUserByToken(token)
	if err != nil {
		c.JSON(http.StatusOK, models.Response{StatusCode: 1, StatusMsg: "User1 doesn't exist"})
		return
	}
	user2, err2 := dao.FindUserById(int64(user2Id))
	if err2 != nil {
		c.JSON(http.StatusOK, models.Response{StatusCode: 1, StatusMsg: "User2 doesn't exist"})
		return
	}
	if err3 := dao.AddFollow(user1, user2); err3 == nil {
		c.JSON(http.StatusOK, models.Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, models.Response{StatusCode: 1, StatusMsg: "add fail"})
	}
}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	id := c.Query("user_id")
	userId, _ := strconv.Atoi(id)
	if _, err := dao.FindUserById(int64(userId)); err == nil {
		users := dao.GetFollow(int64(userId))
		c.JSON(http.StatusOK, UserListResponse{
			Response: models.Response{
				StatusCode: 0,
			},
			UserList: *users,
		})
	} else {
		c.JSON(http.StatusOK, UserListResponse{
			Response: models.Response{
				StatusCode: 1,
				StatusMsg:  "no user",
			},
		})
	}
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	id := c.Query("user_id")
	userId, _ := strconv.Atoi(id)
	if _, err := dao.FindUserById(int64(userId)); err == nil {
		users := dao.GetFollower(int64(userId))
		c.JSON(http.StatusOK, UserListResponse{
			Response: models.Response{
				StatusCode: 0,
			},
			UserList: *users,
		})
	} else {
		c.JSON(http.StatusOK, UserListResponse{
			Response: models.Response{
				StatusCode: 1,
				StatusMsg:  "no user",
			},
		})
	}
}

// FriendList all users have same friend list
func FriendList(c *gin.Context) {
	id := c.Query("user_id")
	userId, _ := strconv.Atoi(id)
	if _, err := dao.FindUserById(int64(userId)); err == nil {
		users := dao.GetFriend(int64(userId))
		c.JSON(http.StatusOK, UserListResponse{
			Response: models.Response{
				StatusCode: 0,
			},
			UserList: *users,
		})
	} else {
		c.JSON(http.StatusOK, UserListResponse{
			Response: models.Response{
				StatusCode: 1,
				StatusMsg:  "no user",
			},
		})
	}
}
