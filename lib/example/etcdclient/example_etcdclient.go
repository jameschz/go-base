package exampleetcdclient

import (
	"fmt"
	"github.com/jameschz/go-base/lib/etcd"
	"github.com/jameschz/go-base/lib/util"
	"time"
)

// TestPut :
func TestPut() {
	// init client
	c := etcd.Client()
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
	c := etcd.Client()
	defer c.Close()
	// get resource
	r, err := c.Get("/room")
	if err != nil {
		util.Dump(err)
	}
	util.Dump("test get >>>")
	for k, v := range r {
		util.Dump(k, v)
	}
	// get with sora
	r, err = c.GetWithSort("/host", "value|desc")
	if err != nil {
		util.Dump(err)
	}
	util.Dump("test get sorted >>>")
	for k, v := range r {
		util.Dump(k, v)
	}
}

// TestDel :
func TestDel() {
	// init client
	c := etcd.Client()
	defer c.Close()
	// set resource
	c.Del("/host")
}

// TestKA :
func TestKA() {
	// init client
	c := etcd.Client()
	defer c.Close()
	// keep alive
	c.KeepAlive("/host/ka", "1", 3)
}

// TestSync :
func TestSync() {
	// init client
	c := etcd.Client()
	defer c.Close()
	// do trans
	s1 := "abc"
	c.Sync("mutex2", func() error {
		util.Dump(s1)
		time.Sleep(5 * time.Second)
		return nil
	})
}

// TestWatch :
func TestWatch() {
	// init client
	c := etcd.Client()
	defer c.Close()
	// watch resource
	rch := c.WatchWithPrefix("foo")
	for wresp := range rch {
		for _, ev := range wresp.Events {
			fmt.Printf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}
}
