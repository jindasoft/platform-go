package xlogger

import "go.uber.org/zap"

func SysDebugf(template string, args ...any) {
	zap.S().Debugf(template, args...)
}

func SysInfof(template string, args ...any) {
	zap.S().Infof(template, args...)
}

func SysWarnf(template string, args ...any) {
	zap.S().Warnf(template, args...)
}

func SysErrorf(template string, args ...any) {
	zap.S().Errorf(template, args...)
}
