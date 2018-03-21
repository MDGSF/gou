package main

import (
	"time"

	"github.com/MDGSF/gou/logs"
)

func main() {

	logs.SetLogPath("C:\\")
	logs.SetFileNamePrefix("lala")

	logs.Emergency("Emergency")
	logs.Alert("Alert")
	logs.Critical("critical")
	logs.Error("Error")
	logs.Warn("warning")
	logs.Notice("Notice")
	logs.Info("info")
	logs.Debug("debug")
	logs.Trace("trace")

	for {
		time.Sleep(time.Second * 20)
		logs.Debug("debug")
	}
}
