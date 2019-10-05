package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_errMsg(t *testing.T) {
	assert.Equal(t, errMsg("foo"), "ERROR: foo env var is missing or incorrect.")
}

func Test_warnMsg(t *testing.T) {
	assert.Equal(t, warnMsg("foo"), "WARNING: foo env var is missing or incorrect.")
}
