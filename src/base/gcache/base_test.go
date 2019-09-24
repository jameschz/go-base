package gcache

import (
	"base/gcache/redis"
	"base/util"
	"testing"
	"time"
)

func TestSet(t *testing.T) {
	cache := D("default")
	util.Dump(cache.Set("test1", "test1"))
	util.Dump(cache.SetTTL("test2", "test1", 5*time.Second))
}

func TestGet(t *testing.T) {
	cache := D("default")
	util.Dump(cache.Get("test1"))
	util.Dump(cache.Get("test2"))
}

func TestDel(t *testing.T) {
	cache := D("default")
	util.Dump(cache.Del("test1"))
}

func TestRSet(t *testing.T) {
	cache := R("user")
	util.Dump(cache.Set("user1", "user1"))
	util.Dump(cache.SetTTL("user2", "user2", 5*time.Second))
}

func TestRGet(t *testing.T) {
	cache := R("user")
	util.Dump(cache.Get("user1"))
}

func TestRDel(t *testing.T) {
	cache := R("user")
	util.Dump(cache.Del("user1"))
}

func TestQueue(t *testing.T) {
	r := &redis.Redis{}
	e := r.Connect("127.0.0.1:6379")
	if e != nil {
		util.Dump(e)
	}
	r.Conn.LPush("q1", 1)
	rs := r.Conn.RPop("q1").Val()
	util.Dump(rs)
}
