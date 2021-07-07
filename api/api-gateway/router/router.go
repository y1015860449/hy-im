package router

import (
	"github.com/gin-gonic/gin"
)

type Options struct {

}

type Option func(*Options)

type router struct {

}

func NewRouter(opts ...Option) *router {
	o := Options{}
	for _, opt := range opts {
		opt(&o)
	}
	return &router{

	}
}

// 总的路由
func (p *router) Route(router gin.IRouter) {

//	mAuth := auth.NewAuth(p.idClient)
//	r := router.Group("/", mAuth.Auth)
//	_ = r
}
