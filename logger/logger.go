package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// zap日志对象
var zapLogger *zap.Logger

// 初始化日志库
func InitLogger(logLevel string, serviceName string, wxHost string, wxEmail string) {
	hook := lumberjack.Logger{
		Filename:   "./logs/" + serviceName + ".log", // 日志文件路径
		MaxSize:    128,                              // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 30,                               // 日志文件最多保存多少个备份
		MaxAge:     7,                                // 文件最多保存多少天
		Compress:   true,                             // 是否压缩
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "linenum",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}

	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(convertToZapLogLevel(logLevel))

	// 报警代码
	wxpush := NewWxPush(wxHost, wxEmail)

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),                                           // 编码器配置
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // 打印到控制台和文件
		atomicLevel, // 日志级别
	)

	// 错误日志上报
	errorLevel := zap.NewAtomicLevel()
	errorLevel.SetLevel(zap.ErrorLevel)
	errorReportCore := zapcore.NewCore(zapcore.NewJSONEncoder(encoderConfig), zapcore.NewMultiWriteSyncer(wxpush), errorLevel)

	teeCore := zapcore.NewTee(core, errorReportCore)

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()

	// 开启文件及行号
	development := zap.Development()

	// 设置初始化字段
	filed := zap.Fields(zap.String("serviceName", serviceName))

	// 构造日志
	zapLogger = zap.New(teeCore, caller, development, filed)
}

func convertToZapLogLevel(logLevel string) zapcore.Level {
	m := map[string]zapcore.Level{
		"debug": zapcore.DebugLevel,
		"info":  zapcore.InfoLevel,
		"warn":  zapcore.WarnLevel,
		"error": zapcore.ErrorLevel,
	}
	if level, ok := m[logLevel]; ok {
		return level
	}
	return zapcore.InfoLevel
}

func Debug(msg string, fields ...zapcore.Field) {
	zapLogger.Debug(msg, fields...)
}

func Info(msg string, fields ...zapcore.Field) {
	zapLogger.Info(msg, fields...)
}

func Warn(msg string, fields ...zapcore.Field) {
	zapLogger.Warn(msg, fields...)
}

func Error(msg string, fields ...zapcore.Field) {
	zapLogger.Error(msg, fields...)
}
