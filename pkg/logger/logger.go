package logger

import (
	"github.com/natefinch/lumberjack"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"sync"
)

var (
	appLogger Logger
	doOnce    sync.Once
)

func InitLogger(production bool) {
	doOnce.Do(func() {
		writerSyncer := getLogWriter(production)
		encoder := getEncoder(production)
		levelEnabler := getLevelEnabler(production)
		core := zapcore.NewCore(encoder, writerSyncer, levelEnabler)
		logger := zap.New(core, zap.AddCaller())
		appLogger = logger.Sugar()
	})
}

func getLevelEnabler(production bool) zapcore.LevelEnabler {
	/*const (
		// DebugLevel logs are typically voluminous, and are usually disabled in production.
		DebugLevel Level = iota - 1

		// InfoLevel is the default logging priority.
		InfoLevel

		// WarnLevel logs are more important than Info, but don't need individual human review.
		WarnLevel

		// ErrorLevel logs are high-priority. If an application is running smoothly,
		// it shouldn't generate any error-level logs.
		ErrorLevel

		// DPanicLevel logs are particularly important errors. In development the
		// logger panics after writing the message.
		DPanicLevel

		// PanicLevel logs a message, then panics.
		PanicLevel

		// FatalLevel logs a message, then calls os.Exit(1).
		FatalLevel

		_minLevel = DebugLevel
		_maxLevel = FatalLevel
	)*/
	if production {
		return zap.InfoLevel
	}
	return zap.DebugLevel
}

func getEncoder(production bool) zapcore.Encoder {
	if production {
		return zapcore.NewJSONEncoder(
			zapcore.EncoderConfig{
				TimeKey:        "ts",
				LevelKey:       "level",
				NameKey:        "logger",
				CallerKey:      "caller",
				MessageKey:     "msg",
				StacktraceKey:  "stacktrace",
				LineEnding:     zapcore.DefaultLineEnding,
				EncodeLevel:    zapcore.CapitalLevelEncoder,
				EncodeTime:     zapcore.ISO8601TimeEncoder,
				EncodeDuration: zapcore.SecondsDurationEncoder,
				EncodeCaller:   zapcore.ShortCallerEncoder,
			},
		)
	}
	return zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
}

func getLogWriter(production bool) zapcore.WriteSyncer {
	if production {
		lumberJackLogger := &lumberjack.Logger{
			Filename:   viper.GetString("lumber.filename"), // location of log file
			MaxSize:    viper.GetInt("lumber.maxsize"),     // maximum size of log file in MBs, before it is rotated
			MaxBackups: viper.GetInt("lumber.maxbackups"),  // maximum no. of old files to retain
			MaxAge:     viper.GetInt("lumber.maxage"),      // maximum number of days it will retain old files
			Compress:   viper.GetBool("lumber.compress"),   // whether to compress/archive old files
		}
		return zapcore.AddSync(lumberJackLogger)
	}
	return os.Stdout
}

/**-- Implement logger --**/

// Debug uses fmt.Sprint to construct and log a message.
func Debug(args ...interface{}) {
	appLogger.Debug(args)
}

// Info uses fmt.Sprint to construct and log a message.
func Info(args ...interface{}) {
	appLogger.Info(args)
}

// Warn uses fmt.Sprint to construct and log a message.
func Warn(args ...interface{}) {
	appLogger.Warn(args)
}

// Error uses fmt.Sprint to construct and log a message.
func Error(args ...interface{}) {
	appLogger.Error(args)
}

// DPanic uses fmt.Sprint to construct and log a message. In development, the
// logger then panics. (See DPanicLevel for details.)
func DPanic(args ...interface{}) {
	appLogger.DPanic(args)
}

// Panic uses fmt.Sprint to construct and log a message, then panics.
func Panic(args ...interface{}) {
	appLogger.Panic(args)
}

// Fatal uses fmt.Sprint to construct and log a message, then calls os.Exit.
func Fatal(args ...interface{}) {
	appLogger.Fatal(args)
}

// Debugf uses fmt.Sprintf to log a templated message.
func Debugf(template string, args ...interface{}) {
	appLogger.Debugf(template, args)
}

// Infof uses fmt.Sprintf to log a templated message.
func Infof(template string, args ...interface{}) {
	appLogger.Infof(template, args)
}

// Warnf uses fmt.Sprintf to log a templated message.
func Warnf(template string, args ...interface{}) {
	appLogger.Warnf(template, appLogger)
}

// Errorf uses fmt.Sprintf to log a templated message.
func Errorf(template string, args ...interface{}) {
	appLogger.Errorf(template, args)
}

// DPanicf uses fmt.Sprintf to log a templated message. In development, the
// logger then panics. (See DPanicLevel for details.)
func DPanicf(template string, args ...interface{}) {
	appLogger.DPanicf(template, args)
}

// Panicf uses fmt.Sprintf to log a templated message, then panics.
func Panicf(template string, args ...interface{}) {
	appLogger.Panicf(template, args)
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit.
func Fatalf(template string, args ...interface{}) {
	appLogger.Fatalf(template, args)
}

// Debugw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
//
// When debug-level logging is disabled, this is much faster than
//  s.With(keysAndValues).Debug(msg)
func Debugw(msg string, keysAndValues ...interface{}) {
	appLogger.Debugw(msg, keysAndValues)
}

// Infow logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func Infow(msg string, keysAndValues ...interface{}) {
	appLogger.Infow(msg, keysAndValues)
}

// Warnw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func Warnw(msg string, keysAndValues ...interface{}) {
	appLogger.Warnw(msg, keysAndValues)
}

// Errorw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func Errorw(msg string, keysAndValues ...interface{}) {
	appLogger.Errorw(msg, keysAndValues)
}

// DPanicw logs a message with some additional context. In development, the
// logger then panics. (See DPanicLevel for details.) The variadic key-value
// pairs are treated as they are in With.
func DPanicw(msg string, keysAndValues ...interface{}) {
	appLogger.DPanicf(msg, keysAndValues)
}

// Panicw logs a message with some additional context, then panics. The
// variadic key-value pairs are treated as they are in With.
func Panicw(msg string, keysAndValues ...interface{}) {
	appLogger.Panicw(msg, keysAndValues)
}

// Fatalw logs a message with some additional context, then calls os.Exit. The
// variadic key-value pairs are treated as they are in With.
func Fatalw(msg string, keysAndValues ...interface{}) {
	appLogger.Fatalw(msg, keysAndValues)
}

// TODO: call before stop
// Sync flushes any buffered log entries.
func Sync() error {
	return appLogger.Sync()
}
