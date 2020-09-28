package gdocluster

// ShardRange : shard by range
type ShardRange struct {
	DbName string
	FrNum  int
	ToNum  int
}

// GetDbName :
func (s *ShardRange) GetDbName() string {
	return s.DbName
}

// IsMatch :
func (s *ShardRange) IsMatch(seqID int, shardNum int) bool {
	if seqID >= s.FrNum && seqID <= s.ToNum {
		return true
	}
	return false
}

// ShardHash : shard by hush
type ShardHash struct {
	DbName string
	ModRes int
}

// GetDbName :
func (s *ShardHash) GetDbName() string {
	return s.DbName
}

// IsMatch :
func (s *ShardHash) IsMatch(seqID int, shardNum int) bool {
	modRes := seqID % shardNum
	if s.ModRes == modRes {
		return true
	}
	return false
}

// Shard : shard interface
type Shard interface {
	GetDbName() string
	IsMatch(seqID int, shardNum int) bool
}
