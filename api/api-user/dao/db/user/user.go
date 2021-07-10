package user

import "hy-im/api/api-user/dao/db/model"

type ImUserDao interface {
	SaveUserInfo(info *model.UserInfo) error
	GetUserInfoByAccount(account string) (*model.UserInfo, error)
	GetUserStatus(userId int64) (int8, error)
	SetUserStatus(userId int64, status int8) error

	SaveLoginInfo(info *model.LoginInfo) error
	GetLoginInfo(userId int64) (*model.LoginInfo, error)
}