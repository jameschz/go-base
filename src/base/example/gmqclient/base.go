//package rabbitmqdemo
package gmqclient

import (
	"base/gmq"
	"base/util"
)

func RabbitPub() {
	// get driver default
	rmq := gmq.D("default")
	// publish to queue default
	rmq.Publish("default", `{name:"james",text:"hello"}`)
}

func RabbitSub() {
	// get driver default
	rmq := gmq.D("default")
	// comsume from queue default
	rmq.Consume("default", func(done chan bool, body []byte) {
		// print messages from queue
		util.Dump("done status:", done)
		util.Dump("msg body bytes:", body)
		util.Dump("msg body string:", string(body))
		// quiting the loop
		done <- true
	})
}
