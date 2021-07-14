package model

import "time"

type FriendInfo struct {
	Id         int64     `json:"id"`
	UserId     int64     `json:"user_id"`
	FriendId   int64     `json:"friend_id"`
	Status     int8      `json:"status"`
	SecretKey  string    `json:"secret_key"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}
