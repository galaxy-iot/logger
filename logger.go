package logger

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sync"
)

type LoggingInterface interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})

	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Close()

	SetOutput(w io.Writer) *logger
	SetLevel(level LoggingLevel) *logger
}

type LoggingLevel int

const (
	DEBUG_LEVEL LoggingLevel = iota
	INFO_LEVEL
	WARN_LEVEL
	ERROR_LEVEL
	FATAL_LEVEL
)

func (level LoggingLevel) String() string {
	switch level {
	case DEBUG_LEVEL:
		return "Debug"
	case INFO_LEVEL:
		return "Info"
	case WARN_LEVEL:
		return "Warn"
	case ERROR_LEVEL:
		return "Error"
	case FATAL_LEVEL:
		return "Fatal"
	default:
		return "Unknown"
	}
}

type LoggerConfig struct {
	Level        LoggingLevel
	EnableCaller bool
	CallerLevel  int
	Out          io.Writer
	Name         string

	Formater func(string, LoggingLevel, int, *BufferPool, string, ...interface{}) *bytes.Buffer
}

func DefaultLoggerConfig() *LoggerConfig {
	return &LoggerConfig{
		Level:        INFO_LEVEL,
		EnableCaller: true,
		CallerLevel:  2,
		Out:          os.Stdout,

		Formater: DefaultFormater,
	}
}

type logger struct {
	mux sync.Mutex

	Pool *BufferPool
	LoggerConfig
}

func NewLogger(conf *LoggerConfig) LoggingInterface {
	if conf == nil {
		conf = DefaultLoggerConfig()
	}

	if conf.Level < DEBUG_LEVEL || conf.Level > FATAL_LEVEL {
		conf.Level = INFO_LEVEL
	}

	if conf.Out == nil {
		conf.Out = os.Stdout
	}

	return &logger{
		Pool:         NewBufferPool(),
		mux:          sync.Mutex{},
		LoggerConfig: *conf,
	}
}

func (l *logger) SetLevel(level LoggingLevel) *logger {
	if level < DEBUG_LEVEL || level > FATAL_LEVEL {
		level = INFO_LEVEL
	}

	l.Level = level
	return l
}

func (l *logger) SetFormatter(formater Formatter) *logger {
	l.Formater = formater
	return l
}

func (l *logger) SetOutput(output io.Writer) *logger {
	l.Out = output
	return l
}

func (l *logger) SetModuleName(name string) *logger {
	l.Name = name
	return l
}

func (l *logger) SetCaller(enableCaller bool, callLevel int) *logger {
	l.EnableCaller = enableCaller
	l.CallerLevel = callLevel
	return l
}

func (l *logger) Open() {

}

func (l *logger) Close() {
	l.mux.Lock()
	defer l.mux.Unlock()
	if closer, ok := l.Out.(io.WriteCloser); ok {
		closer.Close()
	}
}

func (l *logger) Write(buf *bytes.Buffer) {
	l.mux.Lock()
	defer l.mux.Unlock()

	if _, err := l.Out.Write(buf.Bytes()); err != nil {
		panic(err)
	}

	l.Pool.Put(buf)
}

func (l *logger) Debug(args ...interface{}) {
	if l.Level <= DEBUG_LEVEL {
		s := fmt.Sprint(args...)
		l.Write(l.Formater(l.Name, DEBUG_LEVEL, l.CallerLevel, l.Pool, s))
	}
}

func (l *logger) Info(args ...interface{}) {
	if l.Level <= INFO_LEVEL {
		s := fmt.Sprint(args...)
		l.Write(l.Formater(l.Name, INFO_LEVEL, l.CallerLevel, l.Pool, s))
	}
}

func (l *logger) Warn(args ...interface{}) {
	if l.Level <= WARN_LEVEL {
		s := fmt.Sprint(args...)
		l.Write(l.Formater(l.Name, WARN_LEVEL, l.CallerLevel, l.Pool, s))
	}
}

func (l *logger) Error(args ...interface{}) {
	if l.Level <= ERROR_LEVEL {
		s := fmt.Sprint(args...)
		l.Write(l.Formater(l.Name, ERROR_LEVEL, l.CallerLevel, l.Pool, s))
	}
}

func (l *logger) Fatal(args ...interface{}) {
	if l.Level <= FATAL_LEVEL {
		s := fmt.Sprint(args...)
		l.Write(l.Formater(l.Name, FATAL_LEVEL, l.CallerLevel, l.Pool, s))
	}
	os.Exit(1)
}

func (l *logger) Debugf(format string, args ...interface{}) {
	if l.Level <= DEBUG_LEVEL {
		l.Write(l.Formater(l.Name, DEBUG_LEVEL, l.CallerLevel, l.Pool, format, args...))
	}
}

func (l *logger) Infof(format string, args ...interface{}) {
	if l.Level <= INFO_LEVEL {
		l.Write(l.Formater(l.Name, INFO_LEVEL, l.CallerLevel, l.Pool, format, args...))
	}
}

func (l *logger) Warnf(format string, args ...interface{}) {
	if l.Level <= WARN_LEVEL {
		l.Write(l.Formater(l.Name, WARN_LEVEL, l.CallerLevel, l.Pool, format, args...))
	}
}

func (l *logger) Errorf(format string, args ...interface{}) {
	if l.Level <= ERROR_LEVEL {
		l.Write(l.Formater(l.Name, ERROR_LEVEL, l.CallerLevel, l.Pool, format, args...))
	}
}

func (l *logger) Fatalf(format string, args ...interface{}) {
	if l.Level <= FATAL_LEVEL {
		l.Write(l.Formater(l.Name, FATAL_LEVEL, l.CallerLevel, l.Pool, format, args...))
	}
	os.Exit(1)
}
