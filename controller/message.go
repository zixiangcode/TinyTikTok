package controller

import (
	"TinyTikTok/models"
	"TinyTikTok/service/impl"
	"TinyTikTok/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

var tempChat = map[string][]models.Message{}

var messageIdSequence = int64(1)

type ChatResponse struct {
	models.Response
	MessageList []models.Message `json:"message_list"`
}

// MessageAction no practical effect, just check if token is valid
func MessageAction(c *gin.Context) {
	token := c.Query("token")
	toUserId := c.Query("to_user_id")
	content := c.Query("content")
	actionType := c.Query("action_type")

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

	//验证请求是否合法
	if toUserId == "" || content == "" || actionType != "1" {
		c.JSON(http.StatusBadRequest, models.Response{
			StatusCode: 1,
			StatusMsg:  "Invalid request",
		})
		return
	}

	err1 := impl.MessageServiceImpl{}.SendMessage(toUserId, userID, content)
	if err1 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status_code": 1,
			"status_msg":  "Failed to send message",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"status_msg":  "Message sent successfully!",
	})
	return
}

// MessageChat all users have same follow list
func MessageChat(c *gin.Context) {
	token := c.Query("token")
	toUserId := c.Query("to_user_id")

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

	//查询聊天消息
	MessageResponseList, err := impl.MessageServiceImpl{}.GetMessages(userID, toUserId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status_code": 1,
			"status_msg":  "Failed to loading messages",
		})
		return
	}

	c.JSON(http.StatusOK, models.MessageResponseList{
		StatusCode:          0,
		StatusMsg:           "Message loading successfully!",
		MessageResponseList: MessageResponseList,
	})

}
