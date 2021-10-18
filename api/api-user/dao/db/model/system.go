package model

import "time"

type AppVersion struct {
	Id          uint32
	OsType      int8
	AppVersion  string
	DownloadUrl string
	UpdateMsg   string
	Description string
	ForceUpdate int8
	VersionCode uint32
	Created     time.Time
	Updated     time.Time
}

type SystemApp struct {
	Id        uint32
	AppKey    int32
	AppSecret string
	AppName   string
	AppDesc   string
	IsEnabled int8
	Created   time.Time
	Updated   time.Time
}
