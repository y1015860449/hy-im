package user

import "hy-im/api/api-user/dao/db/model"

type ImUserDao interface {
	SaveUserInfo(info *model.Users) error
	GetUserInfoByAccount(account string) (*model.Users, error)
	GetUserStatus(userId int64) (int8, error)
	SetUserStatus(userId int64, status int8) error

	SaveLoginInfo(info *model.UserLogins) error
	GetLoginInfo(userId int64) (*model.UserLogins, error)

}