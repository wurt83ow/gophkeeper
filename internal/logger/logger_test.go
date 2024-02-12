package logger

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewLogger(t *testing.T) {
	logger, err := NewLogger("debug")
	assert.NotNil(t, logger)
	assert.NoError(t, err)

	logger, err = NewLogger("invalid-level")
	assert.Nil(t, logger)
	assert.Error(t, err)
}

func TestLogger_Debug(t *testing.T) {
	logger, _ := NewLogger("debug")
	assert.NotNil(t, logger)

	logger.Debug("Debug message")
}

func TestLogger_Info(t *testing.T) {
	logger, _ := NewLogger("info")
	assert.NotNil(t, logger)

	logger.Info("Info message")
}

func TestLogger_Warn(t *testing.T) {
	logger, _ := NewLogger("warn")
	assert.NotNil(t, logger)

	logger.Warn("Warning message")
}

func TestLogger_Writer(t *testing.T) {
	logger := Logger{}
	assert.NotNil(t, logger.writer())
}
