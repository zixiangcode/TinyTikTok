package models

type FavoriteRelation struct { //储存在数据库中的评论结构
	CommonEntity       // 此中已包含了 CreateTime
	UserID       int64 `json:"user_id"`  //进行点赞操作的用户ID
	VideoID      int64 `json:"video_id"` //被点赞视频的ID
}

func (table *FavoriteRelation) TableName() string {
	return "favoriterelations"
}
