package connectionmanger

import (
	"hy-im/im/im-gateway/clientlink/interface"
	"hy-im/im/im-gateway/common"
	"sync"
)

type ConnectionManager interface {
	AddConnection(key string, sender _interface.Connection)
	GetConnection(key string) _interface.Connection
	DelConnection(key, linkToken string)
}

type connManager struct {
	senders map[string]_interface.Connection
	mtx     sync.RWMutex
}

func (p *connManager) AddConnection(key string, sender _interface.Connection) {
	p.mtx.Lock()
	defer p.mtx.Unlock()
	p.senders[key] = sender
}

func (p *connManager) GetConnection(key string) _interface.Connection {
	p.mtx.RLock()
	defer p.mtx.RUnlock()
	sender := p.senders[key]
	return sender
}

func (p *connManager) DelConnection(key, linkToken string) {
	p.mtx.Lock()
	defer p.mtx.Unlock()
	if sender, ok := p.senders[key]; ok && sender != nil {
		if sender.LinkToken() == linkToken {
			delete(p.senders, key)
		}
	}
}

func NewConnManager() ConnectionManager {
	return &connManager{
		senders: make(map[string]_interface.Connection),
		mtx:     sync.RWMutex{},
	}
}

type GroupConnectionManager interface {
	AddConnection(roomId int64, key string, sender _interface.Connection)
	GetConnections(roomId int64) *sync.Map
	DelConnection(roomId int64, key string)
	DelGroup(roomId int64)
}

type groupManager struct {
	groupMembers map[int64]*sync.Map
	mtx          sync.RWMutex
}

func (r *groupManager) AddConnection(roomId int64, key string, sender _interface.Connection) {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	if value, ok := r.groupMembers[roomId]; ok {
		value.Store(key, sender)
	} else {
		var members sync.Map
		members.Store(key, sender)
		r.groupMembers[roomId] = &members
	}
}

func (r *groupManager) GetConnections(roomId int64) *sync.Map {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	if value, ok := r.groupMembers[roomId]; ok {
		return value
	}
	return nil
}

func (r *groupManager) DelConnection(roomId int64, key string) {
	r.mtx.RLock()
	defer r.mtx.RUnlock()
	if value, ok := r.groupMembers[roomId]; ok {
		value.Delete(key)
	}
}

func (r *groupManager) DelGroup(roomId int64) {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	if roomNumbers, ok := r.groupMembers[roomId]; ok {
		roomNumbers.Range(func(key, value interface{}) bool {
			conn := value.(_interface.Connection)
			connCtx := conn.GetContext().(common.ConnectionCtx)
			connCtx.GroupId = 0
			conn.SetContext(connCtx)

			return true
		})
		delete(r.groupMembers, roomId)
	}
}

func NewRoomConnectionManager() GroupConnectionManager {
	return &groupManager{
		groupMembers: make(map[int64]*sync.Map),
		mtx:          sync.RWMutex{},
	}
}
