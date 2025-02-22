package egrpclog

import (
	"sync"

	"github.com/1477921168/ego/core/elog"
)

var (
	once   sync.Once
	logger *elog.Component
)

// Build builds grpclog instance.
func Build() *elog.Component {
	once.Do(func() {
		logger = elog.EgoLogger.With(elog.FieldComponentName("component.grpc"))
	})
	return logger
}
