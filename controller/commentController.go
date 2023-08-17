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
	videoid := c.Query("video_id")

	//验证请求是否错误
	if actionType != "1" && actionType != "2" || videoid == "" {
		c.JSON(http.StatusBadRequest, models.Response{
			StatusCode: 1,
			StatusMsg:  "Invalid request",
		})
		return
	}

	//将videoid从String转换成Int64
	videoID, err := strconv.ParseInt(videoid, 10, 64)
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
	userID := userClaims.CommonEntity.Id

	//actionType为1，进行发表评论操作
	if actionType == "1" {
		//判断评论内容是否为空
		if c.Query("comment_text") == "" {
			c.JSON(http.StatusBadRequest, models.Response{
				StatusCode: 1,
				StatusMsg:  "Comment text is empty",
			})
			return
		}

		//准备评论数据
		commonEntity := models.CommonEntity{
			CreateTime: time.Now(),
		}
		comment := models.Comment{
			VideoID:      videoID,
			UserID:       userID,
			Content:      c.Query("comment_text"),
			CommonEntity: commonEntity,
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
			Response: models.Response{
				StatusCode: 0,
				StatusMsg:  "Comment added successfully.",
			},
			Comment: commentCommonResponse,
		})
		return
	} else if actionType == "2" { //删除评论操作
		//actionType为2，删除评论操作需要comment_id
		if c.Query("comment_id") == "" {
			c.JSON(http.StatusBadRequest, models.Response{
				StatusCode: 1,
				StatusMsg:  "Comment ID is required for action_type=2",
			})
			return
		}

		//将comment_id从String转换成Int64
		commentID, err := strconv.ParseInt(c.Query("comment_id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, models.Response{
				StatusCode: 1,
				StatusMsg:  "Invalid videoid",
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

		c.JSON(http.StatusOK, models.Response{
			StatusCode: 0,
			StatusMsg:  "Comment deleted successfully.",
		})
		return
	}
}

func CommentList(c *gin.Context) { // 查询视频评论列表操作
	videoid := c.Query("video_id")
	//将videoid从String转换成Int64
	videoID, err := strconv.ParseInt(videoid, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, models.Response{
			StatusCode: 1,
			StatusMsg:  "Invalid videoid",
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
		Response: models.Response{
			StatusCode: 0,
			StatusMsg:  "Success",
		},
		CommentList: myComments,
	})
}
