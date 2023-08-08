package controller

import (
	"net/http"
	"strconv"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

// 初始化GORM数据库连接
func initDB() (*gorm.DB, error) {
	dsn := "root:123456@tcp(192.168.111.129:3306)/TikTok?charset=utf8mb4&parseTime=True&loc=Local"
	//db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{TranslateError: true})
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}

// CommentActionRequest represents the JSON request for the comment action API
type CommentActionRequest struct {
	ActionType  string  `json:"action_type"`            // 1-发布评论，2-删除评论
	CommentID   *string `json:"comment_id,omitempty"`   // 要删除的评论id，在action_type=2的时候使用
	CommentText *string `json:"comment_text,omitempty"` // 用户填写的评论内容，在action_type=1的时候使用
	Token       string  `json:"token"`                  // 用户鉴权token
	VideoID     string  `json:"video_id"`               // 视频id
}

type MyComment struct {
	Id         uint64 `json:"id,omitempty"`
	User       User   `json:"user"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}

type CommentListResponse struct {
	Response
	CommentList []MyComment `json:"comment_list,omitempty"`
}

// CommentActionResponse represents the JSON response for the comment action API
type CommentActionResponse struct {
	Response
	Comment MyComment `json:"comment,omitempty"`
}

// CommentAction handles the comment action API
func CommentAction(c *gin.Context) {
	var request CommentActionRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, Response{
			StatusCode: 1,
			StatusMsg:  "Invalid request",
		})
		return
	}

	//单独开发评论API，暂时没用到token
	token := c.Query("token")
	actionType := c.Query("action_type")
	videoid := c.Query("video_id")

	//单独开发评论API，暂时没用到token
	//if user, exist := usersLoginInfo[token]; exist {
	if true {
		if actionType == "1" {
			if text := c.Query("comment_text"); text == "" {
				c.JSON(http.StatusBadRequest, Response{
					StatusCode: 1,
					StatusMsg:  "Comment text is empty",
				})
				return
			}

			//暂时将token用作userid
			userid, err := strconv.ParseUint(token, 10, 64)
			videoid, err := strconv.ParseUint(videoid, 10, 64)
			comment := Comment{
				VideoID:    videoid,
				UserID:     userid,
				Content:    *request.CommentText,
				CreateDate: time.Now().Format("01-02"), // Replace this with the actual creation date
			}

			// 初始化数据库连接
			db, err := initDB()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"status_code": 1,
					"status_msg":  "Internal Server Error",
				})
				return
			}

			// 创建评论记录
			if err := db.Create(&comment).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"status_code": 1,
					"status_msg":  "Failed to create comment",
				})
				return
			}

			//查询uesr信息，并拼接到response中
			user := User{}
			if err := db.Find(&user, userid).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"status_code": 1,
					"status_msg":  "Failed to create comment",
				})
				return
			}
			myComment := MyComment{
				Id:         comment.Id,
				User:       user,
				Content:    comment.Content,
				CreateDate: comment.CreateDate,
			}

			c.JSON(http.StatusOK, CommentActionResponse{
				Response: Response{
					StatusCode: 0,
					StatusMsg:  "Comment added successfully.",
				},
				Comment: myComment,
			})
			return

			//	删除评论操作
		} else if actionType == "2" {
			if request.CommentID == nil {
				c.JSON(http.StatusBadRequest, Response{
					StatusCode: 1,
					StatusMsg:  "Comment ID is required for action_type=2",
				})
				return
			}
			// 初始化数据库连接
			db, err := initDB()
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"status_code": 1,
					"status_msg":  "Internal Server Error",
				})
				return
			}

			// 删除评论记录
			if err := db.Where("id = ?", request.CommentID).Delete(&Comment{}).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"status_code": 1,
					"status_msg":  "Failed to delete comment",
				})
				return
			}

			c.JSON(http.StatusOK, Response{
				StatusCode: 0,
				StatusMsg:  "Comment deleted successfully.",
			})
			return
		} else {
			c.JSON(http.StatusBadRequest, Response{
				StatusCode: 1,
				StatusMsg:  "Invalid action_type.",
			})
			return
		}
	} else {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "User doesn't exist",
		})
	}
}

func CommentList(c *gin.Context) {
	//_:= c.Query("token")
	videoID := c.Query("video_id")

	db, err := initDB()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status_code": 1,
			"status_msg":  "Internal Server Error",
		})
		return
	}

	// 查询视频的所有评论，并按发布时间倒序排序
	var comments []Comment

	result := db.Preload("User").
		Find(&comments, "video_id = ?", videoID).
		Order("create_date desc")
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status_code": 1,
			"status_msg":  "Failed to get comments2",
		})
		return
	}

	var myComments []MyComment
	for _, comment := range comments {
		myComment := MyComment{
			Id:         comment.Id,
			User:       comment.User,
			Content:    comment.Content,
			CreateDate: comment.CreateDate,
		}
		myComments = append(myComments, myComment)
	}

	c.JSON(http.StatusOK, CommentListResponse{
		Response: Response{
			StatusCode: 0,
			StatusMsg:  "Success",
		},
		CommentList: myComments,
	})
}
