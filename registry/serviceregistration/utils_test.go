package serviceregistration

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetCurrentIP(t *testing.T) {
	ip, err := GetCurrentIP()
	assert.Nil(t, err)
	assert.NotEmpty(t, ip)
}
