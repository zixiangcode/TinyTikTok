package models

import (
	"TinyTikTok/utils"
	"time"
)

// CommonEntity 通用实体，鉴于实体内容较少，不再拆分 service 层和 impl 层
type CommonEntity struct {
	Id         int64     `json:"id,omitempty"`
	CreateTime time.Time `json:"create_time,omitempty"`
	IsDeleted  int64     `json:"is_deleted"`
}

// NewCommonEntity 创建 CommonEntity 实体的实例对象
func NewCommonEntity() CommonEntity {
	sf := utils.NewSnowflake()
	return CommonEntity{
		Id:         sf.Generate(),
		CreateTime: time.Now(),
		IsDeleted:  0,
	}
}
