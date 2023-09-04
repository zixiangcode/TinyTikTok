package controller

import (
	"TinyTikTok/models"
	"TinyTikTok/service/impl"
	"TinyTikTok/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	token := c.Query("token")
	actionTypeStr := c.Query("action_type")
	videoIdStr := c.Query("video_id")

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
	if actionTypeStr != "1" && actionTypeStr != "2" || videoIdStr == "" {
		c.JSON(http.StatusBadRequest, models.Response{
			StatusCode: 1,
			StatusMsg:  "Invalid request",
		})
		return
	}

	//将videoId从String转换成Int64
	videoID, err := strconv.ParseInt(videoIdStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			StatusCode: 1,
			StatusMsg:  "Invalid videoID",
		})
		return
	}

	//actionTypeStr为1，进行点赞操作
	if actionTypeStr == "1" {

		//准备点赞操作数据
		favoriteRelation := models.FavoriteRelation{
			CommonEntity: models.NewCommonEntity(),
			UserID:       userID,
			VideoID:      videoID,
		}

		//调用AddfavoriteRelation函数添加点赞信息
		err := impl.FavoriteRelationServiceImpl{}.AddfavoriteRelation(favoriteRelation)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status_code": 1,
				"status_msg":  "Failed to like",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"StatusCode": 0,
			"StatusMsg":  "Like Succeeded!",
		})
		return
	} else if actionTypeStr == "2" { //取消点赞操作

		//准备取消点赞操作数据
		favoriteRelation := models.FavoriteRelation{
			UserID:  userID,
			VideoID: videoID,
		}

		// 删除评论记录
		err := impl.FavoriteRelationServiceImpl{}.DeletefavoriteRelation(favoriteRelation)
		if err != nil {
			log.Printf("%v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"status_code": 1,
				"status_msg":  "Failed to Unlike",
			})
			return
		}
		c.JSON(http.StatusOK, models.CommentListResponse{
			StatusCode:  0,
			StatusMsg:   "Successfully Unliked",
			CommentList: nil,
		})
		return
	}
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	c.JSON(http.StatusOK, VideoListResponse{
		Response: models.Response{
			StatusCode: 0,
		},
		VideoList: DemoVideos,
	})
}
