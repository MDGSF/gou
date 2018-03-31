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
		log.DebugLevel,
		log.IsTerminal)

	mylog.Println("KDDMJ Start, TDHMJ Start, TJMJ Start")

	mylog.Errorln("I'm error level log")
	mylog.Errorf("I'm error level log = %d", 12)
	mylog.Warnln("I'm warn level log")
	mylog.Warnf("I'm warn level log = %d", 13)
	mylog.Infoln("I'm info level log")
	mylog.Infof("I'm info level log = %d", 14)
	mylog.Debugln("I'm debug level log")
	mylog.Debugf("I'm debug level log = %d", 15)
}
