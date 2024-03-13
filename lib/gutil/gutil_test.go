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

func TestXOR(t *testing.T) {
	key := []byte("KEY")
	encrypt := EncryptXOR("testbyjames1中国", key)
	decrypt := DecryptXOR(encrypt, key)
	Dump("XOR encrypt", encrypt)
	Dump("XOR decrypt", decrypt)
	decrypt1 := DecryptXOR("KjY4Ly04Ii1oeazB9KPp/w==", key)
	Dump("XOR decrypt1", decrypt1)
}

func TestThrow(t *testing.T) {
	defer Catch()
	Throw("exception:1")
	Throw("exception:2")
}
