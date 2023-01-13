package log

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log *zap.SugaredLogger
)

func init() {
	zc := zap.Config{
		Level:             zap.NewAtomicLevelAt(zapcore.DebugLevel),
		Development:       false, // 禁止 Warn 打印栈
		DisableCaller:     false, // 禁止打印文件名
		DisableStacktrace: false, // 禁止 Error 和 Warn 打印栈
		// Sampling:       nil, // 限流
		Encoding: "console", // json, console
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:    "message",
			LevelKey:      "level",
			TimeKey:       "time",
			NameKey:       "name",
			CallerKey:     "caller",
			StacktraceKey: "stacktrace",
			LineEnding:    zapcore.DefaultLineEnding, // 换行符
			EncodeLevel:   zapcore.CapitalColorLevelEncoder,
			EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
				enc.AppendString(t.Format("2006-01-02 15:04:05"))
			}, // 时间
			EncodeCaller: zapcore.ShortCallerEncoder, // 文件名
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}

	logger, err := zc.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}

	log = logger.Sugar()
}

func SetLog(newLog *zap.SugaredLogger) {
	log = newLog
}
