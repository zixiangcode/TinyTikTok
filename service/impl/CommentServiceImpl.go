package impl

import (
	"TinyTikTok/db"
	"TinyTikTok/models"
	"log"
)

type CommentServiceImpl struct {
}

func (commentServiceImpl CommentServiceImpl) Add(comment models.Comment) (int64, error) {
	err := db.GetMysqlDB().Create(&comment).Error
	if err != nil {
		log.Printf("方法 Create() 失败 %v", err)
		return 1, err
	}
	return comment.Id, nil
}

func (commentServiceImpl CommentServiceImpl) Delete(commentID int64) error {
	err := db.GetMysqlDB().Model(&models.Comment{}).
		Where("id = ?", commentID).
		Update("is_deleted", 1).Error
	if err != nil {
		log.Printf("方法 delete() 失败 %v", err)
		return err
	}
	return err
}

func (commentServiceImpl CommentServiceImpl) GetCommentsByVideoID(videoID int64) ([]models.MyComment, error) {
	var comments []models.Comment
	result := db.GetMysqlDB().Preload("User").
		Find(&comments, "video_id = ?", videoID).
		Order("create_date desc")
	if result.Error != nil {
		log.Printf("方法 GetCommentsByVideoID 失败 %v", result.Error)
		return nil, result.Error
	}
	var myComments []models.MyComment
	for _, comment := range comments {
		var user models.User
		results := db.GetMysqlDB().Find(&user).Where("id = ?", comment.UserID)
		if results.Error != nil {
			log.Printf("方法 GetCommentsByVideoID 失败 %v", result.Error)
			return nil, result.Error
		}
		myComment := models.MyComment{
			Id:         comment.Id,
			User:       user,
			Content:    comment.Content,
			CreateDate: comment.CreateTime.Format("01-02"),
		}
		myComments = append(myComments, myComment)
	}
	return myComments, result.Error
}
