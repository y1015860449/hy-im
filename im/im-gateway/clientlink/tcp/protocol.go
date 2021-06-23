package tcp

import (
	"encoding/binary"
	"github.com/Allenxuxu/gev/connection"
	"github.com/Allenxuxu/ringbuffer"
	"github.com/gobwas/pool/pbytes"
)

const msgLen = 4

type protocol struct {}

func (p *protocol) UnPacket(c *connection.Connection, buffer *ringbuffer.RingBuffer) (interface{}, []byte) {
	if buffer.VirtualLength() > msgLen {
		buf := pbytes.GetLen(msgLen)
		defer pbytes.Put(buf)
		_, _ = buffer.VirtualRead(buf)
		dataLen := binary.BigEndian.Uint32(buf)
		if buffer.VirtualLength() >= int(dataLen) {
			ret := make([]byte, dataLen)
			_, _ = buffer.VirtualRead(ret)
			buffer.VirtualFlush()
			return nil, ret
		} else {
			buffer.VirtualRevert()
		}
	}
	return nil, nil
}

func (p *protocol) Packet(c *connection.Connection, data []byte) []byte {
	dataLen := len(data)
	ret := make([]byte, msgLen + dataLen)
	binary.BigEndian.PutUint32(ret, uint32(dataLen))
	copy(ret[msgLen:], data)
	return ret
}



