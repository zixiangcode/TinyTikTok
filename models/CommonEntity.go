package models

import "time"

// CommonEntity 通用实体，鉴于实体内容较少，不再拆分 service 层和 impl 层
type CommonEntity struct {
	Id         int64     `json:"id,omitempty"`
	CreateTime time.Time `json:"create_time,omitempty"`
	IsDeleted  int64     `json:"is_deleted"`
}



