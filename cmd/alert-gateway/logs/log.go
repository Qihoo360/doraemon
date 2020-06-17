package logs

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/astaxie/beego/logs"
)

var (
	Logger      = logs.NewLogger(10000)
	Alertloger  = logs.NewLogger(10000)
	Originloger = logs.NewLogger(10000)
	Panic       = logs.NewLogger(10000)
)

var (
	// default log dir
	LogDir = "log"
	// log to file config
	// log file will auto rotate, keeping logs for last two weeks
	DefaultFileLoggerConfig = FileLoggerConfig{
		Filename: "log/alert-gateway.log",
		MaxFiles: 14,
		MaxSize:  1 << 26,
		Daily:    true,
		MaxDays:  14,
		Rotate:   true,
		Level:    logs.LevelTrace,
	}
)

type FileLoggerConfig struct {
	Filename string `json:"filename"`
	MaxFiles int    `json:"maxfiles"`
	MaxSize  int    `json:"maxsize"`
	Daily    bool   `json:"daily"`
	MaxDays  int64  `json:"maxdays"`
	Rotate   bool   `json:"rotate"`
	Level    int    `json:"level"`
}

func init() {

	if !isExists(LogDir) {
		if err := os.Mkdir(LogDir, 0755); err != nil {
			panic(err)
		}
	}

	setLogger(Logger, "logs.txt")
	setLogger(Alertloger, "alerts.txt")
	setLogger(Originloger, "origin.txt")
	setLogger(Panic, "panic.txt")
}

func setLogger(logger *logs.BeeLogger, file string) {
	DefaultFileLoggerConfig.Filename = filepath.Join(LogDir, file)
	bs, _ := json.Marshal(DefaultFileLoggerConfig)
	logger.SetLogger(logs.AdapterFile, string(bs))
	logger.SetLogger(logs.AdapterConsole, `{"level":7}`)
	logger.EnableFuncCallDepth(true)
	logger.SetLogFuncCallDepth(3)
}

func isExists(file string) bool {
	f, err := os.Stat(file)
	fmt.Println(f, ":", err)
	return err == nil || !os.IsNotExist(err)
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
