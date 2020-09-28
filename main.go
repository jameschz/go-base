package main

import (
	"fmt"
	"github.com/jameschz/go-base/lib/config"
	"github.com/jameschz/go-base/lib/util"
	"os"
)

// -- example code start
import (
	"github.com/jameschz/go-base/lib/example/etcdclient"
	"github.com/jameschz/go-base/lib/example/gcacheclient"
	"github.com/jameschz/go-base/lib/example/gdoclient"
	"github.com/jameschz/go-base/lib/example/gmqclient"
	"github.com/jameschz/go-base/lib/example/httpserver"
	"github.com/jameschz/go-base/lib/example/socketclient"
	"github.com/jameschz/go-base/lib/example/socketserver"
	"github.com/jameschz/go-base/lib/logger"
	"net/http"
)

// -- example code end

func help() {
	fmt.Println("============================================")
	fmt.Println("> rootpath", util.GetRootPath())
	fmt.Println("============================================")
	// -- example code start
	fmt.Println(" base http_server : example http server")
	fmt.Println(" base so_server : example socket server")
	fmt.Println(" base so_client : example socket client")
	fmt.Println(" base gdo_query : example gdo query test")
	fmt.Println(" base gdo_insert : example gdo insert test")
	fmt.Println(" base gdo_update : example gdo update test")
	fmt.Println(" base gdo_delete : example gdo delete test")
	fmt.Println(" base gdo_tx_basic : example gdo transaction test")
	fmt.Println(" base gcache_all : example gcache all tests")
	fmt.Println(" base gmq_pub : example gmq publish test")
	fmt.Println(" base gmq_sub : example gmq comsume test")
	fmt.Println(" base etcd_ka : example etcd keepalive test")
	fmt.Println(" base etcd_sync : example etcd sync trans test")
	// -- example code end
	fmt.Println("============================================")
	os.Exit(0)
}

var (
	_serverHTTPAddr   string
	_serverSocketAddr string
)

func init() {
	// init config
	_serverHTTPAddr = config.Load("config").GetString("server.http.addr")
	_serverSocketAddr = config.Load("config").GetString("server.socket.addr")
}

func main() {

	// get args
	args := os.Args
	if len(args) < 2 || args == nil {
		help()
	}

	// main logic
	action := args[1]
	switch action {
	// -- example code start
	case "http_server":
		{
			fmt.Println("> base http server")
			fmt.Println("> listening on", _serverHTTPAddr, "...")
			http.HandleFunc("/", examplehttpserver.HTTPIndex)
			http.HandleFunc("/hello", examplehttpserver.HTTPHello)
			err := http.ListenAndServe(_serverHTTPAddr, nil)
			if err != nil {
				logger.Error("base server err:", err)
			}
		}
	case "so_server":
		{
			fmt.Println("> base socket server")
			fmt.Println("> bind on " + _serverSocketAddr)
			examplesocketserver.Server(_serverSocketAddr)
		}
	case "so_client":
		{
			examplesocketclient.Client(_serverSocketAddr)
		}
	case "gdo_query":
		{
			examplegdoclient.MysqlQueryBasic()
		}
	case "gdo_insert":
		{
			examplegdoclient.MysqlInsertBasic()
		}
	case "gdo_update":
		{
			examplegdoclient.MysqlUpdateBasic()
		}
	case "gdo_delete":
		{
			examplegdoclient.MysqlDeleteBasic()
		}
	case "gdo_tx_basic":
		{
			examplegdoclient.MysqlTxBasic()
		}
	case "gcache_all":
		{
			examplegcacheclient.TestDriver()
			examplegcacheclient.TestRegion()
		}
	case "gmq_pub":
		{
			examplegmqclient.RabbitPub()
		}
	case "gmq_sub":
		{
			examplegmqclient.RabbitSub()
		}
	case "etcd_ka":
		{
			exampleetcdclient.TestKA()
		}
	case "etcd_sync":
		{
			exampleetcdclient.TestSync()
		}
	// -- example code end
	default:
		{
			help()
		}
	}
}
