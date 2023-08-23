package controller

import (
	"TinyTikTok/models"
	"TinyTikTok/service/impl"
	"TinyTikTok/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type UserListResponse struct {
	models.Response
	UserList []models.User `json:"user_list"`
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	token := c.Query("token")
	actionTypeStr := c.Query("action_type")
	followUserIDStr := c.Query("to_user_id")
	//验证token
	userClaims, err := utils.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.Response{
			StatusCode: 1,
			StatusMsg:  "Unauthorized",
		})
		return
	}
	userID := userClaims.JWTCommonEntity.Id

	//验证请求是否错误
	if actionTypeStr != "1" && actionTypeStr != "2" || followUserIDStr == "" {
		c.JSON(http.StatusBadRequest, models.Response{
			StatusCode: 1,
			StatusMsg:  "Invalid request",
		})
		return
	}

	//将videoId从String转换成Int64
	followUserID, err := strconv.ParseInt(followUserIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			StatusCode: 1,
			StatusMsg:  "Invalid followID",
		})
		return
	}

	actionType, err := strconv.Atoi(actionTypeStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			StatusCode: 1,
			StatusMsg:  "Invalid actionType",
		})
		return
	}

	follow := models.UserFollow{
		UserID:       userID,
		FollowUserID: followUserID,
		ActionType:   actionType,
		CommonEntity: models.CommonEntity{CreateTime: time.Now()},
		UpdateTime:   time.Now(),
	}

	err = impl.UserFollowServiceImpl{}.AddUserFollow(follow)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			StatusCode: 1,
			StatusMsg:  "operator error",
		})
		return
	}
	c.JSON(http.StatusOK, models.Response{
		StatusCode: 200,
		StatusMsg:  "ok",
	})
	return
}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	token := c.Query("token")
	userIDStr := c.Query("user_id")
	//验证token
	_, err := utils.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.Response{
			StatusCode: 1,
			StatusMsg:  "Unauthorized",
		})
		return
	}

	//将videoId从String转换成Int64
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			StatusCode: 1,
			StatusMsg:  "Invalid userID",
		})
		return
	}

	userFollowResps, err := impl.UserFollowServiceImpl{}.GetUserFollowByUserID(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			StatusCode: 1,
			StatusMsg:  "request userID",
		})
		return
	}

	c.JSON(http.StatusOK, models.UserFollowListResponse{
		StatusCode:     200,
		StatusMsg:      "ok",
		UserFollowResp: userFollowResps,
	})
	return
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: models.Response{
			StatusCode: 0,
		},
		UserList: []models.User{DemoUser},
	})
}

// FriendList all users have same friend list
func FriendList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: models.Response{
			StatusCode: 0,
		},
		UserList: []models.User{DemoUser},
	})
}
