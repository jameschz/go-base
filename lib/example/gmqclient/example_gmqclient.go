package examplegmqclient

import (
	"github.com/jameschz/go-base/lib/gmq"
	"github.com/jameschz/go-base/lib/gutil"
)

// RabbitPub :
func RabbitPub() {
	// get driver default
	rmq := gmq.D("default")
	// publish to queue default
	rmq.Publish("default", `{name:"james",text:"hello"}`)
}

// RabbitSub :
func RabbitSub() {
	// get driver default
	rmq := gmq.D("default")
	// comsume from queue default
	rmq.Consume("default", func(done chan bool, body []byte) {
		// print messages from queue
		gutil.Dump("done status:", done)
		gutil.Dump("msg body bytes:", body)
		gutil.Dump("msg body string:", string(body))
		// quiting the loop
		done <- true
	})
}
