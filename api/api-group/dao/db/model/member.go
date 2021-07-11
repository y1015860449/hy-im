package model

import "time"

type GroupMemberInfo struct {
	Id         int64     `json:"id"`
	UserId     int64     `json:"user_id"`
	GroupId    int64     `json:"group_id"`
	GroupType  int8      `json:"group_type"`
	Status     int8      `json:"status"`
	Mute       int8      `json:"mute"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}
