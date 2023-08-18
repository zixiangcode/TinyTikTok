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

func CommentAction(c *gin.Context) { //添加或者删除评论操作

	token := c.Query("token")
	actionType := c.Query("action_type")
	videoId := c.Query("video_id")
	commentText := c.Query("comment_text")
	commentId := c.Query("comment_id")

	//验证请求是否错误
	if actionType != "1" && actionType != "2" || videoId == "" {
		c.JSON(http.StatusBadRequest, models.Response{
			StatusCode: 1,
			StatusMsg:  "Invalid request",
		})
		return
	}

	//将videoId从String转换成Int64
	videoID, err := strconv.ParseInt(videoId, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			StatusCode: 1,
			StatusMsg:  "Invalid videoid",
		})
		return
	}

	//验证token
	userClaims, err := utils.ParseToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, models.Response{
			StatusCode: 1,
			StatusMsg:  "Unauthorized",
		})
	}
	// 从 userClaims 中获取 UserID
	userID := userClaims.JWTCommonEntity.Id

	//actionType为1，进行发表评论操作
	if actionType == "1" {
		//判断评论内容是否为空
		if commentText == "" {
			c.JSON(http.StatusBadRequest, models.Response{
				StatusCode: 1,
				StatusMsg:  "Comment text is empty",
			})
			return
		}

		//准备评论数据
		comment := models.Comment{
			VideoID:      videoID,
			UserID:       userID,
			Content:      commentText,
			CommonEntity: models.CommonEntity{CreateTime: time.Now()},
		}

		//调用AddComment函数添加评论信息
		commentCommonResponse, err := impl.CommentServiceImpl{}.AddComment(comment)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status_code": 1,
				"status_msg":  "Failed to create comment",
			})
			return
		}
		c.JSON(http.StatusOK, models.CommentActionResponse{
			StatusCode: 0,
			StatusMsg:  "Comment added successfully.",
			Comment:    commentCommonResponse,
		})
		return
	} else if actionType == "2" { //删除评论操作
		//actionType为2，删除评论操作需要comment_id
		if commentId == "" {
			c.JSON(http.StatusBadRequest, models.Response{
				StatusCode: 1,
				StatusMsg:  "Comment ID is required for action_type=2",
			})
			return
		}

		//将comment_id从String转换成Int64
		commentID, err := strconv.ParseInt(commentId, 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.Response{
				StatusCode: 1,
				StatusMsg:  "Invalid videoId",
			})
			return
		}
		// 删除评论记录
		err1 := impl.CommentServiceImpl{}.DeleteComment(commentID)
		if err1 != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status_code": 1,
				"status_msg":  "Failed to delete comment",
			})
			return
		}

		c.JSON(http.StatusOK, models.CommentListResponse{
			StatusCode:  0,
			StatusMsg:   "Comment deletion successful",
			CommentList: nil,
		})
		return
	}
}

func CommentList(c *gin.Context) { // 查询视频评论列表操作
	//将videoID从String转换成Int64
	videoID, err := strconv.ParseInt(c.Query("video_id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			StatusCode: 1,
			StatusMsg:  "Invalid videoId",
		})
		return
	}

	myComments, err := impl.CommentServiceImpl{}.GetCommentsByVideoID(videoID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status_code": 1,
			"status_msg":  "Failed to get comments",
		})
		return
	}
	c.JSON(http.StatusOK, models.CommentListResponse{
		StatusCode:  0,
		StatusMsg:   "Success",
		CommentList: myComments,
	})
}
