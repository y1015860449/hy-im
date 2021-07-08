package model

import "time"

type LoginInfo struct {
	Id         int64     `json:"id"`
	userId     int64     `json:"user_id"`
	Account    string    `json:"account"`
	loginToken string    `json:"login_token"`
	CreateTime time.Time `json:"create_time"`
}
