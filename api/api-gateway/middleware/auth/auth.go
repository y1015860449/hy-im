package auth

import (
	"dm-im/http-gateway/global"
	idProto "github.com/dm-common/proto/identify-srv"
	"github.com/dm-common/wrapper/gintracer"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"strings"
)

type Auth interface {
	Auth(ctx *gin.Context)
}

type auth struct {
	idClient idProto.IdentifyService
}

func NewAuth(idClient idProto.IdentifyService) Auth {
	return &auth{idClient: idClient}
}

func (p *auth) Auth(ctx *gin.Context) {
	if !p.authRequired(ctx) {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		ctx.Abort()
		return
	}
	ctx.Next()
}

func (p *auth) authRequired(ctx *gin.Context) bool {

	authStr := ctx.GetHeader("Authorization")
	if authStr == "" || !strings.Contains(authStr, ":") {
		log.Warnf("invalid auth info (%s)", authStr)
		return false
	}

	split := strings.Split(authStr, ":")
	if len(split) < 2 {
		log.Warnf("invalid auth info (%s)", authStr)
		return false
	}

	userId, err := strconv.ParseInt(split[1], 10, 64)
	if err != nil {
		log.Warnf("parse userId err(%v)", err)
		return false
	}

	resp, err := p.idClient.Check(gintracer.ContextWithSpan(ctx), &idProto.CheckRequest{Token: split[0]})
	if err != nil {
		log.Warnf("check token by rpc err(%v)", err)
		return false
	}

	ctx.Set(global.CtxKeyUserID, userId)

	return resp.Status == idProto.TokenStatus_Normal && resp.UserId == userId
}
