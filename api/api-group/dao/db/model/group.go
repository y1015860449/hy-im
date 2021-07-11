package model

import "time"

type GroupInfo struct {
	Id         int64     `json:"id"`
	Status     int8      `json:"status"`
	Type       int8      `json:"type"`
	Mute       int8      `json:"mute"`
	CreatorId  int64     `json:"creator_id"`
	MasterId   int64     `json:"master_id"`
	Total      int32     `json:"total"`
	SecretKey  string    `json:"secret_key"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}
