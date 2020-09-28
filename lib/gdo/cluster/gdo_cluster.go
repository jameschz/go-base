package gdocluster

import (
	"github.com/jameschz/go-base/lib/config"
	"github.com/jameschz/go-base/lib/gdo/driver"
)

// Cluster :
type Cluster struct {
	Type   string
	Algo   string
	SeqID  string
	Shards []Shard
}

// Shard :
func (c *Cluster) Shard(seqID int) (db string) {
	if len(c.Shards) == 0 {
		panic("gdo> cluster has no shard")
	}
	shardNum := len(c.Shards)
	for _, shard := range c.Shards {
		if shard.IsMatch(seqID, shardNum) {
			db = shard.GetDbName()
			break
		}
	}
	return db
}

var (
	_dbClusters map[string]*Cluster
)

// Init :
func Init() bool {
	// init once
	if len(_dbClusters) > 0 {
		return true
	}
	// init drivers
	gdodriver.Init()
	// init clusters
	clusters := config.Load("database").GetStringMap("clusters")
	_dbClusters = make(map[string]*Cluster, len(clusters))
	for _cName, _cData := range clusters {
		// new cluster ptr
		_cCluster := _cData.(map[string]interface{})
		cluster := &Cluster{}
		cluster.Type = _cCluster["type"].(string)
		cluster.Algo = _cCluster["algo"].(string)
		cluster.SeqID = _cCluster["seqID"].(string)
		// get cluster shards
		_cShards := _cCluster["shards"].([]interface{})
		shards := make([]Shard, len(_cShards))
		var i = 0
		for _, _shardData := range _cShards {
			_cShard := make(map[string]interface{})
			for k, v := range _shardData.(map[interface{}]interface{}) {
				_cShard[k.(string)] = v
			}
			switch cluster.Algo {
			case "range":
				shards[i] = &ShardRange{
					DbName: _cShard["db"].(string),
					FrNum:  _cShard["fr_num"].(int),
					ToNum:  _cShard["to_num"].(int),
				}
			case "hash":
				shards[i] = &ShardHash{
					DbName: _cShard["db"].(string),
					ModRes: _cShard["mod_res"].(int),
				}
			}
			i++
		}
		cluster.Shards = shards
		// check cluster
		if len(cluster.Algo) == 0 ||
			len(cluster.SeqID) == 0 ||
			len(cluster.Shards) == 0 {
			panic("gdo> init cluster error")
		}
		// save cluster
		_dbClusters[_cName] = cluster
	}
	return true
}

// GetCluster :
func GetCluster(cs string) (cluster *Cluster) {
	if _, r := _dbClusters[cs]; !r {
		panic("gdo> can not find cluster")
	}
	cluster = _dbClusters[cs]
	return cluster
}
