package friend

import "hy-im/api/api-friend/dao/db/model"

type ImFriendDao interface {
	GetFriendInfo(userId, friendId int64) (*model.FriendInfo, error)
	SaveFriendRelation(info *model.FriendInfo) error
	UpdateFriendRelationStatus(userId, friendId int64, status int8) error
}
