package main

import (
	"fmt"
	"os"

	"github.com/MDGSF/utils/log"
	"github.com/MDGSF/utils/log/mlog"
)

func main() {

	var mylog1 = log.New(os.Stdout,
		"I'm log1, ",
		" ||| roomid = 1024, UserID = 110",
		log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile,
		log.InfoLevel,
		log.IsTerminal)

	file2, err := os.OpenFile("log2.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Println("open file log2.txt failed.")
		return
	}
	var mylog2 = log.New(file2,
		"I'm log2, ",
		" ||| roomid = 1024, UserID = 110",
		log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile,
		log.VerboseLevel,
		log.NotTerminal)

	mlog.AddOneLogger(mylog1)
	mlog.AddOneLogger(mylog2)

	mlog.Println("KDDMJ Start, TDHMJ Start, TJMJ Start")

	mlog.Error("I'm error level log = %d", 12)
	mlog.Errorf("I'm errorf level log = %d", 12)
	mlog.Errorln("I'm errorln level log")

	mlog.Warn("I'm warnf level log = %d", 13)
	mlog.Warnf("I'm warnf level log = %d", 13)
	mlog.Warnln("I'm warnln level log")

	mlog.Info("I'm info level log = %d", 14)
	mlog.Infof("I'm infof level log = %d", 14)
	mlog.Infoln("I'm infoln level log")

	mlog.Debug("I'm debug level log = %d", 15)
	mlog.Debugf("I'm debugf level log = %d", 15)
	mlog.Debugln("I'm debugln level log")

	mlog.Verbose("I'm verbose level log = %d", 16)
	mlog.Verbosef("I'm verbosef level log = %d", 16)
	mlog.Verboseln("I'm verboseln level log")
}
