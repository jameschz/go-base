package gcache

import (
	"testing"
	"time"

	gcacheredis "github.com/jameschz/go-base/lib/gcache/redis"
	"github.com/jameschz/go-base/lib/gutil"
)

func TestSet(t *testing.T) {
	cache := D("default")
	gutil.Dump(cache.Set("test1", "test1"))
	gutil.Dump(cache.SetTTL("test2", "test1", 5*time.Second))
}

func TestGet(t *testing.T) {
	cache := D("default")
	gutil.Dump(cache.Get("test1"))
	gutil.Dump(cache.Get("test2"))
}

func TestDel(t *testing.T) {
	cache := D("default")
	gutil.Dump(cache.Del("test1"))
}

func TestRSet(t *testing.T) {
	cache := R("user")
	gutil.Dump(cache.Set("user1", "user1"))
	gutil.Dump(cache.SetTTL("user2", "user2", 5*time.Second))
}

func TestRGet(t *testing.T) {
	cache := R("user")
	gutil.Dump(cache.Get("user1"))
}

func TestRDel(t *testing.T) {
	cache := R("user")
	gutil.Dump(cache.Del("user1"))
}

func TestQueue(t *testing.T) {
	r := &gcacheredis.Redis{}
	e := r.Connect("127.0.0.1:6379")
	if e != nil {
		gutil.Dump(e)
	}
	r.RedisConn.LPush("q1", 1)
	rs := r.RedisConn.RPop("q1").Val()
	gutil.Dump(rs)
}
