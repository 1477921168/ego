package ejob

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/1477921168/ego/core/elog"
)

func TestComponent_new(t *testing.T) {
	comp := newComponent("test-cmp", defaultConfig(), elog.EgoLogger)
	assert.Equal(t, "test-cmp", comp.Name())
	assert.Equal(t, "task.ejob", comp.PackageName())
}
