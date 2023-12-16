package pkg

import (
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	once      sync.Once
	sharedLog *ZapLogger
)

type ZapLogger struct {
	log  *zap.Logger
	Sync func() error
}

var zapLevels = map[log.Level]zapcore.Level{
	log.LevelDebug: zapcore.DebugLevel,
	log.LevelInfo:  zapcore.InfoLevel,
	log.LevelWarn:  zapcore.WarnLevel,
	log.LevelError: zapcore.ErrorLevel,
	log.LevelFatal: zapcore.FatalLevel,
}

func Logger(name string) log.Logger {
	once.Do(func() {
		sharedLog = newLogger(name)
	})
	return sharedLog
}

func newLogger(name string) *ZapLogger {
	encoder := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
	}

	return &ZapLogger{
		log: newZapLogger(name, encoder),
	}
}

func newZapLogger(name string, encoder zapcore.EncoderConfig) *zap.Logger {
	// 初始化日志文件分割
	writeSyncer := generateLogWriter(name)

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoder),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(writeSyncer)),
		zap.NewAtomicLevelAt(zapcore.InfoLevel),
	)

	return zap.New(core, zap.Development(), zap.AddCaller(), zap.AddCallerSkip(2))
}

// getLogWriter split log
func generateLogWriter(name string) zapcore.WriteSyncer {
	logFileName := fmt.Sprintf("/data/%s/logs/info.log", name)
	// 获取当前日期，用于日志文件名
	today := time.Now().Format("2006-01-02")
	logFileName = fmt.Sprintf("/data/%s/logs/%s-info.log", name, today)

	// 创建 Lumberjack Logger 配置
	lumberJackLogger := &lumberjack.Logger{
		Filename:   logFileName,
		MaxSize:    1000, // 每个日志文件的最大大小（MB）
		MaxBackups: 10,   // 保留的旧日志文件数
		MaxAge:     30,   // 保留的旧日志文件的最大天数
		Compress:   true, // 是否压缩旧日志文件
	}

	return zapcore.AddSync(lumberJackLogger)
}

func (zapLogger *ZapLogger) Log(level log.Level, kvs ...interface{}) error {
	if len(kvs) == 0 || len(kvs)%2 != 0 {
		zapLogger.log.Warn("Key values must appear in pairs", zap.Any("key_values", kvs))
		return nil
	}

	var fields []zap.Field
	for i := 0; i < len(kvs); i += 2 {
		fields = append(fields, zap.Any(fmt.Sprint(kvs[i]), kvs[i+1]))
	}

	logLevel, exists := zapLevels[level]
	if !exists {
		logLevel = zapcore.InfoLevel
	}

	zapLogger.log.With(fields...).Log(logLevel, "")
	return nil
}
