package impl

import (
	"TinyTikTok/dao"
	"TinyTikTok/models"
	"errors"
	"log"
)

type FollowServiceimpl struct {
}

func (followServiceimpl FollowServiceimpl) AddFollowAction(userID int64, to_userID int64) error {
	if userID == to_userID {
		log.Printf("相同用户")
		return errors.New("the same user")
	}
	return dao.AddFollow(userID, to_userID)
}

func (followServiceimpl FollowServiceimpl) DelFollowAction(userID int64, to_userID int64) error {
	if userID == to_userID {
		log.Printf("相同用户")
		return errors.New("the same user")
	}
	return dao.DelFollow(userID, to_userID)
}

func (followServiceimpl FollowServiceimpl) GetFollowList(Id int64) (*[]models.User, error) {
	_, err := UserServiceImpl{}.GetUserById(Id)
	if err != nil {
		log.Printf("方法getUserById 失败")
		return nil, errors.New("no user")
	}
	return dao.GetFollow(Id), nil
}

func (followServiceimpl FollowServiceimpl) GetFollowerList(Id int64) (*[]models.User, error) {
	_, err := UserServiceImpl{}.GetUserById(Id)
	if err != nil {
		log.Printf("方法getUserById 失败")
		return nil, errors.New("no user")
	}
	return dao.GetFollower(Id), nil
}

func (followServiceimpl FollowServiceimpl) GetFriendList(Id int64) (*[]models.User, error) {
	_, err := UserServiceImpl{}.GetUserById(Id)
	if err != nil {
		log.Printf("方法getUserById 失败")
		return nil, errors.New("no user")
	}
	return dao.GetFriend(Id), nil
}
