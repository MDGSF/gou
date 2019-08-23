package main

import (
	"flag"
	"net/http"
	"sync"
	"time"

	"github.com/MDGSF/utils/log"
	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:12306", "http service address")

var upgrader = websocket.Upgrader{} // use default options

func main() {
	log.Info("server start listen at %v", *addr)
	flag.Parse()
	http.HandleFunc("/", echo)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

func echo(w http.ResponseWriter, r *http.Request) {
	defer func() {
		log.Info("connection closed.")
	}()

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error("upgrade:", err)
		return
	}

	log.Info("new conn comming")
	client := NewClientConn(conn)
	client.Start()
}

type TClientConn struct {
	conn *websocket.Conn
}

func NewClientConn(conn *websocket.Conn) *TClientConn {
	c := &TClientConn{}
	c.conn = conn
	return c
}

func (c *TClientConn) Start() {
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go c.readFromClient()
	go c.writeToClient()

	wg.Wait()
}

func (c *TClientConn) readFromClient() {
	defer func() {
		c.conn.Close()
	}()

	for {
		mt, message, err := c.conn.ReadMessage()
		if err != nil {
			log.Error("read: err = %v", err)
			break
		}

		log.Printf("%v: %s", mt, message)
	}
}

func (c *TClientConn) writeToClient() {
	defer func() {
		c.conn.Close()
	}()

	for {

		data := []byte("test server")
		if err := c.conn.WriteMessage(websocket.BinaryMessage, data); err != nil {
			log.Error("failed to send binary message to: %s", err)
			return
		}

		time.Sleep(time.Millisecond * 500)
	}
}
