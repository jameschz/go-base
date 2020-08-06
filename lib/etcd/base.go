package etcd

import (
	"go-base/lib/config"
	"context"
	"fmt"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/concurrency"
)

type Conf struct {
	Timeout   int
	Endpoints []string
}

type Etcd struct {
	Client *clientv3.Client
}

// global
var (
	_etcd_conf    *Conf
	_etcd_timeout = 5 * time.Second
)

// private
func _sortStrategy(strategy string) (op clientv3.OpOption) {
	switch strategy {
	case "key|asc":
		op = clientv3.WithSort(clientv3.SortByKey, clientv3.SortAscend)
	case "key|desc":
		op = clientv3.WithSort(clientv3.SortByKey, clientv3.SortDescend)
	case "value|asc":
		op = clientv3.WithSort(clientv3.SortByValue, clientv3.SortAscend)
	case "value|desc":
		op = clientv3.WithSort(clientv3.SortByValue, clientv3.SortDescend)
	}
	if op == nil {
		panic("etcd> bad sort strategy")
	}
	return op
}

func Init() bool {
	// init once
	if _etcd_conf != nil {
		return true
	}
	// init timeout
	confs := config.Load("etcd").GetStringMap("config")
	_etcd_conf = &Conf{Timeout: confs["timeout"].(int)}
	// init endpoints
	endpoints := confs["endpoints"].([]interface{})
	_etcd_conf.Endpoints = make([]string, len(endpoints))
	for id, conf := range endpoints {
		_etcd_conf.Endpoints[id] = conf.(string)
	}
	return true
}

func Client() *Etcd {
	Init() // init config
	c, err := clientv3.New(clientv3.Config{
		Endpoints:   _etcd_conf.Endpoints,
		DialTimeout: time.Duration(_etcd_conf.Timeout) * time.Second,
	})
	if err != nil {
		panic("etcd> init client error")
	}
	etcd := &Etcd{Client: c}
	return etcd
}

func (c *Etcd) Close() {
	c.Client.Close()
}

func (c *Etcd) Put(k string, v string) error {
	ctx, cancel := context.WithTimeout(context.Background(), _etcd_timeout)
	_, err := c.Client.Put(ctx, k, v)
	cancel()
	if err != nil {
		return err
	}
	return nil
}

func (c *Etcd) PutWithLease(k string, v string, ts int64) error {
	ctx, cancel := context.WithTimeout(context.Background(), _etcd_timeout)
	resp, err := c.Client.Grant(ctx, ts)
	cancel()
	if err != nil {
		return err
	}
	ctx, cancel = context.WithTimeout(context.Background(), _etcd_timeout)
	_, err = c.Client.Put(ctx, k, v, clientv3.WithLease(resp.ID))
	cancel()
	if err != nil {
		return err
	}
	return nil

}

func (c *Etcd) Incr(k string) error {
	_, err := concurrency.NewSTM(c.Client, func(s concurrency.STM) error {
		v := s.Get(k)
		v_i := 0
		fmt.Sscanf(v, "%d", &v_i)
		s.Put(k, fmt.Sprintf("%d", (v_i+1)))
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *Etcd) Decr(k string) error {
	_, err := concurrency.NewSTM(c.Client, func(s concurrency.STM) error {
		v := s.Get(k)
		v_i := 0
		fmt.Sscanf(v, "%d", &v_i)
		s.Put(k, fmt.Sprintf("%d", (v_i-1)))
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *Etcd) Get(k string) (res map[string]string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), _etcd_timeout)
	resp, err := c.Client.Get(ctx, k, clientv3.WithPrefix())
	cancel()
	if err != nil {
		return res, err
	} else {
		res = make(map[string]string, len(resp.Kvs))
		for _, v := range resp.Kvs {
			res[string(v.Key)] = string(v.Value)
		}
	}
	return res, err
}

func (c *Etcd) GetWithSort(k string, strategy string) (res map[string]string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), _etcd_timeout)
	resp, err := c.Client.Get(ctx, k, clientv3.WithPrefix(), _sortStrategy(strategy))
	cancel()
	if err != nil {
		return res, err
	} else {
		res = make(map[string]string, len(resp.Kvs))
		for _, v := range resp.Kvs {
			res[string(v.Key)] = string(v.Value)
		}
	}
	return res, err
}

func (c *Etcd) GetWithSortLimit(k string, strategy string, limit int64) (res map[string]string, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), _etcd_timeout)
	resp, err := c.Client.Get(ctx, k, clientv3.WithPrefix(), _sortStrategy(strategy), clientv3.WithLimit(limit))
	cancel()
	if err != nil {
		return res, err
	} else {
		res = make(map[string]string, len(resp.Kvs))
		for _, v := range resp.Kvs {
			res[string(v.Key)] = string(v.Value)
		}
	}
	return res, err
}

func (c *Etcd) Del(k string) error {
	ctx, cancel := context.WithTimeout(context.Background(), _etcd_timeout)
	_, err := c.Client.Delete(ctx, k, clientv3.WithPrefix())
	cancel()
	if err != nil {
		return err
	}
	return nil
}

func (c *Etcd) Sync(mutex_s string, callback func() error) (err error) {
	// new session
	sess, err := concurrency.NewSession(c.Client)
	if err != nil {
		return err
	}
	defer sess.Close()
	// new mutex
	mutex := concurrency.NewMutex(sess, "/_base/sync/"+mutex_s)
	// do lock
	if err = mutex.Lock(context.TODO()); err != nil {
		return err
	}
	// do callback
	if err = callback(); err != nil {
		return err
	}
	// do unlock
	if err = mutex.Unlock(context.TODO()); err != nil {
		return err
	}
	return nil
}

func (c *Etcd) WatchWithPrefix(k string) clientv3.WatchChan {
	return c.Client.Watch(context.Background(), k, clientv3.WithPrefix())
}

func (c *Etcd) KeepAlive(k string, v string, ts int64) error {
	// grant for one more second
	ctx, cancel := context.WithTimeout(context.Background(), _etcd_timeout)
	resp, err := c.Client.Grant(ctx, ts+1)
	cancel()
	if err != nil {
		return err
	}
	ctx, cancel = context.WithTimeout(context.Background(), _etcd_timeout)
	_, err = c.Client.Put(ctx, k, v, clientv3.WithLease(resp.ID))
	cancel()
	if err != nil {
		return err
	}
	// keep alive forever
	var ch = make(chan int)
	go func() {
		for {
			_, err := c.Client.KeepAliveOnce(context.TODO(), resp.ID)
			if err != nil {
				panic("etcd> keep alive error")
			}
			sec := time.Duration(ts)
			time.Sleep(sec * time.Second)
		}
	}()
	<-ch // wait error
	return err
}
