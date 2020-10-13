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
	gutil.Dump(cache.Set("test1", "test1"))
	gutil.Dump(cache.SetTTL("test2", "test2", 5*time.Second))
	gutil.Dump(cache.Get("test1"))
	gutil.Dump(cache.Del("test1"))
	cache.Close()
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
	cache.Close()
}
