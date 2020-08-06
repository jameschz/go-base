package main

import (
	"fmt"
	"net/http"
	"os"

	"go-base/lib/config"
	"go-base/lib/example/etcdclient"
	"go-base/lib/example/gcacheclient"
	"go-base/lib/example/gdoclient"
	"go-base/lib/example/gmqclient"
	"go-base/lib/example/httpserver"
	"go-base/lib/example/socketclient"
	"go-base/lib/example/socketserver"
	"go-base/lib/logger"
	"go-base/lib/util"
)

func help() {
	fmt.Println("============================================")
	fmt.Println("> rootpath", util.GetRootPath())
	fmt.Println("============================================")
	fmt.Println(" base server : http server")
	fmt.Println(" base so_server : socket server")
	fmt.Println(" base so_client : socket client")
	fmt.Println(" base gdo_query : gdo query test")
	fmt.Println(" base gdo_insert : gdo insert test")
	fmt.Println(" base gdo_update : gdo update test")
	fmt.Println(" base gdo_delete : gdo delete test")
	fmt.Println(" base gdo_tx_basic : gdo transaction test")
	fmt.Println(" base gcache_all : gcache all tests")
	fmt.Println(" base gmq_pub : gmq publish test")
	fmt.Println(" base gmq_sub : gmq comsume test")
	fmt.Println(" base etcd_ka : etcd keepalive test")
	fmt.Println(" base etcd_sync : etcd sync trans test")
	fmt.Println("============================================")
	os.Exit(0)
}

var (
	_c_serverHttpAddr   string
	_c_serverSocketAddr string
)

func init() {
	// init config
	_c_serverHttpAddr = config.Load("config").GetString("server.http.addr")
	_c_serverSocketAddr = config.Load("config").GetString("server.socket.addr")
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
	case "server":
		{

			fmt.Println("> base http server")
			fmt.Println("> listening on", _c_serverHttpAddr, "...")
			http.HandleFunc("/", httpserver.HttpIndex)
			http.HandleFunc("/hello", httpserver.HttpHello)
			err := http.ListenAndServe(_c_serverHttpAddr, nil)
			if err != nil {
				logger.Error("base server err:", err)
			}
		}
	case "so_server":
		{
			fmt.Println("> base socket server")
			fmt.Println("> bind on " + _c_serverSocketAddr)
			socketserver.Server(_c_serverSocketAddr)
		}
	case "so_client":
		{
			socketclient.Client(_c_serverSocketAddr)
		}
	case "gdo_query":
		{
			gdoclient.MysqlQueryBasic()
		}
	case "gdo_insert":
		{
			gdoclient.MysqlInsertBasic()
		}
	case "gdo_update":
		{
			gdoclient.MysqlUpdateBasic()
		}
	case "gdo_delete":
		{
			gdoclient.MysqlDeleteBasic()
		}
	case "gdo_tx_basic":
		{
			gdoclient.MysqlTxBasic()
		}
	case "gcache_all":
		{
			gcacheclient.TestDriver()
			gcacheclient.TestRegion()
		}
	case "gmq_pub":
		{
			gmqclient.RabbitPub()
		}
	case "gmq_sub":
		{
			gmqclient.RabbitSub()
		}
	case "etcd_ka":
		{
			etcdclient.TestKA()
		}
	case "etcd_sync":
		{
			etcdclient.TestSync()
		}
	default:
		{
			help()
		}
	}
}
