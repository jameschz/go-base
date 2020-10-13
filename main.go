package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/jameschz/go-base/lib/config"
	"github.com/jameschz/go-base/lib/gutil"

	// -- example code start
	examplegcacheclient "github.com/jameschz/go-base/lib/example/gcacheclient"
	examplegdoclient "github.com/jameschz/go-base/lib/example/gdoclient"
	examplegetcdclient "github.com/jameschz/go-base/lib/example/getcdclient"
	examplegmqclient "github.com/jameschz/go-base/lib/example/gmqclient"
	examplehttpserver "github.com/jameschz/go-base/lib/example/httpserver"
	examplesocketclient "github.com/jameschz/go-base/lib/example/socketclient"
	examplesocketserver "github.com/jameschz/go-base/lib/example/socketserver"
	"github.com/jameschz/go-base/lib/logger"
	// -- example code end
)

func help() {
	fmt.Println("============================================")
	fmt.Println("> rootpath", gutil.GetRootPath())
	fmt.Println("============================================")
	// -- example code start
	fmt.Println("> http_server : example http server")
	fmt.Println("> so_server : example socket server")
	fmt.Println("> so_client : example socket client")
	fmt.Println("> gdo_query : example gdo query test")
	fmt.Println("> gdo_insert : example gdo insert test")
	fmt.Println("> gdo_update : example gdo update test")
	fmt.Println("> gdo_delete : example gdo delete test")
	fmt.Println("> gdo_tx_basic : example gdo transaction test")
	fmt.Println("> gcache_all : example gcache all tests")
	fmt.Println("> gmq_pub : example gmq publish test")
	fmt.Println("> gmq_sub : example gmq comsume test")
	fmt.Println("> etcd_ka : example etcd keepalive test")
	fmt.Println("> etcd_sync : example etcd sync trans test")
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
			examplegetcdclient.TestKA()
		}
	case "etcd_sync":
		{
			examplegetcdclient.TestSync()
		}
	// -- example code end
	default:
		{
			help()
		}
	}
}
