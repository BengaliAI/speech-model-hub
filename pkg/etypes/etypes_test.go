package etypes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEType_SetMessage(t *testing.T) {
	assert := assert.New(t)
	e := EType{
		Code:    1001,
		Message: "File format not supported",
	}
	e.SetMessage("File format not supported")
	assert.Equal(e.Message, "File format not supported")
}

func TestEType_Error(t *testing.T) {
	assert := assert.New(t)
	e := EType{
		Code:    1001,
		Message: "File format not supported",
	}
	assert.Equal(e.Error(), "1001: File format not supported")
}

func TestEType_Error2(t *testing.T) {
	assert := assert.New(t)
	assert.Equal(ERROR_FILE_FORMAT_NOT_SUPPORTED.Message, "File format not supported")
	ERROR_FILE_FORMAT_NOT_SUPPORTED.SetMessage("Text file format not supported")
	assert.NotEqual(ERROR_FILE_FORMAT_NOT_SUPPORTED.Error(), "1001: Text file format not supported")
}
