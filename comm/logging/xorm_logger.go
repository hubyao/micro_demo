package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"xorm.io/core"
)

var _ core.ILogger = &XormLogger{}

type XormLogger struct {
	level   core.LogLevel
	showSQL bool
	logger  *zap.SugaredLogger
}

func (x *XormLogger) Debug(v ...interface{}) {
	x.logger.Debug(v...)
}
func (x *XormLogger) Debugf(format string, v ...interface{}) {

	// TODO: 屏蔽数据库对时区的debug输出
	if format != "empty zone key[%v] : %v | zone: %v | location: %+v\n" {
		x.logger.Debugf(format, v...)
	}

}
func (x *XormLogger) Error(v ...interface{}) {

	x.logger.Error(v...)

}
func (x *XormLogger) Errorf(format string, v ...interface{}) {
	x.logger.Errorf(format, v...)

}
func (x *XormLogger) Info(v ...interface{}) {
	x.logger.Info(v...)
}

func (x *XormLogger) Infof(format string, v ...interface{}) {
	x.logger.Infof(format, v...)
}
func (x *XormLogger) Warn(v ...interface{}) {
	x.logger.Warn(v...)
}
func (x *XormLogger) Warnf(format string, v ...interface{}) {
	x.logger.Warnf(format, v...)
}

func (x *XormLogger) Level() core.LogLevel {
	return x.level
}

func (x *XormLogger) SetLevel(l core.LogLevel) {
	x.level = l

	level := convertLevel(l)
	var err error
	x.logger, err = NewLogger(level)
	if err != nil {
		panic(err)
	}
	return
}

func (x *XormLogger) ShowSQL(show ...bool) {
	if len(show) == 0 {
		x.showSQL = true
		//fmt.Println("1", x.showSQL )
		return
	}
	x.showSQL = show[0]
	//fmt.Println("2", x.showSQL )
}

func (x *XormLogger) IsShowSQL() bool {
	//fmt.Println("3", x.showSQL )

	return x.showSQL
}

func convertLevel(l core.LogLevel) zapcore.Level {
	switch l {
	case core.LOG_INFO:
		return zapcore.InfoLevel
	case core.LOG_WARNING:
		return zapcore.WarnLevel
	case core.LOG_DEBUG:
		return zapcore.DebugLevel
	case core.LOG_ERR:
		return zapcore.ErrorLevel
	default:
		return zapcore.InfoLevel
	}
}
