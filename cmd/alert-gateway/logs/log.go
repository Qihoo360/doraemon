package logs

import (
	"fmt"
	"strings"

	"github.com/astaxie/beego/logs"
)

var (
	Logger      = logs.NewLogger(10000)
	Alertloger  = logs.NewLogger(10000)
	Originloger = logs.NewLogger(10000)
	Panic       = logs.NewLogger(10000)
)

func init() {
	Logger.SetLogger("file", `{"filename":"logs/logs.txt","level":7}`)
	Logger.SetLogger(logs.AdapterConsole, `{"level":7}`)
	Logger.EnableFuncCallDepth(true)
	Logger.SetLogFuncCallDepth(3)
	Alertloger.SetLogger("file", `{"filename":"logs/alerts.txt","level":7}`)
	Alertloger.SetLogger(logs.AdapterConsole, `{"level":7}`)
	Alertloger.EnableFuncCallDepth(true)
	Alertloger.SetLogFuncCallDepth(3)
	Originloger.SetLogger("file", `{"filename":"logs/origin.txt","level":7}`)
	Originloger.SetLogger(logs.AdapterConsole, `{"level":7}`)
	Originloger.EnableFuncCallDepth(true)
	Originloger.SetLogFuncCallDepth(3)

	Panic.SetLogger("file", `{"filename":"logs/panic.txt","level":7}`)
	Panic.SetLogger(logs.AdapterConsole, `{"level":7}`)
	Panic.EnableFuncCallDepth(true)
	Panic.SetLogFuncCallDepth(3)
}

func Error(f interface{}, v ...interface{}) {
	Logger.Error(formatLog(f, v...))
	return
}

func Warning(f interface{}, v ...interface{}) {
	Logger.Warning(formatLog(f, v...))
	return
}

func Critical(f interface{}, v ...interface{}) {
	Logger.Critical(formatLog(f, v...))
	return
}

func Notice(f interface{}, v ...interface{}) {
	Logger.Notice(formatLog(f, v...))
	return
}

func Info(f interface{}, v ...interface{}) {
	Logger.Info(formatLog(f, v...))
	return
}

func Debug(f interface{}, v ...interface{}) {
	Logger.Debug(formatLog(f, v...))
	return
}

// vendor/github.com/astaxie/beego/logs/log.go
func formatLog(f interface{}, v ...interface{}) string {
	var msg string
	switch f.(type) {
	case string:
		msg = f.(string)
		if len(v) == 0 {
			return msg
		}
		if strings.Contains(msg, "%") && !strings.Contains(msg, "%%") {
			//format string
		} else {
			//do not contain format char
			msg += strings.Repeat(" %v", len(v))
		}
	default:
		msg = fmt.Sprint(f)
		if len(v) == 0 {
			return msg
		}
		msg += strings.Repeat(" %v", len(v))
	}
	return fmt.Sprintf(msg, v...)
}
