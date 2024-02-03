package log

import (
	"github.com/apache/rocketmq-client-go/v2/rlog"
	"github.com/xxl6097/go-glog/glog"
)

//type Logger interface {
//	Debug(msg string, fields map[string]interface{})
//	Info(msg string, fields map[string]interface{})
//	Warning(msg string, fields map[string]interface{})
//	Error(msg string, fields map[string]interface{})
//	Fatal(msg string, fields map[string]interface{})
//	Level(level string)
//	OutputPath(path string) (err error)
//}

const defaultLogPath = "/tmp/rocketmq-client.log"

type defaultLogger struct {
}

func (l *defaultLogger) Debug(msg string, fields map[string]interface{}) {
	if LogDebug {
		return
	}
	if msg == "" && len(fields) == 0 {
		return
	}
	//l.logger.WithFields(fields).Debug(msg)
	glog.StdGLog.Debug(msg, fields)
}

func (l *defaultLogger) Info(msg string, fields map[string]interface{}) {
	if LogDebug {
		return
	}
	if msg == "" && len(fields) == 0 {
		return
	}
	//l.logger.WithFields(fields).Info(msg)
	glog.StdGLog.Info(msg, fields)
}

func (l *defaultLogger) Warning(msg string, fields map[string]interface{}) {
	if LogDebug {
		return
	}
	if msg == "" && len(fields) == 0 {
		return
	}
	//l.logger.WithFields(fields).Warning(msg)
	glog.StdGLog.Warn(msg, fields)
}

func (l *defaultLogger) Error(msg string, fields map[string]interface{}) {
	if LogDebug {
		return
	}
	if msg == "" && len(fields) == 0 {
		return
	}
	//l.logger.WithFields(fields).Error(msg)
	glog.StdGLog.Error(msg, fields)
}

func (l *defaultLogger) Fatal(msg string, fields map[string]interface{}) {
	if LogDebug {
		return
	}
	if msg == "" && len(fields) == 0 {
		return
	}
	//l.logger.WithFields(fields).Fatal(msg)
	glog.StdGLog.Fatal(msg, fields)
}

func (l *defaultLogger) Level(level string) {
}

func (l *defaultLogger) OutputPath(path string) (err error) {
	return
}

var LogDebug bool

func InitMQLog() {
	r := &defaultLogger{}
	rlog.SetLogger(r)
}
