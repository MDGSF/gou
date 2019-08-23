package main

import (
	"fmt"
	"net/url"
	"time"

	"github.com/MDGSF/utils/log"
	"github.com/MDGSF/utils/wsclient"
	"github.com/gorilla/websocket"
)

func main() {

	addr1 := &wsclient.TAddress{
		Name: "localhost",
		U: url.URL{
			Scheme: "ws",
			Host:   "192.168.1.178:12306",
			Path:   "/",
		},
	}

	u2, err := url.Parse("ws://127.0.0.1:12306")
	if err != nil {
		panic(err)
	}
	addr2 := &wsclient.TAddress{
		Name: "test",
		U:    *u2,
	}

	addrs := make([]*wsclient.TAddress, 0)
	addrs = append(addrs, addr1)
	addrs = append(addrs, addr2)

	client := wsclient.New("MockServer", addrs)
	client.DoAfterOpenSuccess = DoAfterOpenSuccess
	client.Start()

	ticker := time.NewTicker(time.Second)

	for {
		select {
		case data, ok := <-client.ServerBinaryMsgChan:
			if !ok {
			}
			log.Info("len(data) = %v", len(data))

		case <-ticker.C:
			client.SendBinary([]byte{'a', 'b', 'c'})
			client.SendBinaryWithTimeout([]byte{'d', 'e', 'f'}, time.Second)
			client.SendText([]byte("hello"))
			client.SendTextWithTimeout([]byte("world"), time.Second)

			client.SendJSON(map[string]interface{}{"key": "value"})
			client.SendJSONWithTimeout(map[string]interface{}{"key": "value with timeout"}, time.Second)

			if client.IsOnline() {
				fmt.Println("client is online")
			} else {
				fmt.Println("client is offline")
			}
		}
	}

	shutdown := make(chan bool)
	<-shutdown
}

// DoAfterOpenSuccess check auth
func DoAfterOpenSuccess(serverName string, curAddr *wsclient.TAddress, conn *websocket.Conn) error {
	return nil
}
