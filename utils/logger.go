package utils

import (
	"os"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	log  *zap.Logger
	once sync.Once
)

func NewLogger() *zap.Logger {
	once.Do(func() {
		log = initLogger()
	})
	return log
}

func initLogger() *zap.Logger {
	var err error
	level := new(zap.AtomicLevel)
	*level = zap.NewAtomicLevel()
	err = level.UnmarshalText([]byte("info"))
	if err != nil {
		os.Exit(1)
	}
	zapConfig := zap.Config{
		Level:             *level,
		Development:       true,
		DisableStacktrace: true,
		Encoding:          "console",
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "T",
			LevelKey:       "L",
			NameKey:        "N",
			CallerKey:      "C",
			MessageKey:     "M",
			StacktraceKey:  "S",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalColorLevelEncoder,
			EncodeTime:     zapcore.TimeEncoderOfLayout("01-02 15:04:05.006"),
			EncodeDuration: zapcore.StringDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}
	log, err = zapConfig.Build()
	if err != nil {
		os.Exit(1)
	}
	return log
}
