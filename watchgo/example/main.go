package main

import (
	"fmt"
	"os"
	"time"

	"github.com/MDGSF/utils/watchgo"
)

func main() {
	watchgo.Register("watchgo", 2, func(item *watchgo.TItem) {
		fmt.Println("expired, item = ", item)
		os.Exit(1)
	})

	for {
		time.Sleep(time.Second)
		watchgo.Feed("watchgo")
	}
}
