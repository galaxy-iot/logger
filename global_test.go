package logger

import (
	"testing"
)

func TestGloablPrint(t *testing.T) {
	SetLogLevel(DEBUG_LEVEL)
	Info("info")
	Debug("debug")
	Warn("warn")
	Error("error")
}
