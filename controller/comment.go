package controller

//import (
//	"net/http"
//
//	"github.com/gin-gonic/gin"
//)
//
//type CommentListResponse struct {
//	Response
//	CommentList []Comment `json:"comment_list,omitempty"`
//}
//
//type CommentActionResponse struct {
//	Response
//	Comment Comment `json:"comment,omitempty"`
//}
//
//// CommentAction no practical effect, just check if token is valid
//func CommentAction(c *gin.Context) {
//	token := c.Query("token")
//	actionType := c.Query("action_type")
//
//	if user, exist := usersLoginInfo[token]; exist {
//		if actionType == "1" {
//			text := c.Query("comment_text")
//			c.JSON(http.StatusOK, CommentActionResponse{Response: Response{StatusCode: 0},
//				Comment: Comment{
//					Id:         1,
//					User:       user,
//					Content:    text,
//					CreateDate: "05-01",
//				}})
//			return
//		}
//		c.JSON(http.StatusOK, Response{StatusCode: 0})
//	} else {
//		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
//	}
//}
//
//// CommentList all videos have same demo comment list
//func CommentList(c *gin.Context) {
//	c.JSON(http.StatusOK, CommentListResponse{
//		Response:    Response{StatusCode: 0},
//		CommentList: DemoComments,
//	})
//}

import (
	"net/http"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

// 初始化GORM数据库连接
func initDB() (*gorm.DB, error) {
	dsn := "root:123456@tcp(192.168.111.129:3306)/TikTok?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// 自动迁移数据库，创建Comment表
	err = db.AutoMigrate(&Comment{})
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

type videoComment struct {
	VideoId   string `json:"video_id"`
	CommentId int64  `json:"comment_id"`
}

// CommentListResponse represents the JSON response for the comment list API
type CommentListResponse struct {
	Response
	CommentList []Comment `json:"comment_list,omitempty"`
}

// CommentActionResponse represents the JSON response for the comment action API
type CommentActionResponse struct {
	Response
	Comment Comment `json:"comment,omitempty"`
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

	token := c.Query("token")
	actionType := c.Query("action_type")

	if user, exist := usersLoginInfo[token]; exist {
		if actionType == "1" {
			if text := c.Query("comment_text"); text == "" {
				c.JSON(http.StatusBadRequest, Response{
					StatusCode: 1,
					StatusMsg:  "Comment text is empty",
				})
				return
			}

			comment := Comment{
				User:       user,
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
			defer db.Close()

			// 开启事务
			tx := db.Begin()
			if tx.Error != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"status_code": 1,
					"status_msg":  "Failed to create comment",
				})
				return
			}

			// 创建评论记录
			if err := tx.Create(&comment).Error; err != nil {
				// 回滚事务
				tx.Rollback()
				c.JSON(http.StatusInternalServerError, gin.H{
					"status_code": 1,
					"status_msg":  "Failed to create comment",
				})
				return
			}
			commentID := comment.Id
			if err := tx.Create(videoComment{CommentId: commentID}).Error; err != nil {
				// 回滚事务
				tx.Rollback()
				c.JSON(http.StatusInternalServerError, gin.H{
					"status_code": 1,
					"status_msg":  "Failed to video comment",
				})
				return
			}

			// 提交事务
			tx.Commit()

			c.JSON(http.StatusOK, CommentActionResponse{
				Response: Response{
					StatusCode: 0,
					StatusMsg:  "Comment added successfully.",
				},
				Comment: comment,
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
			defer db.Close()

			// 删除评论记录
			if err := db.Where("id = ?", ID).Delete(&Comment{}).Error; err != nil {
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
	token := c.Query("token")
	videoID := c.Query("video_id")

	// 查询视频的所有评论，并按发布时间倒序排序
	var comments []Comment
	if err := db.Where("video_id = ?", videoID).Order("create_date desc").Find(&comments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status_code": 1,
			"status_msg":  "Failed to get comments",
		})
		return
	}

	comments := []Comment{
		{
			ID: 1,
			User: User{
				ID:              1,
				Name:            "User1",
				FollowCount:     10,
				FollowerCount:   20,
				IsFollow:        true,
				Avatar:          "avatar_url",
				BackgroundImage: "background_url",
				Signature:       "User1 signature",
				TotalFavorited:  "100",
				WorkCount:       5,
				FavoriteCount:   50,
			},
			Content:    "This is the first comment.",
			CreateDate: "2023-08-05 12:34:56",
		},
		{
			ID: 2,
			User: User{
				ID:              2,
				Name:            "User2",
				FollowCount:     15,
				FollowerCount:   25,
				IsFollow:        false,
				Avatar:          "avatar_url",
				BackgroundImage: "background_url",
				Signature:       "User2 signature",
				TotalFavorited:  "200",
				WorkCount:       10,
				FavoriteCount:   100,
			},
			Content:    "This is the second comment.",
			CreateDate: "2023-08-05 13:45:32",
		},
	}

	c.JSON(http.StatusOK, CommentListResponse{
		Response: Response{
			StatusCode: 0,
			StatusMsg:  "Success",
		},
		CommentList: comments,
	})
}
