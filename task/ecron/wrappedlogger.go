package ecron

import (
	"github.com/1477921168/ego/core/elog"
)

type wrappedLogger struct {
	*elog.Component
}

// Info logs routine messages about cron's operation.
func (wl *wrappedLogger) Info(msg string, keysAndValues ...interface{}) {
	wl.Component.ZapSugaredLogger().Infow("cron "+msg, keysAndValues...)
}

// Error logs an error condition.
func (wl *wrappedLogger) Error(err error, msg string, keysAndValues ...interface{}) {
	wl.Component.ZapSugaredLogger().Errorw("cron "+msg, append(keysAndValues, "err", err)...)
}
