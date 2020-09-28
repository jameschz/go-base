package examplesocketclient

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

// Client :
func Client(addr string) {
	// dial and handshake to get connections
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		fmt.Println("> connect error:", err)
		os.Exit(1)
	}
	fmt.Println("> connect success!")
	defer conn.Close()
	// keep connection
	i := 0
	for {
		i++
		// send to server
		words := strconv.Itoa(i) + " Hello I'm heartbeat client."
		msg, err := conn.Write([]byte(words))
		if err != nil {
			fmt.Println("> write error:", err)
			break
		}
		fmt.Println("> send:", msg, "text:", words)
		// lost connection
		if i >= 5 {
			time.Sleep(5 * time.Second)
			continue
		}
		// keep connection
		time.Sleep(1 * time.Second)
	}
}
