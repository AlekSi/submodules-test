package integration

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/AlekSi/submodules-test/v3/internal/pack"
)

func TestIntegration(t *testing.T) {
	_ = pack.B{}

	assert.True(t, true)
}
