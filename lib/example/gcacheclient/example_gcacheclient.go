package examplegcacheclient

import (
	"time"

	"github.com/jameschz/go-base/lib/gcache"
	gcachepool "github.com/jameschz/go-base/lib/gcache/pool"
	"github.com/jameschz/go-base/lib/gutil"
)

// TestDriver :
func TestDriver() {
	// print debug info
	gcachepool.SetDebug(true)
	// get by driver
	cache := gcache.D("default")
	cache.Set("test0", "test1")
	cache.Close()

	cache = gcache.D("default")
	cache.Set("test1", "test1")

	cache = gcache.D("default")
	cache.Set("test2", "test1")
	cache.Close()

	cache = gcache.D("default")
	cache.Set("test3", "test1")
	cache.Close()

	cache = gcache.D("default")
	cache.Set("test4", "test1")
	cache.Close()

	cache = gcache.D("default")
	cache.Set("test5", "test1")
	cache.Close()
	cache = gcache.D("default")
	cache.Set("test6", "test1")
	cache.Close()
	cache = gcache.D("default")
	cache.Set("test7", "test1")
	cache.Close()
	cache = gcache.D("default")
	cache.Set("test8", "test1")
	cache.Close()
	cache = gcache.D("default")
	cache.Set("test9", "test1")
	cache.Close()
	cache = gcache.D("default")
	cache.Set("test10", "test1")
	cache.Close()
	cache = gcache.D("default")
	cache.Set("test11", "test1")
	cache.Close()

	// test connection timeout
	// for i := 0; i < 5; i++ {
	// 	time.Sleep(10 * time.Second)
	// 	cache = gcache.D("default")
	// 	cache.Set("test11", "test1")
	// 	cache.Close()
	// }

}

// TestRegion :
func TestRegion() {
	// print debug info
	gcachepool.SetDebug(true)
	// get user region
	cache := gcache.R("user")
	gutil.Dump(cache.Set("user1", "user1"))
	gutil.Dump(cache.SetTTL("user2", "user2", 5*time.Second))
	gutil.Dump(cache.Get("user1"))
	gutil.Dump(cache.Del("user1"))

	incr, _ := cache.Incr("incr")
	gutil.Dump("IncrR:", incr)

	incr, _ = cache.Incr("incr")
	gutil.Dump("IncrR:", incr)

	incrBy, _ := cache.IncrBy("incr", 10)
	gutil.Dump("IncrByR:", incrBy)

	cache.Close()
}
