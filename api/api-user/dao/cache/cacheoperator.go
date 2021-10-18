package cache

import (
	"encoding/json"
	"fmt"
	"github.com/tal-tech/go-zero/core/stores/redis"
)

type CacheOperator struct {
	redisCli *redis.Redis
}

var preTokenKey = "ic:token:%s"

func getTokenKey(token string) string {
	return fmt.Sprintf(preTokenKey, token)
}

func (p *CacheOperator) SaveUserToken(loginToken string, info *UserTokenInfo, expired int64) error {
	key := getTokenKey(loginToken)
	data, err := json.Marshal(info)
	if err != nil {
		return err
	}
	return p.redisCli.Setex(key, string(data), int(expired))
}

func (p *CacheOperator) GetUserToken(loginToken string) (*UserTokenInfo, error) {
	key := getTokenKey(loginToken)
	rest, err := p.redisCli.Get(key)
	if err != nil {
		return nil, err
	}
	var info UserTokenInfo
	if err := json.Unmarshal([]byte(rest), &info); err != nil {
		return nil, err
	}
	return &info, err
}

func (p *CacheOperator) DeleteUserToken(loginToken string) error {
	key := getTokenKey(loginToken)
	_, err := p.redisCli.Del(key)
	return err
}

var prePcLogin = "ic:pcLogin:%s"

func getPcLoginKey(uniqueSign string) string {
	return fmt.Sprintf(prePcLogin, uniqueSign)
}

func (p *CacheOperator) SavePcUniqueSignInfo(uniqueSign string, info *PcUniqueSignInfo, expired int64) error {
	key := getPcLoginKey(uniqueSign)
	data, err := json.Marshal(info)
	if err != nil {
		return err
	}
	return p.redisCli.Setex(key, string(data), int(expired))
}

func (p *CacheOperator) GetPcUniqueSignInfo(uniqueSign string) (*PcUniqueSignInfo, error) {
	key := getPcLoginKey(uniqueSign)
	rest, err := p.redisCli.Get(key)
	if err != nil {
		return nil, err
	}
	if len(rest) <= 0 {
		return nil, nil
	}
	var info PcUniqueSignInfo
	if err := json.Unmarshal([]byte(rest), &info); err != nil {
		return nil, err
	}
	return &info, err
}

func NewCacheOperator(redisCli *redis.Redis) CacheDao {
	return &CacheOperator{redisCli: redisCli}
}
