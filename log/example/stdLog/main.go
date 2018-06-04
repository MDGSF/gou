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
	//log.SetCallDepth(3)

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

	logWrapperError()
}

func logWrapperError() {
	log.Error("I'm error level log = %d", 12)
}
