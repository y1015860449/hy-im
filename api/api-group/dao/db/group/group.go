package group

import "hy-im/api/api-group/dao/db/model"

type ImGroupDao interface {
	SaveGroupInfo(info *model.GroupInfo) error
}