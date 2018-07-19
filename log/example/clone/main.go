package main

import (
	"fmt"
	"os"
	"time"

	"github.com/MDGSF/utils/log"
)

func main() {
	file2, err := os.OpenFile("log2.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("open file log2.txt failed.")
		return
	}
	//var mylog2 = log.New(file2,
	//	"I'm log2, ",
	//	" ||| roomid = 1024, UserID = 110",
	//	log.Ldate|log.Ltime|log.Lmicroseconds|log.Lshortfile,
	//	log.VerboseLevel,
	//	log.NotTerminal)

	log.SetOutput(file2)
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
	log.SetPrefix("[I'm log prefix] ")
	log.SetSuffix("I'm suffix")

	newLog := log.Clone()
	log.SetPrefix("[I'm prefix] ")
	newLog.SetPrefix("[I'm newlog prefix] ")

	go func() {
		for {
			log.Info("I'm info level log = %d", 14)
			time.Sleep(time.Millisecond)
		}
	}()

	for {
		newLog.Info("I'm info level log = %d", 14)
		time.Sleep(time.Millisecond)
	}

	log.Error("I'm error level log = %d", 12)
	log.Errorf("I'm errorf level log = %d", 12)
	log.Errorln("I'm errorln level log")

	log.Warn("I'm warnf level log = %d", 13)
	log.Warnf("I'm warnf level log = %d", 13)
	log.Warnln("I'm warnln level log")

	log.Info("I'm info level log = %d", 14)
	log.Infof("I'm infof level log = %d", 14)
	log.Infoln("I'm infoln level log")

	log.Debug("I'm debug level log = %d", 15)
	log.Debugf("I'm debugf level log = %d", 15)
	log.Debugln("I'm debugln level log")

	log.Verbose("I'm verbose level log = %d", 16)
	log.Verbosef("I'm verbosef level log = %d", 16)
	log.Verboseln("I'm verboseln level log")

}
