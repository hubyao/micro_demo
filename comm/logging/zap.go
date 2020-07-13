package logging

import (
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	// LevelDebug  Very verbose messages for debugging specific issues
	LevelDebug = "debug"
	// LevelInfo Default log level, informational
	LevelInfo = "info"
	// LevelWarn Warnings are messages about possible issues
	LevelWarn = "warn"
	// LevelError Errors are messages about things we know are problems
	LevelError = "error"
)

var logger *zap.SugaredLogger
var level = zapcore.DebugLevel
var once sync.Once
var sugaredLogger *zap.SugaredLogger

// getZapLevel 等级
func getZapLevel(level string) zapcore.Level {
	switch level {
	case LevelInfo:
		return zapcore.InfoLevel
	case LevelWarn:
		return zapcore.WarnLevel
	case LevelDebug:
		return zapcore.DebugLevel
	case LevelError:
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}

// NewLogger 实例化
func NewLogger(l zapcore.Level) (lg *zap.SugaredLogger, err error) {

	logger, err := SetLogger(l)
	if err != nil {
		return
	}

	lg = logger.Sugar()
	return
}

// SetLogger 设置
func SetLogger(l zapcore.Level) (*zap.Logger, error) {

	lowPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return  lev >= zap.DebugLevel
	})

	// err
	highPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev >= zap.ErrorLevel
	})

	var lowFilename, highFilename string

	// 判断是否是linux环境
	if runtime.GOOS == "linux" {
		lowFilename = "./.log/all.log" // TODO: 日志会自动输出到stdout.log文件中如果日志地址重新设置stdout.log会出现重复输出
		highFilename = "./.log/err.log"
	}

	highCore, err := setLumberjack(highFilename, highPriority)
	if err != nil {
		return nil, err
	}

	lowCore, err := setLumberjack(lowFilename, lowPriority)
	if err != nil {
		return nil, err
	}

	// 开启文件及行号
	development := zap.Development()
	addStacktrace := zap.AddStacktrace(zap.ErrorLevel)

	zapLogger := zap.New(
		zapcore.NewTee(
			highCore,
			lowCore,
		),
		zap.AddCaller(),
		addStacktrace,
		development,
	)

	return zapLogger, nil
}

// 设置切割
func setLumberjack(filename string, priority zapcore.LevelEnabler) (core zapcore.Core, err error) {

	hook := zapcore.AddSync(&lumberjack.Logger{
		MaxSize:    50,    // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 30,    // 日志文件最多保存多少个备份
		MaxAge:     7,     // 文件最多保存多少天
		Compress:   false, // 是否压缩
		Filename:   filename,
	})

	core = zapcore.NewCore(
		zapcore.NewConsoleEncoder(newEncoderConfig()),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), hook), priority)

	return core, nil
}

// 格式化时间
func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

// newEncoderConfig 解析模式
func newEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		// Keys can be anything except the empty string.
		TimeKey:        "T",                           // json时时间键
		LevelKey:       "L",                           // json时日志等级键
		NameKey:        "N",                           // json时日志记录器名
		CallerKey:      "C",                           // json时日志文件信息键
		MessageKey:     "M",                           // json时日志消息键
		StacktraceKey:  "S",                           // json时堆栈键
		LineEnding:     zapcore.DefaultLineEnding,     // 友好日志换行符
		EncodeLevel:    zapcore.CapitalLevelEncoder,   // 友好日志等级名大小写（info INFO）
		EncodeTime:     timeEncoder,                   // 友好日志时日期格式化
		EncodeDuration: zapcore.StringDurationEncoder, // 时间序列化
		EncodeCaller:   zapcore.ShortCallerEncoder,    // 日志文件信息（包/文件.go:行号）
	}
}

// Logger 初始化
func Logger() *zap.SugaredLogger {
	if logger != nil {
		return logger
	}

	// 单例模式只运行一次
	once.Do(func() {
		logger, err := NewLogger(level)
		if err != nil {
			panic(err)
		}
		sugaredLogger = logger
	})

	return sugaredLogger
}

// SetLogLevel 设置日志等级
func SetLogLevel(l string) {
	level = getZapLevel(l)
	logger = nil
}
