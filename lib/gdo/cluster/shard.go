package cluster

// shard by range
type ShardRange struct {
	DbName string
	FrNum  int
	ToNum  int
}

func (s *ShardRange) GetDbName() string {
	return s.DbName
}

func (s *ShardRange) IsMatch(seq_id int, shard_num int) bool {
	if seq_id >= s.FrNum && seq_id <= s.ToNum {
		return true
	}
	return false
}

// shard by hush
type ShardHash struct {
	DbName string
	ModRes int
}

func (s *ShardHash) GetDbName() string {
	return s.DbName
}

func (s *ShardHash) IsMatch(seq_id int, shard_num int) bool {
	mod_res := seq_id % shard_num
	if s.ModRes == mod_res {
		return true
	}
	return false
}

// shard interface
type Shard interface {
	GetDbName() string
	IsMatch(seq_id int, shard_num int) bool
}
