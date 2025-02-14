package logger

import (
	"runtime"
	"strconv"
	"strings"
	"testing"
)

func TestDefaultFormatter(t *testing.T) {
	bufferPool := NewBufferPool()
	_, caller, line, _ := runtime.Caller(1)

	testMessage := "Test"
	buf := DefaultFormater("test", INFO_LEVEL, 2, bufferPool, testMessage)

	if strings.Contains(buf.String(), caller+":"+strconv.Itoa(line)+" "+"Info"+" msg: "+testMessage+"\n") {
		t.Error("test failed for default formatter")
		return
	}

	buf1 := DefaultFormater("test", INFO_LEVEL, 2, bufferPool, "Test %d", 1)
	s1 := "Test 1"
	if strings.Contains(buf1.String(), caller+":"+strconv.Itoa(line)+" "+"Info"+" msg: "+s1+"\n") {
		t.Error("test failed for default formatter")
		return
	}

	buf2 := DefaultFormater("test", INFO_LEVEL, 2, bufferPool, "Test %.1f", 1.2)
	s2 := "Test 1.2"
	if strings.Contains(buf2.String(), caller+":"+strconv.Itoa(line)+" "+"Info"+" msg: "+s2+"\n") {
		t.Error("test failed for default formatter")
		return
	}
}
