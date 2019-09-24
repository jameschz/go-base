package socketserver

import (
	"base/logger"
	"fmt"
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
			msg_s := make(chan []byte)
			msg_h := make(chan byte)
			// handle message
			go _getMessage(conn, msg_s)
			// handle heartbeat
			go _heartBeating(conn, msg_h, 4)
			// send messages to channel
			msg_s <- data
			msg_h <- data[0]
		}
	}
	// connection closed log
	fmt.Println(">", conn.RemoteAddr().String(), "disconnected.")
	logger.Info(conn.RemoteAddr().String(), "connection closed!")
	// close connection
	conn.Close()
}

func _getMessage(conn net.Conn, msg_s chan []byte) {
	select {
	case data := <-msg_s:
		// get data string
		data_s := string(data)
		// print message string
		fmt.Println("> get message :", data_s)
		// todo : handle message logic
		// ...
	}
	close(msg_s)
}

func _heartBeating(conn net.Conn, msg_h chan byte, timeout int) {
	select {
	case hb := <-msg_h:
		// set next deadline time
		hb_s := string(hb)
		fmt.Println("> get heartbeat :", hb_s)
		logger.Info(conn.RemoteAddr().String(), "get heartbeat :", hb_s)
		conn.SetDeadline(time.Now().Add(time.Duration(timeout) * time.Second))
		break
	}
	close(msg_h)
}

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
