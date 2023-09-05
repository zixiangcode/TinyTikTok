package controller

import (
	"TinyTikTok/dao"
	"TinyTikTok/models"
	"TinyTikTok/service/impl"
	"TinyTikTok/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type UserListResponse struct {
	models.Response
	UserList []models.User `json:"user_list"`
}

// GetRelationServiceImpl 实例化 RelationService
func GetRelationServiceImpl() impl.RelationServiceImpl {
	var relationService impl.RelationServiceImpl
	return relationService
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	token := c.Query("token")
	actionTypeStr := c.Query("action_type")
	followUserIDStr := c.Query("to_user_id")
	// 验证 token
	userClaims, err := utils.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.Response{
			StatusCode: 1,
			StatusMsg:  "Unauthorized",
		})
		return
	}
	userID := userClaims.JWTCommonEntity.Id

	// 验证请求是否错误
	if actionTypeStr != "1" && actionTypeStr != "2" || followUserIDStr == "" {
		c.JSON(http.StatusBadRequest, models.Response{
			StatusCode: 1,
			StatusMsg:  "Invalid request",
		})
		return
	}

	// 将 Id 从 String 转换成 Int64
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

	if userID == followUserID {
		c.JSON(http.StatusBadRequest, models.Response{
			StatusCode: 1,
			StatusMsg:  "不能关注/取关自己",
		})
		return
	}

	errFollow := GetRelationServiceImpl().FollowUser(userID, followUserID, actionType)
	if errFollow != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			StatusCode: 1,
			StatusMsg:  "关注/取关用户失败",
		})
		return
	}
	err = dao.UpdateUserFollowByUserId(userID)
	if err != nil {
		log.Printf("更新 user 表的 is_follow 属性列失败")
	}

	c.JSON(http.StatusOK, models.Response{
		StatusCode: 0,
		StatusMsg:  "关注成功",
	})
	return
}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	token := c.Query("token")
	userIDStr := c.Query("user_id")
	// 验证 token
	_, err := utils.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.Response{
			StatusCode: 1,
			StatusMsg:  "Unauthorized",
		})
		return
	}

	// 将 videoId 从 String 转换成 Int64
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			StatusCode: 1,
			StatusMsg:  "Invalid userID",
		})
		return
	}

	userFollows, err := GetRelationServiceImpl().GetFollows(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			StatusCode: 1,
			StatusMsg:  "request userID",
		})
		return
	}
	c.JSON(http.StatusOK, models.FollowListResponse{
		StatusCode:         0,
		StatusMsg:          "获取关注列表成功",
		UserFollowResponse: userFollows,
	})
	return
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	token := c.Query("token")
	userIDStr := c.Query("user_id")
	// 验证 token
	_, err := utils.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.Response{
			StatusCode: 1,
			StatusMsg:  "Unauthorized",
		})
		return
	}

	// 将 videoId 从 String 转换成 Int64
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			StatusCode: 1,
			StatusMsg:  "Invalid userID",
		})
		return
	}

	followers, err := GetRelationServiceImpl().GetFollowers(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			StatusCode: 1,
			StatusMsg:  "request userID",
		})
		return
	}
	c.JSON(http.StatusOK, models.FollowListResponse{
		StatusCode:         0,
		StatusMsg:          "获取粉丝列表成功",
		UserFollowResponse: followers,
	})
	return
}

// FriendList all users have same friend list
func FriendList(c *gin.Context) {
	token := c.Query("token")
	userIDStr := c.Query("user_id")
	// 验证 token
	_, err := utils.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.Response{
			StatusCode: 1,
			StatusMsg:  "Unauthorized",
		})
		return
	}

	// 将 videoId 从 String 转换成 Int64
	userID, err := strconv.ParseInt(userIDStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			StatusCode: 1,
			StatusMsg:  "Invalid userID",
		})
		return
	}

	friends, err := GetRelationServiceImpl().GetFriends(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.Response{
			StatusCode: 1,
			StatusMsg:  "request userID",
		})
		return
	}
	c.JSON(http.StatusOK, models.FollowListResponse{
		StatusCode:         0,
		StatusMsg:          "获取朋友列表成功",
		UserFollowResponse: friends,
	})
	return
}
