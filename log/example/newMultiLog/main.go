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
		log.ErrorLevel,
		log.NotTerminal)

	multilog := mlog.New()
	multilog.AddOneLogger(mylog1)
	multilog.AddOneLogger(mylog2)

	multilog.Println("KDDMJ Start, TDHMJ Start, TJMJ Start")

	multilog.Error("I'm error level log = %d", 12)
	multilog.Errorf("I'm errorf level log = %d", 12)
	multilog.Errorln("I'm errorln level log")

	multilog.Warn("I'm warnf level log = %d", 13)
	multilog.Warnf("I'm warnf level log = %d", 13)
	multilog.Warnln("I'm warnln level log")

	multilog.Info("I'm info level log = %d", 14)
	multilog.Infof("I'm infof level log = %d", 14)
	multilog.Infoln("I'm infoln level log")

	multilog.Debug("I'm debug level log = %d", 15)
	multilog.Debugf("I'm debugf level log = %d", 15)
	multilog.Debugln("I'm debugln level log")

	multilog.Verbose("I'm verbose level log = %d", 16)
	multilog.Verbosef("I'm verbosef level log = %d", 16)
	multilog.Verboseln("I'm verboseln level log")
}
