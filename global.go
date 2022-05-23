package logger

import (
	"bytes"
	"fmt"
	"io"
	"runtime"
	"strconv"
)

var (
	globalLogger LoggingInterface
)

func globalLogFormatter(name string, level LoggingLevel, callerLevel int, pool *BufferPool, format string, args ...interface{}) *bytes.Buffer {
	var (
		s string
	)

	if len(args) == 0 {
		s = format
	} else {
		s = fmt.Sprintf(format, args...)
	}

	time := CacheTime()
	_, caller, line, _ := runtime.Caller(callerLevel)

	buf := pool.Get()
	buf.Reset()
	buf.WriteString(time)
	buf.WriteString(" ")
	buf.WriteString(caller)
	buf.WriteString(":")
	buf.WriteString(strconv.Itoa(line))
	buf.WriteString(" ")
	buf.WriteString(level.String())
	buf.WriteString(" msg: ")
	buf.WriteString(s)
	buf.WriteString("\n")

	return buf
}

func SetLogLevel(level LoggingLevel) {
	globalLogger.SetLevel(level)
}

func Debug(args ...interface{}) {
	globalLogger.Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	globalLogger.Debugf(format, args...)
}

func Info(args ...interface{}) {
	globalLogger.Info(args...)
}

func Infof(format string, args ...interface{}) {
	globalLogger.Infof(format, args...)
}

func Warn(args ...interface{}) {
	globalLogger.Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	globalLogger.Warnf(format, args...)
}

func Error(args ...interface{}) {
	globalLogger.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	globalLogger.Errorf(format, args...)
}

func Fatal(args ...interface{}) {
	globalLogger.Fatal(args...)
}

func Fatalf(format string, args ...interface{}) {
	globalLogger.Fatalf(format, args...)
}

func Println(args ...interface{}) {
	globalLogger.Info(args...)
}

func Printf(format string, args ...interface{}) {
	globalLogger.Infof(format, args...)
}

func Fatalln(args ...interface{}) {
	globalLogger.Fatal(args...)
}

func SetOutPut(w io.Writer) {
	globalLogger.SetOutput(w)
}
