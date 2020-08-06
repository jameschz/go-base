package redis

import (
	"go-base/lib/gcache/base"
	"time"

	"github.com/go-redis/redis"
)

var (
	_c_timeout = time.Hour
)

type Redis struct {
	base.Cache               // extends Driver, NodeName
	Conn       *redis.Client // redis connection
}

func (r *Redis) Connect(node string) error {
	r.Conn = redis.NewClient(&redis.Options{
		Addr: node,
	})
	return nil
}

func (r *Redis) Close() error {
	if r.Conn != nil {
		return r.Conn.Close()
	}
	return nil
}

func (r *Redis) Shard(k string) error {
	// if not connected
	if r.Conn == nil {
		if r.Region != nil {
			// shard by region
			return r.Connect(r.Region.Driver.GetShardNode(k))
		} else {
			// shard by driver
			return r.Connect(r.Driver.GetShardNode(k))
		}
	}
	return nil
}

func (r *Redis) Set(k string, v string) error {
	// default exp
	exp := _c_timeout
	// use region exp
	if r.Region != nil {
		exp = r.Region.GetExp()
	}
	return r.SetTTL(k, v, exp)
}

func (r *Redis) SetTTL(k string, v string, exp time.Duration) error {
	// use region key
	if r.Region != nil {
		k = r.Region.GetKey(k)
	}
	// sharding
	r.Shard(k)
	// set kv
	_, err := r.Conn.Set(k, v, exp).Result()
	if err != nil {
		return err
	}
	return nil
}

func (r *Redis) Get(k string) (string, error) {
	// use region key
	if r.Region != nil {
		k = r.Region.GetKey(k)
	}
	// sharding
	r.Shard(k)
	// get kv
	v, err := r.Conn.Get(k).Result()
	if err != nil {
		return "", err
	}
	return v, nil
}

func (r *Redis) Del(k string) error {
	// use region key
	if r.Region != nil {
		k = r.Region.GetKey(k)
	}
	// sharding
	r.Shard(k)
	// del kv
	_, err := r.Conn.Del(k).Result()
	if err != nil {
		return err
	}
	return nil
}
