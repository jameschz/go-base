package gutil

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	"os"
	"path/filepath"
	"reflect"
	"runtime/debug"
	"strings"
	"time"

	"github.com/bwmarrin/snowflake"
	uuid "github.com/satori/go.uuid"
)

/////////////////////////////////////////////////////////////////
// debug funcs

// Dump : gutil.Dump
func Dump(vals ...interface{}) {
	fmt.Print("dump> ")
	for _, v := range vals {
		fmt.Printf("%#v ", v)
	}
	fmt.Printf("\n")
}

// Throw : gutil.Throw
func Throw(s string, p ...interface{}) {
	panic(fmt.Sprintf(s, p...))
}

// Catch : gutil.Catch
func Catch() {
	if err := recover(); err != nil {
		fmt.Println("catch>", err)
		debug.PrintStack()
	}
}

/////////////////////////////////////////////////////////////////
// base funcs

var (
	_baseEnv  = os.Getenv("GO_baseEnv")
	_baseRoot = os.Getenv("GO_baseRoot")
)

// SetEnv : gutil.SetEnv
func SetEnv(env string) {
	_baseEnv = env
}

// GetEnv : gutil.GetEnv
func GetEnv() string {
	// get from file
	if len(_baseEnv) == 0 {
		_baseEnv = GetFileContent(GetRootPath() + "/etc/env.txt")
		_baseEnv = strings.Trim(_baseEnv, "\n")
	}
	// default local
	if len(_baseEnv) == 0 {
		_baseEnv = "local"
	}
	return _baseEnv
}

// SetRootPath : gutil.SetRootPath
func SetRootPath(path string) {
	_baseRoot = path
}

// GetRootPath : gutil.GetRootPath
func GetRootPath() string {
	if len(_baseRoot) == 0 {
		_baseRoot, _ = filepath.Abs(".")
	}
	return _baseRoot
}

// GetFileContent : gutil.GetFileContent
func GetFileContent(file string) string {
	f, err := os.Open(file)
	defer f.Close()
	if err != nil {
		return ""
	}
	c, err := ioutil.ReadAll(f)
	if err != nil {
		return ""
	}
	return string(c)
}

/////////////////////////////////////////////////////////////////
// call funcs

// CallFunc : gutil.CallFunc
func CallFunc(fn interface{}, args []interface{}) {
	// check func type
	fnType := reflect.TypeOf(fn).String()
	// if strings.Contains(fnType, "func") == false {
	if !strings.Contains(fnType, "func") {
		Throw("CallFunc error: first param expected to be a func")
	}
	// fill func args
	var as = make([]reflect.Value, len(args))
	i := 0
	for _, v := range args {
		as[i] = reflect.ValueOf(v)
		i++
	}
	// call func with args
	reflect.ValueOf(fn).Call(as)
}

/////////////////////////////////////////////////////////////////
// util funcs

// GetMACs : gutil.GetMACs
func GetMACs() (macs []string) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		Throw("GetMACs error: %v", err)
		return macs
	}
	for _, netInterface := range netInterfaces {
		macAddr := netInterface.HardwareAddr.String()
		if len(macAddr) == 0 {
			continue
		}
		macs = append(macs, macAddr)
	}
	return macs
}

// GetIPs : gutil.GetIPs
func GetIPs() (ips []string) {
	interfaceAddrs, err := net.InterfaceAddrs()
	if err != nil {
		Throw("GetIPs error: %v", err)
		return ips
	}
	for _, interfaceAddr := range interfaceAddrs {
		ipNet, isValidIPNet := interfaceAddr.(*net.IPNet)
		if isValidIPNet && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ips = append(ips, ipNet.IP.String())
			}
		}
	}
	return ips
}

// UUID : gutil.UUID
func UUID() string {
	uuid, err := uuid.NewV4()
	if err != nil {
		Throw("UUID error: %v", err)
	}
	return uuid.String()
}

// SFID : gutil.SFID
func SFID() int64 {
	rand.Seed(time.Now().UnixNano())
	node, err := snowflake.NewNode(rand.Int63n(1023))
	if err != nil {
		Throw("SFID error: %v", err)
	}
	return node.Generate().Int64()
}

// RangeInt : gutil.RangeInt
func RangeInt(f int, t int) []int {
	a := []int{}
	if f <= t {
		for {
			a = append(a, f)
			f++
			if f > t {
				break
			}
		}
	} else {
		for {
			a = append(a, f)
			f--
			if f < t {
				break
			}
		}
	}
	return a
}

// RangeInt64 : gutil.RangeInt64
func RangeInt64(f int64, t int64) []int64 {
	a := []int64{}
	if f <= t {
		for {
			a = append(a, f)
			f++
			if f > t {
				break
			}
		}
	} else {
		for {
			a = append(a, f)
			f--
			if f < t {
				break
			}
		}
	}
	return a
}
