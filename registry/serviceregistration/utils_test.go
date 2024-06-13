package serviceregistration

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetCurrentIP(t *testing.T) {
	ip := GetCurrentIP()
	assert.NotEmpty(t, ip)
}
