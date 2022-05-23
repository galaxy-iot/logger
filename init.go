package logger

func init() {
	logger = NewLoggingWithFormater("global", INFO_LEVEL, 3, globalLogFormatter)
}
