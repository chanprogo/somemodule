package log

import (
	"fmt"
	"strings"

	"github.com/astaxie/beego/logs"
)

var Logger *log

type log struct {
	logPath   string
	logLevel  int
	beeLogger *logs.BeeLogger
}

func NewLogger(logPath, logLevel string) {
	Logger = &log{
		logPath:   logPath,
		beeLogger: logs.NewLogger(1000),
	}

	// 输出文件名和行号
	Logger.beeLogger.EnableFuncCallDepth(true)
	Logger.beeLogger.SetLogFuncCallDepth(3)

	switch logLevel {
	case "debug":
		Logger.logLevel = logs.LevelDebug
	case "info":
		Logger.logLevel = logs.LevelInfo
	case "error":
		Logger.logLevel = logs.LevelError
	default:
		Logger.logLevel = logs.LevelDebug
	}

	Logger.beeLogger.SetLogger(logs.AdapterFile, fmt.Sprintf(`{"filename":"%s","level":%d,"maxlines":0,"maxsize":0,"daily":true,"maxdays":60}`, Logger.logPath, Logger.logLevel))
}

func (l *log) Debug(v ...interface{}) {
	l.beeLogger.Debug(l.generateFmtStr(len(v)), v...)
}

func (l *log) Info(v ...interface{}) {
	l.beeLogger.Info(l.generateFmtStr(len(v)), v...)
}

func (l *log) Error(v ...interface{}) {
	l.beeLogger.Error(l.generateFmtStr(len(v)), v...)
}

func (l *log) generateFmtStr(n int) string {
	return strings.Repeat("%v ", n)
}
