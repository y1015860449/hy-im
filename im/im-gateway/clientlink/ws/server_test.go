package ws

import (
	"log"
	"testing"
)



func TestNewWebsocketServer(t *testing.T) {
	srv, err := NewWebsocketServer(func(options *Options) {
		options.Addr = ":8899"
		options.IdleTime = 20
		options.ConnManager = nil
	})
	if err != nil {
		log.Printf("error %#v", err)
		return
	}
	_ = srv.Start()
}
