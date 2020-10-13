# go-base

A simple go framework for development

> This framework could be easliy copy & create a new project as a project template. And you can also sync the "src/base" libraries to another project by using build-in command line tools.

## Features

* Full Stack. Use Makefile as development & build command line tools, across platform support.

* NOT only for WEB but for ALL. Proper framework offers the advantage of being extremely scalable.

* Operational Friendly. Could be simply configured in different environments, such as local, test, release ...

* Rich Components Support. Provides GDO (GO Database OOP Libs), GMO (Go MQ OOP Libs), GCACHE (Go Cache OOP Libs), and also support Etcd, Protobuf ...

* Simple to Start. EASY to use, EASY to expand Libs. EASY to copy & create new project. EASY to sync codes between projects.

## Installing

**Set system ENV (Linux & Mac)**

```
vi ~/.bash_profile
...
export GO_BASE_ENV=[local|test|...|release]
export GO_BASE_ROOT=/path/to/go-base
...
```

**Run basic examples**

```
git clone http://192.168.1.201:8888/go/go-base.git
cd go-base
go mod download
make # Build in Linux & Mac
build.bat # Build in Windows
bin/go-base
...
============================================
> rootpath /data/code/go-base
============================================
 base http_server : example http server
 base so_server : example socket server
 base so_client : example socket client
 base gdo_query : example gdo query test
 base gdo_insert : example gdo insert test
 base gdo_update : example gdo update test
 base gdo_delete : example gdo delete test
 base gdo_tx_basic : example gdo transaction test
 base gcache_all : example gcache all tests
 base gmq_pub : example gmq publish test
 base gmq_sub : example gmq comsume test
 base etcd_ka : example etcd keepalive test
 base etcd_sync : example etcd sync trans test
============================================
...
bin/go-base http_server
...
> base http server
> listening on 0.0.0.0:80 ...
...
bin/go-base so_server
...
> base socket server
> bind on 0.0.0.0:8080
...
bin/go-base so_client
...
> connect success!
> send: 29 text: 1 Hello I'm heartbeat client.
...
```

## Development

**Path introduction**

* introductions of go-base framework's paths are as follow

Path|Introduction
--|--
bin/|run path
log/|log files
etc/local/|default config files
etc/{BASE_ENV}/|config files in GO_BASE_ENV enviornment
src/base/config|config package
src/base/logger|logger package
src/base/gcache|cache package, support : redis ...
src/base/gmq|mq package, support : rabbitmq ...
src/base/gdo|db package, support : mysql ...
src/base/getcd|etcd support package
src/base/gutil|utils functions
src/base/proto|protobuf base package
src/base/example|examples package


**How to start new project**

* use "make create" command to create a new project as follow

```
make create path=/path/to/project-name
...
make> Create project 'project-name' at /path/to/project-name ...
make> Change project base ENV to PROJECT_NAME_ENV ...
...
cd /path/to/project-name
make
...
make> Prepare dirs ...
make> Cleaning build cache ...
make> Checking missing dependencies ...
make> Building binary files ...
...
bin/project-name
...
dump> "gutil.GetEnv()" "local" 
dump> "gutil.GetRootPath()" "/path/to/project-name"
...
```

**How to use base library**

* use "make synclib" command to sync libraries between projects

```
make synclib path=/path/to/project-name
...
make> Sync all libraries to /path/to/project-name ...
make> Change project base ENV to PROJECT_NAME_ENV ...
...
cd /path/to/project-name
git status
...
```

## Examples

**main.go**

> Main examples console, including the example how to use config library.

**lib/example/gdoclient/example_gdoclient.go**

> Including db CRUD function examples, now support mysql, keep upgrading.

**lib/example/gmqclient/example_gmqclient.go**

> Including message queue basic examples, now support rabbitmq, keep upgrading.

**lib/example/gcacheclient/example_gcacheclient.go**

> Including cache basic examples, now support redis, keep upgrading.

**lib/example/etcdclient/example_etcdclient.go**

> Including etcd basic examples, keep upgrading.

**lib/example/socketserver/example_socketserver.go**

> A simple socket server demo, including heartbeating implemention and how to use logger library.

**lib/example/socketclient/example_socketclient.go**

> A simple socket client demo, including heartbeating implemention.

**lib/example/httpserver/example_httpserver.go**

> A simple http server demo.
