package examplesocketserver

import (
	"fmt"
	"github.com/jameschz/go-base/lib/logger"
	"net"
	"time"
)

func _handleConn(conn net.Conn) {
	buffer := make([]byte, 1024)
	for {
		// recv message
		n, err := conn.Read(buffer)
		if err != nil {
			logger.Warn(conn.RemoteAddr().String(), "buffer error", err)
			break
		}
		// init message
		data := buffer[:n]
		if len(data) > 0 {
			// init channles
			msgS := make(chan []byte)
			msgH := make(chan byte)
			// handle message
			go _getMessage(conn, msgS)
			// handle heartbeat
			go _heartBeating(conn, msgH, 4)
			// send messages to channel
			msgS <- data
			msgH <- data[0]
		}
	}
	// connection closed log
	fmt.Println(">", conn.RemoteAddr().String(), "disconnected.")
	logger.Info(conn.RemoteAddr().String(), "connection closed!")
	// close connection
	conn.Close()
}

func _getMessage(conn net.Conn, msgS chan []byte) {
	select {
	case data := <-msgS:
		// get data string
		dataS := string(data)
		// print message string
		fmt.Println("> get message :", dataS)
		// todo : handle message logic
		// ...
	}
	close(msgS)
}

func _heartBeating(conn net.Conn, msgH chan byte, timeout int) {
	select {
	case hb := <-msgH:
		// set next deadline time
		hbS := string(hb)
		fmt.Println("> get heartbeat :", hbS)
		logger.Info(conn.RemoteAddr().String(), "get heartbeat :", hbS)
		conn.SetDeadline(time.Now().Add(time.Duration(timeout) * time.Second))
		break
	}
	close(msgH)
}

// Server :
func Server(addr string) {
	listener, err := net.Listen("tcp", addr)
	// listen error
	if err != nil {
		fmt.Println("> error:", err)
	}
	defer listener.Close()
	// wait for client
	for {
		conn, err := listener.Accept()
		// accept error
		if err != nil {
			logger.Error(conn.RemoteAddr().String(), "accept err", err)
			continue
		}
		// set timeout
		conn.SetReadDeadline(time.Now().Add(time.Second * 10))
		// connected log
		fmt.Println(">", conn.RemoteAddr().String(), "connected.")
		logger.Info(conn.RemoteAddr().String(), "connect success!")
		// conn handler
		go _handleConn(conn)
	}
}
