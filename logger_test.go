package logger

import (
	"bytes"
	"testing"
)

func TestInfo(t *testing.T) {
	buf := &bytes.Buffer{}
	testMessage := "Test Message"

	logging := NewLogger(nil)
	logging.SetOutput(buf)

	buf.Reset()
	logging.Info(testMessage)

	buf.Reset()
	logging.Infof(testMessage)

	buf.Reset()
	logging.Infof("Test %s", "Message")
}

func TestLevel(t *testing.T) {
	buf := &bytes.Buffer{}
	testMessage := "Test Message"

	logging := NewLogger(nil)
	logging.SetOutput(buf)

	buf.Reset()
	logging.Info(testMessage)

	buf.Reset()
	logging.Debug(testMessage)
}

func BenchmarkInfo(b *testing.B) {
	testMessage := "Test Message"
	buf := &bytes.Buffer{}

	logging := NewLogger(nil)
	logging.SetOutput(buf)

	for i := 0; i < b.N; i++ {
		buf.Reset()
		logging.Info(testMessage)
	}
}
