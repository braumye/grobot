package grobot

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUtils_StringBuilder(t *testing.T) {
	assert.Equal(t, "abc", stringBuilder("a", "b", "c"))
}

func TestUtils_NewError_String(t *testing.T) {
	assert.Equal(t, "type:abc", newError("type", "abc").Error())
}

func TestUtils_NewError_Int(t *testing.T) {
	assert.Equal(t, "type:123", newError("type", 123).Error())
}

func TestUtils_NewError_Error(t *testing.T) {
	assert.Equal(t, "type:error", newError("type", errors.New("error")).Error())
}

func TestUtils_NewError_Default(t *testing.T) {
	assert.Equal(t, "type", newError("type", t).Error())
}
