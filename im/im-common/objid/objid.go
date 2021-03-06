package objid

import (
	"fmt"
	"github.com/y1015860449/go-tools/hy_utils"
	"sync"
)


var incrCount int64 = 0
var incrLock sync.Mutex

func GetObjectId(sessionType int8, sessionId int64) string {
	nowTime := hy_utils.GetMillisecond()
	incrLock.Lock()
	if incrCount >= 16384 {
		incrCount = 0
	}
	tmp := incrCount
	incrCount++
	incrLock.Unlock()
	pre := uint64(nowTime << 22 | tmp << 8 | int64(sessionType))
	return fmt.Sprintf("%x%x", pre, uint32(sessionId))
}
