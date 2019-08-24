package main

import (
	"github.com/MDGSF/utils/ini"
)

func main() {
	conf, err := ini.NewConfig("conf.ini")
	if err != nil {
		panic(err)
	}
	conf.Set("--enable_ldw", "0")
	conf.Set("--pcw_level", "1")
	conf.SaveConfigFile("conf2.ini")
}
