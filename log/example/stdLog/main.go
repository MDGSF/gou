package main

import (
	"os"

	"github.com/MDGSF/utils/log"
)

func main() {

	log.SetOutput(os.Stderr)
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
	log.SetPrefix("[I'm prefix] ")
	log.SetSuffix("I'm suffix")

	log.Errorln("I'm error level log")
	log.Warnln("I'm warn level log")
	log.Infoln("I'm info level log")
	log.Debugln("I'm debug level log")
}
