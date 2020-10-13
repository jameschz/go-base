package gutil

import (
	"testing"
)

func TestUtil(t *testing.T) {
	Dump("GetRootPath", GetRootPath())
	Dump("GetMACs", GetMACs())
	Dump("GetIPs", GetIPs())
	Dump("UUID", UUID())
	Dump("SFID", SFID())
}

func TestThrow(t *testing.T) {
	defer Catch()
	Throw("exception:1")
	Throw("exception:2")
}
