package model

import "time"

type UserInfo struct {
	Id         int64     `json:"id"`
	Status     int8      `json:"status"`
	Account    string    `json:"account"`
	SecretKey  string    `json:"secret_key"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}
