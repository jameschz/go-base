package examplegetcdclient

import (
	"fmt"
	"time"

	"github.com/jameschz/go-base/lib/getcd"
	"github.com/jameschz/go-base/lib/gutil"
)

// TestPut :
func TestPut() {
	// init client
	c := getcd.Client()
	defer c.Close()
	// set resource
	c.Put("/host/192.168.1.1", "123")
	c.Put("/host/102.168.1.2", "333")
	c.Put("/room/100001", "192.168.1.1")
	c.Put("/room/100002", "192.168.1.2")
	c.PutWithLease("/host/tmp", "ttt", 5)
}

// TestGet :
func TestGet() {
	// init client
	c := getcd.Client()
	defer c.Close()
	// get resource
	r, err := c.Get("/room")
	if err != nil {
		gutil.Dump(err)
	}
	gutil.Dump("test get >>>")
	for k, v := range r {
		gutil.Dump(k, v)
	}
	// get with sora
	r, err = c.GetWithSort("/host", "value|desc")
	if err != nil {
		gutil.Dump(err)
	}
	gutil.Dump("test get sorted >>>")
	for k, v := range r {
		gutil.Dump(k, v)
	}
}

// TestDel :
func TestDel() {
	// init client
	c := getcd.Client()
	defer c.Close()
	// set resource
	c.Del("/host")
}

// TestKA :
func TestKA() {
	// init client
	c := getcd.Client()
	defer c.Close()
	// keep alive
	c.KeepAlive("/host/ka", "1", 3)
}

// TestSync :
func TestSync() {
	// init client
	c := getcd.Client()
	defer c.Close()
	// do trans
	s1 := "abc"
	c.Sync("mutex2", func() error {
		gutil.Dump(s1)
		time.Sleep(5 * time.Second)
		return nil
	})
}

// TestWatch :
func TestWatch() {
	// init client
	c := getcd.Client()
	defer c.Close()
	// watch resource
	rch := c.WatchWithPrefix("foo")
	for wresp := range rch {
		for _, ev := range wresp.Events {
			fmt.Printf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}
}
