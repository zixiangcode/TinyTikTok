package impl

import (
	"TinyTikTok/dao"
	"TinyTikTok/models"
	"log"
)

type CommentServiceImpl struct {
}

func (commentServiceImpl CommentServiceImpl) AddComment(comment models.Comment) (models.CommentCommonResponse, error) {

	commentId, err := dao.AddComment(comment)
	if err != nil {
		log.Printf("方法 AddComment() 失败 %v", err)
		return models.CommentCommonResponse{}, err
	}

	//查询uesr信息，并拼接到response中
	user, err := UserServiceImpl{}.GetUserById(comment.UserID)
	if err != nil {
		log.Printf("方法 AddComment() 失败 %v", err)
		return models.CommentCommonResponse{}, err
	}

	commentuserinfo := models.CommentUserInfo{
		ID:              user.Id,
		Name:            user.Name,
		FollowCount:     user.FollowCount,
		FollowerCount:   user.FollowerCount,
		IsFollow:        user.IsFollow,
		Avatar:          user.Avatar,
		BackgroundImage: user.BackgroundImage,
		Signature:       user.Signature,
		TotalFavorited:  user.TotalFavorited,
		WorkCount:       user.WorkCount,
		FavoriteCount:   user.FavoriteCount,
	}

	commentCommonResponse := models.CommentCommonResponse{
		Id:         commentId,
		User:       commentuserinfo,
		Content:    comment.Content,
		CreateTime: comment.CreateTime.Format("01-02"),
	}
	return commentCommonResponse, nil
}

func (commentServiceImpl CommentServiceImpl) DeleteComment(commentID int64, userID int64) error {

	err := dao.DeleteComment(commentID, userID)
	if err != nil {
		log.Printf("方法 DeleteComment() 失败 %v", err)
		return err
	}
	return nil
}

func (commentServiceImpl CommentServiceImpl) GetCommentsByVideoID(videoID int64) ([]models.CommentCommonResponse, error) {

	//获取评论信息，先拿到userID
	comments, err := dao.GetCommentsByVideoID(videoID)
	if err != nil {
		log.Printf("方法 GetCommentsByVideoID 失败 %v", err)
		return []models.CommentCommonResponse{}, err
	}

	var commentCommonResponses []models.CommentCommonResponse
	for _, comment := range comments {
		//从comments中读取userID查询uesr信息，并拼接到response中
		user, err := UserServiceImpl{}.GetUserById(comment.UserID)
		if err != nil {
			log.Printf("方法 GetCommentsByVideoID() 失败 %v", err)
			return []models.CommentCommonResponse{}, err
		}

		commentuserinfo := models.CommentUserInfo{
			ID:              user.Id,
			Name:            user.Name,
			FollowCount:     user.FollowCount,
			FollowerCount:   user.FollowerCount,
			IsFollow:        user.IsFollow,
			Avatar:          user.Avatar,
			BackgroundImage: user.BackgroundImage,
			Signature:       user.Signature,
			TotalFavorited:  user.TotalFavorited,
			WorkCount:       user.WorkCount,
			FavoriteCount:   user.FavoriteCount,
		}

		myComment := models.CommentCommonResponse{
			Id:         comment.Id,
			User:       commentuserinfo,
			Content:    comment.Content,
			CreateTime: comment.CreateTime.Format("01-02"),
		}
		commentCommonResponses = append(commentCommonResponses, myComment)
	}
	return commentCommonResponses, nil
}
