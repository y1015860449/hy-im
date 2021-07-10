package cache

import (
	"fmt"
	"github.com/y1015860449/go-tools/hyredis"
)


type cacheOperator struct {
	rdCli *hyredis.HyRedis
}

func (c *cacheOperator) GetRoomNumbers(roomId int64) (map[string]int, error) {
	key := getRoomNumberKey(roomId)
	return c.rdCli.ZRangeWithScores(key, 0, -1)
}

func (c *cacheOperator) AddRoomNumbers(roomId int64, members map[string]int) error {
	key := getRoomNumberKey(roomId)
	return c.rdCli.ZAdd(key, members)
}

func (c *cacheOperator) RemoveNumbers(roomId int64, members []string) error {
	key := getRoomNumberKey(roomId)
	return c.rdCli.ZRem(key, members...)
}

func (c *cacheOperator) GetRoomNumbersByRang(roomId int64, min, max interface{}, limit int) (map[string]int, error) {
	panic("implement me")
}

func NewCacheOperator(rdCli *hyredis.HyRedis) CacheDao {
	return &cacheOperator{rdCli: rdCli}
}


func getRoomNumberKey(roomId int64) string {
	return fmt.Sprintf("room_number:%d", roomId)
}