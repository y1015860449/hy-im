package server

import (
	"github.com/afex/hystrix-go/hystrix"
	"net/http"
)

func StartHystrix(addr string) (err error) {
	r := "81"
	if addr != "" {
		r = addr
	}
	streamHandler := hystrix.NewStreamHandler()
	streamHandler.Start()
	//go http.ListenAndServe(net.JoinHostPort("", "81"), streamHandler)
	go http.ListenAndServe(r, streamHandler)
	return
}
