package impl

import (
	"TinyTikTok/dao"
	"TinyTikTok/db"
	"TinyTikTok/models"
	"errors"
	"log"
)

type RelationServiceImpl struct {
}

// GetUserServiceImpl 实例化 UserServiceImpl
func GetUserServiceImpl() UserServiceImpl {
	var userService UserServiceImpl
	return userService
}

// GetRelationServiceImpl 实例化 RelationServiceImpl
func GetRelationServiceImpl() RelationServiceImpl {
	var relationServiceImpl RelationServiceImpl
	return relationServiceImpl
}

// FollowUser 关注用户
func (relationServiceImpl RelationServiceImpl) FollowUser(userId int64, toUserId int64, actionType int) error {
	tx := db.GetMysqlDB().Begin()
	isCommit := true // 默认提交事务
	defer func() {
		if isCommit {
			tx.Commit()
		} else {
			tx.Rollback()
		}
	}()

	if actionType == 1 {
		exists, err := dao.GetFollowByUserIdAndToUserId(userId, toUserId)
		count := 1 // 对用户关注总量的影响
		if err != nil {
			log.Printf("查询关注记录发生异常 = %v", err)
			return err
		}
		if exists.Id != 0 && exists.IsDeleted == 0 {
			log.Printf("该用户已关注")
			return errors.New("已关注")
		}
		// 不存在这个记录，进行关注操作，插入数据库
		err1 := dao.FollowUser(tx, userId, toUserId)
		if err1 != nil {
			log.Printf("插入关注关系失败：%v", err1)
			isCommit = false // 出现报错，取消提交事务
			return err1
		}
		// 关注总数增 1
		err2 := GetUserServiceImpl().UpdateFollowTotalCount(tx, userId, count)
		if err2 != nil {
			log.Printf("更新关注数失败：%v", err2)
			isCommit = false
			return err2
		}
		// 被关注者粉丝总数增 1
		err3 := GetUserServiceImpl().UpdateFollowerTotalCount(tx, toUserId, count)
		if err3 != nil {
			log.Printf("更新被关注者粉丝数失败：%v", err3)
			isCommit = false
			return err3
		}
	} else if actionType == 2 {
		exists, err := dao.GetFollowByUserIdAndToUserId(userId, toUserId)
		count := -1 // 取消关注对关注总量的影响
		if err != nil {
			log.Printf("查询关注记录发生异常 = %v", err)
			return err
		}
		if exists == (models.Follow{}) || exists.Id == 0 {
			log.Printf("未找到要取消关注的记录")
			return errors.New("未关注")
		}
		// 取消关注的数据库操作
		err1 := dao.UnFollowUser(tx, userId, toUserId)
		if err1 != nil {
			log.Printf("取消关注失败：%v", err1)
			isCommit = false // 出现报错，取消提交事务
			return err1
		}
		// 关注总数减 1
		err2 := GetUserServiceImpl().UpdateFollowTotalCount(tx, userId, count)
		if err2 != nil {
			log.Printf("更新关注数失败：%v", err2)
			isCommit = false
			return err2
		}
		// 被关注者粉丝总数减 1
		err3 := GetUserServiceImpl().UpdateFollowerTotalCount(tx, toUserId, count)
		if err3 != nil {
			log.Printf("更新被关注者粉丝数失败：%v", err3)
			isCommit = false
			return err3
		}
	}
	return nil
}

// GetFollows 获取自身关注列表
func (relationServiceImpl RelationServiceImpl) GetFollows(userId int64) ([]models.User, error) {
	var usersId []int64
	err := db.GetMysqlDB().Table("follow").
		Where("user_id = ? AND is_deleted = ?", userId, 0).
		Pluck("follow_user_id", &usersId).Error
	if err != nil {
		log.Printf("方法 GetFollows 失败: %v", err)
		return []models.User{}, err
	}
	users := make([]models.User, len(usersId))
	for i := range usersId {
		var user models.User
		err := db.GetMysqlDB().Table("user").Where("id = ?", usersId[i]).Find(&user).Error
		if err != nil {
			log.Printf("关注列表查找用户数据失败")
			return []models.User{}, err
		}
		users[i] = user
	}
	return users, nil
}

// GetFollowers 获取粉丝列表
func (relationServiceImpl RelationServiceImpl) GetFollowers(userId int64) ([]models.User, error) {
	var usersId []int64
	err := db.GetMysqlDB().Table("follow").
		Where("follow_user_id = ? AND is_deleted = ?", userId, 0).
		Pluck("user_id", &usersId).Error
	if err != nil {
		log.Printf("方法 GetFollowers 失败: %v", err)
		return []models.User{}, err
	}
	users := make([]models.User, len(usersId))
	for i := range usersId {
		var user models.User
		err := db.GetMysqlDB().Table("user").Where("id = ?", usersId[i]).Find(&user).Error
		if err != nil {
			log.Printf("粉丝列表查找用户数据失败")
			return []models.User{}, err
		}
		users[i] = user
	}
	return users, nil
}

// GetFriends 获取好友列表
func (relationServiceImpl RelationServiceImpl) GetFriends(userId int64) ([]models.User, error) {
	follows, err := GetRelationServiceImpl().GetFollows(userId)
	if err != nil {
		return nil, err
	}
	followers, err := GetRelationServiceImpl().GetFollowers(userId)
	if err != nil {
		return nil, err
	}

	friends := make([]models.User, 0)
	for _, user := range followers {
		if func(arr []models.User, id int64) bool {
			for _, u := range arr {
				if u.Id == id {
					return true
				}
			}
			return false
		}(follows, user.Id) {
			friends = append(friends, user)
		}
	}
	return friends, nil
}
