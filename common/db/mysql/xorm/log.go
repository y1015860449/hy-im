package xorm

import (
	log "github.com/sirupsen/logrus"
	"xorm.io/core"
)

type cLog struct {
}

func (p *cLog) Debug(v ...interface{}) {
	log.Debug(v...)
}
func (p *cLog) Debugf(format string, v ...interface{}) {
	log.Debugf(format, v...)

}
func (p *cLog) Error(v ...interface{}) {
	log.Error(v...)
}
func (p *cLog) Errorf(format string, v ...interface{}) {
	log.Errorf(format, v...)
}
func (p *cLog) Info(v ...interface{}) {
	log.Info(v...)
}
func (p *cLog) Infof(format string, v ...interface{}) {
	log.Infof(format, v...)
}
func (p *cLog) Warn(v ...interface{}) {
	log.Warn(v...)
}
func (p *cLog) Warnf(format string, v ...interface{}) {
	log.Warnf(format, v...)
}

func (p *cLog) Level() core.LogLevel {
	switch log.GetLevel() {
	case log.DebugLevel:
		return core.LOG_DEBUG
	case log.ErrorLevel:
		return core.LOG_ERR
	case log.WarnLevel:
		return core.LOG_WARNING
	case log.InfoLevel:
		return core.LOG_INFO
	default:
		return core.LOG_WARNING
	}
}
func (p *cLog) SetLevel(l core.LogLevel) {
	switch l {
	case core.LOG_DEBUG:
		log.SetLevel(log.DebugLevel)
	case core.LOG_ERR:
		log.SetLevel(log.ErrorLevel)
	case core.LOG_WARNING:
		log.SetLevel(log.WarnLevel)
	case core.LOG_INFO:
		log.SetLevel(log.InfoLevel)
	}
}

func (p *cLog) ShowSQL(show ...bool) {
}
func (p *cLog) IsShowSQL() bool {
	if log.GetLevel() <= log.WarnLevel {
		return true
	}
	return false
}
