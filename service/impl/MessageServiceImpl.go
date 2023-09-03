package impl

import (
	"TinyTikTok/dao"
	"TinyTikTok/models"
	"log"
	"strconv"
	"time"
)

type MessageServiceImpl struct {
}

func (messageServiceImpl MessageServiceImpl) SendMessage(toUserid string, fromUserId int64, content string) error {

	//将toUserid从String转换成Int64
	toUserId, err := strconv.ParseInt(toUserid, 10, 64)
	if err != nil {
		log.Printf("toUserid从String转换成Int64失败 %v", err)
		return err
	}

	//准备储存到数据库中的数据
	var message = models.Message{
		CommonEntity: models.CommonEntity{
			CreateTime: time.Now(),
			IsDeleted:  0,
		},
		FromUserID: fromUserId,
		Content:    content,
		ToUserID:   toUserId,
	}

	//存入数据
	if err := dao.SendMessage(message); err != nil {
		log.Printf("方法 SendMessage 失败 %v", err)
		return err
	}
	return nil
}

func (messageServiceImpl MessageServiceImpl) GetMessages(userID int64, toUserid string) ([]models.MessageResponse, error) {

	//将toUserid从String转换成Int64
	toUserId, err := strconv.ParseInt(toUserid, 10, 64)
	if err != nil {
		log.Printf("toUserid从String转换成Int64失败 %v", err)
		return []models.MessageResponse{}, err
	}

	//查询消息数据
	messages, err := dao.GetMessages(userID, toUserId)
	if err != nil {
		log.Printf("方法 SendMessage 失败 %v", err)
		return []models.MessageResponse{}, err
	}

	//将查询到的message数据拼接到返回结果中
	var messageResponses = make([]models.MessageResponse, len(messages))

	for k, message := range messages {
		var messageResponse = models.MessageResponse{
			Content:    message.Content,
			CreateTime: message.CreateTime.Format("2006-01-02 15:04:05"),
			FromUserID: message.FromUserID,
			ID:         message.Id,
			ToUserID:   message.ToUserID,
		}
		messageResponses[k] = messageResponse
	}
	return messageResponses, nil
}
