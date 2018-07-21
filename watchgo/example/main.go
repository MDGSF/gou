package main

import (
	"fmt"
	"os"
	"time"

	"github.com/MDGSF/utils/watchgo"
)

func main() {
	watchgo.Register("watchgo", 2, func() {
		fmt.Println("expired")
		os.Exit(1)
	})

	for {
		time.Sleep(time.Second)
		watchgo.Feed("watchgo")
	}
}
