package main

import (
	"github.com/MDGSF/utils/log"

	"github.com/MDGSF/utils/log/mwriter"
)

const M1 = 10 * 1024 * 1024
const M3 = 30 * 1024 * 1024

func main() {
	w := mwriter.New("log.txt", M1, M3)
	log.SetOutput(w)
	log.SetLevel(log.VerboseLevel)
	log.SetIsTerminal(log.NotTerminal)

	for {
		log.Info("huangjian")
	}
}
