package log

import "strings"

type Level uint8

const (
	ERROR Level = iota
	WARN
	INFO
	DEBUG
	TRACE
)

var (
	logLevel = INFO
)

func Trace(args ...interface{}) {
	if logLevel >= TRACE {
		log.Debug(args)
	}
}

func Debug(args ...interface{}) {
	if logLevel >= DEBUG {
		log.Debug(args)
	}
}

func Info(args ...interface{}) {
	if logLevel >= INFO {
		log.Info(args)
	}
}

func Warn(args ...interface{}) {
	if logLevel >= WARN {
		log.Warn(args)
	}
}

func Error(args ...interface{}) {
	if logLevel >= ERROR {
		log.Error(args)
	}
}

func SetLevel(level Level) {
	logLevel = level
}

func ChangeLevel(level string) {
	level = strings.ToLower(level)
	switch level {
	case "error":
		SetLevel(ERROR)
	case "warn":
		SetLevel(WARN)
	case "info":
		SetLevel(INFO)
	case "debug":
		SetLevel(DEBUG)
	case "trace":
		SetLevel(TRACE)
	}
}
