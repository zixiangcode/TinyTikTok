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

func (messageServiceImpl MessageServiceImpl) SendMessage(toUserId int64, fromUserId int64, content string) error {

	//todo 应在存入数据库之前验证发送者与接收者是否为朋友关系

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

func (messageServiceImpl MessageServiceImpl) GetMessages(userID int64, toUserid string, preMsgTimeStr string) ([]models.MessageResponse, error) {

	//将toUserid从String转换成Int64
	toUserId, err := strconv.ParseInt(toUserid, 10, 64)
	if err != nil {
		log.Printf("toUserid从String转换成Int64失败 %v", err)
		return []models.MessageResponse{}, err
	}

	//将preMsgTime从String转换成Int64时间戳
	preMsgTimeInt, err1 := strconv.ParseInt(preMsgTimeStr, 10, 64)
	if err1 != nil {
		log.Printf("toUserid从String转换成Int64失败 %v", err)
		return []models.MessageResponse{}, err
	}

	// 创建一个时间对象
	t := time.Unix(preMsgTimeInt, 0)
	//将时间戳数据转换成yyyy-mm-dd hh:ss:mm格式
	preMsgTime := t.Format("2006-01-02 15:04:05")

	//查询消息数据
	messages, err := dao.GetMessages(userID, toUserId, preMsgTime)
	if err != nil {
		log.Printf("方法 SendMessage 失败 %v", err)
		return []models.MessageResponse{}, err
	}

	//将查询到的message数据拼接到返回结果中
	var messageResponses = make([]models.MessageResponse, len(messages))

	for k, message := range messages {

		var messageResponse = models.MessageResponse{
			Content:    message.Content,
			CreateTime: message.CreateTime.Unix(),
			FromUserID: message.FromUserID,
			ID:         message.Id,
			ToUserID:   message.ToUserID,
		}
		messageResponses[k] = messageResponse
	}
	return messageResponses, nil
}
