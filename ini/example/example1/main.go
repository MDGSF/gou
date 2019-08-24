package main

import (
	"fmt"

	"github.com/MDGSF/utils/ini"
)

func main() {
	conf, err := ini.NewConfig("conf.ini")
	if err != nil {
		panic(err)
	}
	conf.Set("runmode", "dev")

	Show(conf)

	conf.SaveConfigFile("conf2.ini")
}

// Show print config in console, just for debug.
func Show(conf *ini.Config) {
	fmt.Printf("%v", string(conf.FormatBeauty()))
}
