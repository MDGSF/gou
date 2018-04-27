package main

import (
	"io"
	"os"

	"github.com/MDGSF/utils/log"
)

func main() {
	var out []io.Writer
	out = append(out, os.Stdout)

	file, err := os.OpenFile("log.txt", os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		out = append(out, file)
	}

	var mylog = log.New(os.Stdout,
		"GameID = kddMJ, ",
		" ||| roomid = 1024, UserID = 110",
		log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile,
		log.VerboseLevel,
		log.IsTerminal)

	mylog.Println("KDDMJ Start, TDHMJ Start, TJMJ Start")

	mylog.Error("I'm error level log = %d", 12)
	mylog.Errorf("I'm errorf level log = %d", 12)
	mylog.Errorln("I'm errorln level log")

	mylog.Warn("I'm warnf level log = %d", 13)
	mylog.Warnf("I'm warnf level log = %d", 13)
	mylog.Warnln("I'm warnln level log")

	mylog.Info("I'm info level log = %d", 14)
	mylog.Infof("I'm infof level log = %d", 14)
	mylog.Infoln("I'm infoln level log")

	mylog.Debug("I'm debug level log = %d", 15)
	mylog.Debugf("I'm debugf level log = %d", 15)
	mylog.Debugln("I'm debugln level log")

	mylog.Verbose("I'm verbose level log = %d", 16)
	mylog.Verbosef("I'm verbosef level log = %d", 16)
	mylog.Verboseln("I'm verboseln level log")
}
