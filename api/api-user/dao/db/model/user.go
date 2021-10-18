package model

import "time"

type Users struct {
	Id           int64
	Status       int8
	HeadImgUrl   string
	NickName     string
	Gender       int8
	Birthday     string
	PerSignature string
	Account      string
	Email        string
	Mobile       string
	CountryCode  string
	RegisterType int8
	LoginPwd     string
	PwdSalt      string
	PublicKey    string
	CreateTime   time.Time
	UpdateTime   time.Time
}

type UserDevice struct {
	Id             int64
	UserId         int64
	DeviceToken    string
	PushType       int16
	PushToken      string
	VoipToken      string
	DeviceName     string
	IsCheck        int8
	IsCurrentUsed  int8
	LastLoginToken string
	LastLoginTime  int64
	AppVersion     int64
	CreateTime     time.Time
	UpdateTime     time.Time
}
