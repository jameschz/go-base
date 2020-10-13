package gmq

import (
	"testing"

	"github.com/jameschz/go-base/lib/gutil"
)

func TestPub(t *testing.T) {
	rmq := D("default")
	rmq.Publish("q1", `{name:"james",text:"hello"}`)
}

func TestSub(t *testing.T) {
	rmq := D("default")
	rmq.Consume("q1", func(done chan bool, body []byte) {
		// print messages from queue
		gutil.Dump("done status:", done)
		gutil.Dump("msg body bytes:", body)
		gutil.Dump("msg body string:", string(body))
		// quiting the loop
		done <- true
	})
}
