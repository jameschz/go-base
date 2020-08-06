package gmq

import (
	"go-base/lib/util"
	"testing"
)

func TestPub(t *testing.T) {
	rmq := D("default")
	rmq.Publish("q1", `{name:"james",text:"hello"}`)
}

func TestSub(t *testing.T) {
	rmq := D("default")
	rmq.Consume("q1", func(done chan bool, body []byte) {
		// print messages from queue
		util.Dump("done status:", done)
		util.Dump("msg body bytes:", body)
		util.Dump("msg body string:", string(body))
		// quiting the loop
		done <- true
	})
}
