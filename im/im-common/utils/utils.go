package utils

import (
	"fmt"
	"github.com/spaolacci/murmur3"
)

func GetChatId(fromId, toId int64) int64 {
	rest := fmt.Sprintf("%x", fromId + toId)
	return int64(murmur3.Sum64([]byte(rest)))
}

