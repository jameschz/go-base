package logger

import (
	"testing"
)

func TestLogger(t *testing.T) {
	Info("test info")
	Warn("test warn")
	Error("test error")
}
