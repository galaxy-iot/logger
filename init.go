package logger

func init() {
	globalLogger = NewLogger(&LoggerConfig{
		Name:         "global",
		Level:        INFO_LEVEL,
		CallerLevel:  3,
		EnableCaller: true,
		Formater:     globalLogFormatter,
	})
}
