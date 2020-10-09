package main

import (
	"time"

	"github.com/MDGSF/utils/log"

	"github.com/MDGSF/utils/log/mwriter"
)

const M3 = 30 * 1024 * 1024

func main() {
	w := mwriter.New("log.txt", M3, time.Hour)
	log.SetOutput(w)
	log.SetLevel(log.VerboseLevel)
	log.SetIsTerminal(log.NotTerminal)

	for {
		log.Info("huangjian")
	}
}
