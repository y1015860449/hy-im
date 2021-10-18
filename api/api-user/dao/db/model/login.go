package model

import "time"

type UserLogins struct {
	UserId      int64
	Account     string
	Email       string
	Mobile      string
	CountryCode string
	LoginType   int8
	IsVerified  int8
	IsDefault   int8
	CreateTime  time.Time
}

type UserLoginLogs struct {
	Id          int64
	UserId      int64
	UserToken   string
	LoginName   string
	LoginType   int8
	Host        string
	DeviceToken string
	IsSuccess   int8
	IsCheck     uint8
	AppVersion  string
	MbOsVersion string
	CreateTime  time.Time
	UpdateTime  time.Time
}
