package util

import (
	"fmt"
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

func Dump(vals ...interface{}) {
	fmt.Print("dump> ")
	for _, v := range vals {
		fmt.Printf("%#v ", v)
	}
	fmt.Printf("\n")
}

func Throw(s string, p ...interface{}) {
	panic(fmt.Sprintf(s, p...))
}

func Catch() {
	if err := recover(); err != nil {
		fmt.Println("catch>", err)
		debug.PrintStack()
	}
}

/////////////////////////////////////////////////////////////////
// base funcs

var (
	_base_env  = os.Getenv("GO_BASE_ENV")
	_base_root = os.Getenv("GO_BASE_ROOT")
)

func SetEnv(env string) {
	_base_env = env
}

func GetEnv() string {
	if len(_base_env) == 0 {
		_base_env = "local"
	}
	return _base_env
}

func SetRootPath(path string) {
	_base_root = path
}

func GetRootPath() string {
	if len(_base_root) == 0 {
		_base_root, _ = filepath.Abs(".")
	}
	return _base_root
}

/////////////////////////////////////////////////////////////////
// call funcs

func CallFunc(fn interface{}, args []interface{}) {
	// check func type
	fn_type := reflect.TypeOf(fn).String()
	if strings.Contains(fn_type, "func") == false {
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

func GetIPs() (ips []string) {
	interfaceAddrs, err := net.InterfaceAddrs()
	if err != nil {
		Throw("GetIPs error: %v", err)
		return ips
	}
	for _, interfaceAddr := range interfaceAddrs {
		ipNet, isValidIpNet := interfaceAddr.(*net.IPNet)
		if isValidIpNet && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				ips = append(ips, ipNet.IP.String())
			}
		}
	}
	return ips
}

func UUID() string {
	uuid, err := uuid.NewV4()
	if err != nil {
		Throw("UUID error: %v", err)
	}
	return uuid.String()
}

func SFID() int64 {
	rand.Seed(time.Now().UnixNano())
	node, err := snowflake.NewNode(rand.Int63n(1023))
	if err != nil {
		Throw("SFID error: %v", err)
	}
	return node.Generate().Int64()
}

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
