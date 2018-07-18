package main

import (
	"time"

	"github.com/MDGSF/utils/cache"
	"github.com/MDGSF/utils/log"
)

func main() {
	log.Info("cache demo")

	bm, err := cache.NewCache("memory", `{"interval":60}`)
	if err != nil {
		log.Error("new cache failed, err = %v", err)
		return
	}

	err = bm.Put("astaxie", 1, 10*time.Second)
	if err != nil {
		log.Error("cache put failed, err = %v", err)
		return
	}

	allkeys := bm.GetAllKeys()
	log.Info("allkeys = %v", allkeys)

	v := bm.Get("astaxie")
	if v == nil {
		log.Error("get value failed")
		return
	} else {
		log.Info("key = astaxie, value = %v", v)
	}

	bExist := bm.IsExist("astaxie")
	if bExist {
		log.Info("exist astaxie")
	} else {
		log.Info("not exist astaxie")
	}

	err = bm.Delete("astaxie")
	if err != nil {
		log.Error("cache delete failed, err = %v", err)
		return
	}
}
