package cluster

import (
	"go-base/lib/config"
	"go-base/lib/gdo/driver"
)

type Cluster struct {
	Type   string
	Algo   string
	SeqId  string
	Shards []Shard
}

func (c *Cluster) Shard(seq_id int) (db string) {
	if len(c.Shards) == 0 {
		panic("gdo> cluster has no shard")
	}
	shard_num := len(c.Shards)
	for _, shard := range c.Shards {
		if shard.IsMatch(seq_id, shard_num) {
			db = shard.GetDbName()
			break
		}
	}
	return db
}

var (
	_db_clusters map[string]*Cluster
)

func Init() bool {
	// init once
	if len(_db_clusters) > 0 {
		return true
	}
	// init drivers
	driver.Init()
	// init clusters
	clusters := config.Load("database").GetStringMap("clusters")
	_db_clusters = make(map[string]*Cluster, len(clusters))
	for c_name, c_data := range clusters {
		// new cluster ptr
		c_cluster := c_data.(map[string]interface{})
		cluster := &Cluster{}
		cluster.Type = c_cluster["type"].(string)
		cluster.Algo = c_cluster["algo"].(string)
		cluster.SeqId = c_cluster["seq_id"].(string)
		// get cluster shards
		c_shards := c_cluster["shards"].([]interface{})
		shards := make([]Shard, len(c_shards))
		var i = 0
		for _, s_data := range c_shards {
			c_shard := make(map[string]interface{})
			for k, v := range s_data.(map[interface{}]interface{}) {
				c_shard[k.(string)] = v
			}
			switch cluster.Algo {
			case "range":
				shards[i] = &ShardRange{
					DbName: c_shard["db"].(string),
					FrNum:  c_shard["fr_num"].(int),
					ToNum:  c_shard["to_num"].(int),
				}
			case "hash":
				shards[i] = &ShardHash{
					DbName: c_shard["db"].(string),
					ModRes: c_shard["mod_res"].(int),
				}
			}
			i++
		}
		cluster.Shards = shards
		// check cluster
		if len(cluster.Algo) == 0 ||
			len(cluster.SeqId) == 0 ||
			len(cluster.Shards) == 0 {
			panic("gdo> init cluster error")
		}
		// save cluster
		_db_clusters[c_name] = cluster
	}
	return true
}

func GetCluster(cs string) (cluster *Cluster) {
	if _, r := _db_clusters[cs]; !r {
		panic("gdo> can not find cluster")
	}
	cluster = _db_clusters[cs]
	return cluster
}
