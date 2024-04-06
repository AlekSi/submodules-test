package integration

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/AlekSi/submodules-test/internal/pack"
)

func TestIntegration(t *testing.T) {
	_ = pack.A{}

	assert.True(t, true)
}
