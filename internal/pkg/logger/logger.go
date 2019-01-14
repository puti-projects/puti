package logger

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

// Logger zap logger
var logger *zap.Logger

// InitLogger init zap logger
func InitLogger(runmode string) *zap.Logger {
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:   viper.GetString("log.logger_file"),
		MaxSize:    viper.GetInt("log.logger_max_size"), // megabytes
		MaxBackups: viper.GetInt("log.logger_max_backups"),
		MaxAge:     viper.GetInt("log.logger_max_age"), // days
	})

	if runmode == "release" {
		logger = InitProductionLogger(w)
	} else {
		logger = InitDevelopmentLogger(w)
	}

	return logger
}

// InitProductionLogger init the logger for production environment
func InitProductionLogger(w zapcore.WriteSyncer) *zap.Logger {
	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	consoleErrors := zapcore.Lock(os.Stderr)
	jsonEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewProductionEncoderConfig())

	core := zapcore.NewTee(
		zapcore.NewCore(jsonEncoder, w, highPriority),
		zapcore.NewCore(consoleEncoder, consoleErrors, highPriority),
	)
	logger := zap.New(core)
	defer logger.Sync()

	return logger
}

// InitDevelopmentLogger init the logger for development environment
func InitDevelopmentLogger(w zapcore.WriteSyncer) *zap.Logger {
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel
	})
	consoleDebugging := zapcore.Lock(os.Stdout)
	jsonEncoder := zapcore.NewJSONEncoder(zap.NewDevelopmentEncoderConfig())
	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())

	core := zapcore.NewTee(
		zapcore.NewCore(jsonEncoder, w, lowPriority),
		zapcore.NewCore(consoleEncoder, consoleDebugging, lowPriority),
	)
	logger := zap.New(core)
	defer logger.Sync()

	return logger
}

// Info logs a message at InfoLevel. The message includes any fields passed at the log site.
func Info(msg string, fields ...zapcore.Field) {
	logger.Info(msg, fields...)
}

// Infof uses fmt.Sprintf to log a templated message.
func Infof(template string, args ...interface{}) {
	Info(fmt.Sprintf(template, args...))
}

// Warn logs a message at WarnLevel. The message includes any fields passed at the log site.
func Warn(msg string, args ...zapcore.Field) {
	logger.Warn(msg, args...)
}

// Warnf uses fmt.Sprintf to log a templated message.
func Warnf(template string, args ...interface{}) {
	Warn(fmt.Sprintf(template, args...))
}

// Error logs a message at ErrorLevel. The message includes any fields passed at the log site.
func Error(msg string, args ...zapcore.Field) {
	logger.Error(msg, args...)
}

// Errorf uses fmt.Sprintf to log a templated message.
func Errorf(template string, args ...interface{}) {
	Error(fmt.Sprintf(template, args...))
}

// DPanic logs a message at DPanicLevel. The message includes any fields passed at the log site.
func DPanic(msg string, args ...zapcore.Field) {
	logger.DPanic(msg, args...)
}

// DPanicf uses fmt.Sprintf to log a templated message.
// In development, the logger then panics. (See DPanicLevel for details.)
func DPanicf(template string, args ...interface{}) {
	DPanic(fmt.Sprintf(template, args...))
}

// Panic logs a message at PanicLevel. The message includes any fields passed at the log site.
func Panic(msg string, args ...zapcore.Field) {
	logger.Panic(msg, args...)
}

// Panicf uses fmt.Sprintf to log a templated message, then panics.
func Panicf(template string, args ...interface{}) {
	Panic(fmt.Sprintf(template, args...))
}

// Fatal logs a message at FatalLevel. The message includes any fields passed at the log site.
func Fatal(msg string, args ...zapcore.Field) {
	logger.Panic(msg, args...)
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit.
func Fatalf(template string, args ...interface{}) {
	Fatal(fmt.Sprintf(template, args...))
}
