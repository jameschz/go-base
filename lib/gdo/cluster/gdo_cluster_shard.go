package gdocluster

// ShardRange : shard by range
type ShardRange struct {
	DrName string
	FrNum  int64
	ToNum  int64
}

// GetDrName :
func (s *ShardRange) GetDrName() string {
	return s.DrName
}

// IsMatch :
func (s *ShardRange) IsMatch(seqID int64, shardNum int64) bool {
	if seqID >= s.FrNum && seqID <= s.ToNum {
		return true
	}
	return false
}

// ShardHash : shard by hush
type ShardHash struct {
	DrName string
	ModRes int64
}

// GetDrName :
func (s *ShardHash) GetDrName() string {
	return s.DrName
}

// IsMatch :
func (s *ShardHash) IsMatch(seqID int64, shardNum int64) bool {
	modRes := seqID % shardNum
	if s.ModRes == modRes {
		return true
	}
	return false
}

// Shard : shard interface
type Shard interface {
	GetDrName() string
	IsMatch(seqID int64, shardNum int64) bool
}
