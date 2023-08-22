package models

import "time"

type Relation struct {
	OwnerId    int64     `json:"owner_id,omitempty"`
	TargetID   int64     `json:"target_id,omitempty"`
	CreateTime time.Time `json:"create_time,omitempty"`
}
