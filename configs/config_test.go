package configs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	config := InitConfig()

	assert.Equal(t, "Muhamad Wildan Faz", config.AppAuthor, "Invalid Author Name")
}
