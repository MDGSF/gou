package logs

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/robfig/cron"
)

var levelstr2int = map[string]int{
	"Emergency": logs.LevelEmergency,
	"Alert":     logs.LevelAlert,
	"critical":  logs.LevelCritical,
	"Error":     logs.LevelError,
	"warning":   logs.LevelWarning,
	"Notice":    logs.LevelNotice,
	"info":      logs.LevelInformational,
	"debug":     logs.LevelDebug,
	"trace":     logs.LevelDebug,
}

var (
	log      *logs.BeeLogger
	logPath  = ""
	fileName = "gou"
	maxDays  = 15
	level    = logs.LevelDebug

	//timeFormat = "2006-01-02_15-04-05"
	timeFormat = "2006-01-02"
)

func init() {
	log = logs.NewLogger(10000)
	log.EnableFuncCallDepth(true)
	log.SetLogFuncCallDepth(3)

	glogEnv := os.Getenv("GLOG")
	if len(glogEnv) > 0 {
		if v, ok := levelstr2int[glogEnv]; ok {
			level = v
		}
	}

	logPath = genDefaultLogDir()

	setLogger()

	startDailyCheckTask()
}

func startDailyCheckTask() {
	task := cron.New()
	task.AddFunc("@midnight", setLogger)
	task.Start()
}

func setLogger() {
	fileNameWithCurTime := logPath + fileName + "_" + time.Now().Format(timeFormat) + ".log"
	oldJSON := fmt.Sprintf(`{"filename":"%s","maxDays":%d,"level":%d}`, fileNameWithCurTime, maxDays, level)
	newJSON := strings.Replace(oldJSON, "\\", "\\\\", -1)
	fmt.Println("setLogger = ", oldJSON)
	fmt.Println("setLogger = ", newJSON)
	log.DelLogger(logs.AdapterFile)
	log.SetLogger(logs.AdapterFile, newJSON)
}

func genDefaultLogDir() string {

	fmt.Println("os.Args[0] = ", os.Args[0])

	PathWithExeName := os.Args[0]
	PathName := ""
	ExeName := ""

	i := strings.LastIndex(PathWithExeName, string(os.PathSeparator))
	if i == -1 {
		PathName = "." + string(os.PathSeparator)
		ExeName = PathWithExeName[:]
	} else {
		PathName = PathWithExeName[:i]
		ExeName = PathWithExeName[i+1:]
	}
	fmt.Printf("PathName = %v, ExeName = %v\n", PathName, ExeName)

	logPathName := PathName + string(os.PathSeparator) + "log" + string(os.PathSeparator) + ExeName + string(os.PathSeparator)
	fmt.Println("genDefaultLogDir logPathName = ", logPathName)

	_, err := os.Stat(logPathName)
	if err != nil {
		if err := os.MkdirAll(logPathName, os.ModePerm); err != nil {
			fmt.Println("MkdirAll failed", logPathName, err)
			return "." + string(os.PathSeparator)
		}
	}

	return logPathName
}

// SetLogPath set log path
func SetLogPath(path string) {
	logPath = path
	setLogger()
}

// SetFileNamePrefix set log file name's prefix.
func SetFileNamePrefix(prefix string) {
	fileName = prefix
	setLogger()
}

// Emergency Log EMERGENCY level message.
func Emergency(format string, v ...interface{}) {
	log.Emergency(format, v...)
}

// Alert Log ALERT level message.
func Alert(format string, v ...interface{}) {
	log.Alert(format, v...)
}

// Critical Log CRITICAL level message.
func Critical(format string, v ...interface{}) {
	log.Critical(format, v...)
}

// Error Log ERROR level message.
func Error(format string, v ...interface{}) {
	log.Error(format, v...)
}

// Warn Log WARN level message.
// compatibility alias for Warning()
func Warn(format string, v ...interface{}) {
	log.Warn(format, v...)
}

// Notice Log NOTICE level message.
func Notice(format string, v ...interface{}) {
	log.Notice(format, v...)
}

// Info Log Info level message.
func Info(format string, v ...interface{}) {
	log.Info(format, v...)
}

// Debug Log DEBUG level message.
func Debug(format string, v ...interface{}) {
	log.Debug(format, v...)
}

// Trace Log TRACE level message.
// compatibility alias for Debug()
func Trace(format string, v ...interface{}) {
	log.Trace(format, v...)
}
