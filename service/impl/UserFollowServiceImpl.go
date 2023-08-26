package impl

import (
	"TinyTikTok/dao"
	"TinyTikTok/db"
	"TinyTikTok/models"
	"gorm.io/gorm"
	"log"
)

type UserFollowServiceImpl struct {
}

func (commentServiceImpl UserFollowServiceImpl) AddUserFollow(follow models.UserFollow) (err error) {
	tx := db.GetMysqlDB().Begin()
	isCommit := false
	defer func() {
		if isCommit {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()

	userFollow, err := dao.GetUserFollowBy(tx, follow.UserID, follow.FollowUserID)
	if err != nil && err != gorm.ErrRecordNotFound {
		return
	}
	count := 1
	if err == gorm.ErrRecordNotFound {
		err = dao.AddUserFollow(tx, follow)
		if err != nil {
			log.Printf("查询user关注数据出错,user_id:%d, followUserID:%d", follow.UserID, follow.FollowUserID)
			return
		}
		goto UpdateFollowTotalCount
	}

	if userFollow.ActionType == 1 && follow.ActionType == 1 {
		isCommit = true
		return
	}

	// 取消关注后用户名的关注总量得减去1
	if userFollow.ActionType == 1 && follow.ActionType == 2 {
		count = -1
	}
	err = dao.UpdateFollow(tx, userFollow.Id, follow)
	if err != nil {
		log.Printf("更新用户关注记录出错,user_id:%d, followUserID:%d", follow.UserID, follow.FollowUserID)
		return
	}
UpdateFollowTotalCount:
	// 关注总数加一
	err = UserServiceImpl{}.UpdateFollowTotalCount(tx, follow.UserID, count)
	if err != nil {
		log.Printf("更新用户关注总量出错,user_id:%d, followUserID:%d", follow.UserID, follow.FollowUserID)
		return
	}

	// 被关注者粉丝总数加一
	err = UserServiceImpl{}.UpdateFollowerTotalCount(tx, follow.FollowUserID, count)
	if err != nil {
		log.Printf("更新用户关注总量出错,user_id:%d, followUserID:%d", follow.UserID, follow.FollowUserID)
		return
	}

	isCommit = true
	return
}

func (commentServiceImpl UserFollowServiceImpl) GetUserFollowByUserID(userID int64) (userFollowResps []models.UserFollowResponse, err error) {
	userFollows, err := dao.GetUserFollows(userID)
	if len(userFollows) == 0 {
		return
	}

	var followUserIDs []int64
	for _, followUser := range userFollows {
		followUserIDs = append(followUserIDs, followUser.FollowUserID)
	}

	users, err := UserServiceImpl{}.GetUserByIds(followUserIDs)
	if err != nil || len(users) == 0 {
		return
	}

	for _, user := range users {
		userFollowResp := models.UserFollowResponse{
			ID:              user.Id,
			Name:            user.Name,
			FollowCount:     user.FollowCount,
			FollowerCount:   user.FollowerCount,
			IsFollow:        user.IsFollow,
			Avatar:          user.Avatar,
			BackgroundImage: user.BackgroundImage,
			Signature:       user.Signature,
			TotalFavorite:   user.TotalFavorited,
			WorkCount:       user.WorkCount,
			FavoriteCount:   user.FavoriteCount,
		}
		userFollowResps = append(userFollowResps, userFollowResp)
	}
	return
}
