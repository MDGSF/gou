// MIT License
//
// Copyright (c) 2019 Huang Jian
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package wsclient

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"sync/atomic"
	"time"

	"github.com/MDGSF/utils/log"
	"github.com/gorilla/websocket"
	"golang.org/x/net/proxy"
)

const (
	// wsOnline means current time, connect to server success.
	wsOnline = 0

	// wsOffline mean current time, can't connect to server.
	wsOffline = 1
)

const (
	retryToOpen = 5 * time.Second

	handshakeWait = 5 * time.Second

	writeWait = 10 * time.Second

	pongWait = 10 * time.Second

	pingPeriod = (pongWait * 6) / 10
)

// TAddress websocket address to connect
type TAddress struct {
	Name     string
	NeedAuth bool
	U        url.URL
}

// TWSClient websocket client
type TWSClient struct {
	serverName  string
	serverAddrs []*TAddress
	curAddr     *TAddress
	conn        *websocket.Conn
	state       int32

	// DoAfterOpenSuccess will be called after connect to server success.
	// If err != nil, will disconnect to the server.
	DoAfterOpenSuccess func(serverName string, curAddr *TAddress, conn *websocket.Conn) error

	ServerBinaryMsgChan chan []byte
	ServerTextMsgChan   chan []byte

	outBinaryChan chan []byte
	outTextChan   chan []byte

	retryToOpen   int64
	handshakeWait int64
	writeWait     int64
	pongWait      int64
	pingPeriod    int64

	proxyAddr string
}

// New create a new websocket client
func New(serverName string, serverAddrs []*TAddress) *TWSClient {
	return &TWSClient{
		serverName:          serverName,
		serverAddrs:         serverAddrs,
		DoAfterOpenSuccess:  func(serverName string, curAddr *TAddress, conn *websocket.Conn) error { return nil },
		ServerBinaryMsgChan: make(chan []byte, 16),
		ServerTextMsgChan:   make(chan []byte, 16),
		outBinaryChan:       make(chan []byte, 16),
		outTextChan:         make(chan []byte, 16),
		state:               wsOffline,
		retryToOpen:         int64(retryToOpen),
		handshakeWait:       int64(handshakeWait),
		writeWait:           int64(writeWait),
		pongWait:            int64(pongWait),
		pingPeriod:          int64(pingPeriod),
	}
}

// SetProxyAddr set proxy address
func (c *TWSClient) SetProxyAddr(addr string) {
	c.proxyAddr = addr
}

// Start start websocket client
func (c *TWSClient) Start() {
	go c.run()
}

func (c *TWSClient) run() {
	defaultRetryToOpen := atomic.LoadInt64(&c.retryToOpen)
	for {
		if c.open() {
			atomic.StoreInt64(&c.retryToOpen, defaultRetryToOpen)
			c.doIOLoop()
		}

		time.Sleep(time.Duration(atomic.LoadInt64(&c.retryToOpen)))

		// if can't connect to server, delay to retry again.
		if atomic.LoadInt64(&c.retryToOpen) < int64(60*time.Second) {
			atomic.StoreInt64(&c.retryToOpen, c.retryToOpen*2)
		}
	}
}

func (c *TWSClient) open() bool {
	log.Info("Trying to connect [%v]", c.serverName)
	for i, addr := range c.serverAddrs {
		header := http.Header{}
		header.Set("Host", addr.Name)

		tlsConfig := &tls.Config{}
		tlsConfig.ServerName = addr.Name

		var dialer websocket.Dialer

		if len(c.proxyAddr) > 0 {
			netDialer, err := proxy.SOCKS5("tcp", c.proxyAddr, nil, proxy.Direct)
			if err != nil {
				log.Error("[%v] proxy sock5 %v faield, err = %v", c.serverName, c.proxyAddr, err)
			} else {
				dialer.NetDial = netDialer.Dial
			}
		}

		dialer.HandshakeTimeout = time.Duration(atomic.LoadInt64(&c.handshakeWait))
		dialer.TLSClientConfig = tlsConfig
		conn, _, err := dialer.Dial(addr.U.String(), header)
		if err != nil {
			log.Error("failed to connect [%s] serverAddrs[%d](%v): %s", c.serverName, i, addr.U.String(), err)
			continue
		}

		c.curAddr = c.serverAddrs[i]
		c.conn = conn
		log.Info("connected to [%s] serverAddrs[%d](%v)", c.serverName, i, addr.U.String())

		if addr.NeedAuth {
			if err := c.DoAfterOpenSuccess(c.serverName, c.curAddr, c.conn); err != nil {
				log.Error("DoAfterOpenSuccess failed, err = %v", err)
				c.conn.Close()
				return false
			}
		}
		return true
	}

	log.Error("unable to connect [%v] server", c.serverName)
	return false
}

func (c *TWSClient) doIOLoop() {
	log.Info("[%s] goroutine start: TWSClient.doIOLoop", c.serverName)
	atomic.StoreInt32(&c.state, wsOnline)
	defer func() {
		log.Info("[%s] goroutine exit: TWSClient.doIOLoop", c.serverName)
		atomic.StoreInt32(&c.state, wsOffline)
	}()

	wg := &sync.WaitGroup{}
	wg.Add(2)

	go c.doReading(wg)
	go c.doWriting(wg)

	wg.Wait()
}

func (c *TWSClient) doReading(wg *sync.WaitGroup) {
	log.Info("[%s] goroutine start: TWSClient.doReading", c.serverName)
	defer func() {
		c.conn.Close()
		wg.Done()
		log.Info("[%s] goroutine exit: TWSClient.doReading", c.serverName)
	}()

	c.conn.SetReadDeadline(time.Now().Add(time.Duration(atomic.LoadInt64(&c.pongWait))))
	c.conn.SetPongHandler(func(string) error {
		c.conn.SetReadDeadline(time.Now().Add(time.Duration(atomic.LoadInt64(&c.pongWait))))
		return nil
	})

	for {
		msgtype, data, err := c.conn.ReadMessage()
		if err != nil {
			log.Error("[%s] ReadMessage error: %s", c.serverName, err)
			return
		}

		if msgtype == websocket.BinaryMessage {
			select {
			case c.ServerBinaryMsgChan <- data:
			default:
				log.Warn("[%s] server binary channel is full, drop binary message", c.serverName)
			}
		} else if msgtype == websocket.TextMessage {
			select {
			case c.ServerTextMsgChan <- data:
			default:
				log.Warn("[%s] server text channel is full, drop text message", c.serverName)
			}
		} else {
			log.Warn("[%s] unexpected message type: %d", c.serverName, msgtype)
		}
	}
}

func (c *TWSClient) doWriting(wg *sync.WaitGroup) {
	log.Info("[%s] goroutine start: TWSClient.doWriting", c.serverName)

	ticker := time.NewTicker(time.Duration(atomic.LoadInt64(&c.pingPeriod)))
	defer func() {
		ticker.Stop()
		c.conn.Close()
		wg.Done()
		log.Info("[%s] goroutine exit: TWSClient.doWriting", c.serverName)
	}()

	for {
		select {
		case data, ok := <-c.outBinaryChan:
			if !ok {
				log.Error("[%s] out binary channel is not ok.", c.serverName)
				return
			}

			c.conn.SetWriteDeadline(time.Now().Add(time.Duration(atomic.LoadInt64(&c.writeWait))))
			if err := c.conn.WriteMessage(websocket.BinaryMessage, data); err != nil {
				log.Error("[%s] send binary message failed, err = %v", c.serverName, err)
				return
			}

		case data, ok := <-c.outTextChan:
			if !ok {
				log.Error("[%s] out text channel is not ok.", c.serverName)
				return
			}

			c.conn.SetWriteDeadline(time.Now().Add(time.Duration(atomic.LoadInt64(&c.writeWait))))
			if err := c.conn.WriteMessage(websocket.TextMessage, data); err != nil {
				log.Error("[%s] send text message failed, err = %v", c.serverName, err)
				return
			}

		case <-ticker.C:
			if err := c.conn.WriteControl(websocket.PingMessage, nil, time.Now().Add(time.Duration(atomic.LoadInt64(&c.writeWait)))); err != nil {
				log.Error("[%s] send ping message failed, err = %v", c.serverName, err)
				return
			}
		}
	}
}

// SendBinary send binary data
func (c *TWSClient) SendBinary(data []byte) error {
	if !c.IsOnline() {
		return fmt.Errorf("offline")
	}

	select {
	case c.outBinaryChan <- data:
	default:
		return fmt.Errorf("binary channel is full, drop binary message")
	}
	return nil
}

// SendBinaryWithTimeout send binary data with timeout
func (c *TWSClient) SendBinaryWithTimeout(data []byte, d time.Duration) error {
	if !c.IsOnline() {
		return fmt.Errorf("offline")
	}

	select {
	case c.outBinaryChan <- data:
	case <-time.After(d):
		return fmt.Errorf("binary channel is full, send binary time out, drop message")
	}
	return nil
}

// SendText send text data
func (c *TWSClient) SendText(data []byte) error {
	if !c.IsOnline() {
		return fmt.Errorf("[%s] offline", c.serverName)
	}

	select {
	case c.outTextChan <- data:
	default:
		return fmt.Errorf("text channel is full, drop text message")
	}
	return nil
}

// SendTextWithTimeout send text data with timeout
func (c *TWSClient) SendTextWithTimeout(data []byte, d time.Duration) error {
	if !c.IsOnline() {
		return fmt.Errorf("[%s] offline", c.serverName)
	}

	select {
	case c.outTextChan <- data:
	case <-time.After(d):
		return fmt.Errorf("text channel is full, send text time out, drop message")
	}
	return nil
}

// SendJSON send json data
func (c *TWSClient) SendJSON(v interface{}) error {
	if !c.IsOnline() {
		return fmt.Errorf("offline")
	}

	data, err := json.Marshal(v)
	if err != nil {
		return err
	}

	select {
	case c.outTextChan <- data:
	default:
		return fmt.Errorf("text channel is full, drop json message")
	}
	return nil
}

// SendJSONWithTimeout send json data with timeout
func (c *TWSClient) SendJSONWithTimeout(v interface{}, d time.Duration) error {
	if c.IsOnline() {
		data, err := json.Marshal(v)
		if err != nil {
			return err
		}

		select {
		case c.outTextChan <- data:
			return nil
		case <-time.After(d):
			return fmt.Errorf("text channel is full, send json time out, drop message")
		}
	}
	return fmt.Errorf("offline")
}

// IsOnline websocket client is connected to websocket server.
func (c *TWSClient) IsOnline() bool {
	state := atomic.LoadInt32(&c.state)
	return state == wsOnline
}

// SetTimeOutRetry set retry to connect to server timeout.
func (c *TWSClient) SetTimeOutRetry(d time.Duration) {
	atomic.StoreInt64(&c.retryToOpen, int64(d))
}

// SetTimeOutHandShake set websocket connection handshake timeout.
func (c *TWSClient) SetTimeOutHandShake(d time.Duration) {
	atomic.StoreInt64(&c.handshakeWait, int64(d))
}

// SetTimeoutWrite set websocket write message timeout.
func (c *TWSClient) SetTimeoutWrite(d time.Duration) {
	atomic.StoreInt64(&c.writeWait, int64(d))
}

// SetTimeoutRead set websocket read message timeout.
func (c *TWSClient) SetTimeoutRead(d time.Duration) {
	atomic.StoreInt64(&c.pongWait, int64(d))
}

// SetTimeOutPing set websocket send ping period.
// c.pingPeriod must be smaller than c.pongWait
func (c *TWSClient) SetTimeOutPing(t int64) {
	atomic.StoreInt64(&c.pingPeriod, t*int64(time.Second))
}
