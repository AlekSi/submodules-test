package integration

import (
	"testing"

	"github.com/stretchr/testify/assert"

	packv1 "github.com/AlekSi/submodules-test/internal/pack"
	// packv2 "github.com/AlekSi/submodules-test/v2/internal/pack"
	packv3 "github.com/AlekSi/submodules-test/v3/internal/pack"
)

func TestIntegration(t *testing.T) {
	_ = packv1.A{}
	// _ = packv2.B{}
	_ = packv3.B{}

	assert.True(t, true)
}
