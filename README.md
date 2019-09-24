# go-base

A simple go framework for development

> This framework could be copy to create a new project as a template. And you can also copy the "src/base" library to use it in your project independently.

## Installing

**Set system env**

```
vi ~/.bash_profile
...
export GOPATH=/path/to/go-base:...
export GO_BASE_ENV=[local|test|...|release]
export GO_BASE_ROOT=/path/to/go-base
...
```

**Run basic examples**

```
git clone http://192.168.1.201:8888/go/go-base.git
cd go-base
make
bin/go-base
...
============================================
> rootpath /data/code/go-base
============================================
 base server : http server
 base so_server : socket server
 base so_client : socket client
 base gdo_query : gdo query test
 base gdo_insert : gdo insert test
 base gdo_update : gdo update test
 base gdo_delete : gdo delete test
 base gdo_tx_basic : gdo transaction test
 base gcache_all : gcache all tests
 base gmq_pub : gmq publish test
 base gmq_sub : gmq comsume test
 base etcd_ka : etcd keepalive test
 base etcd_sync : etcd sync trans test
============================================
...
bin/go-base server
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
src/base/etcd|etcd support package
src/base/proto|protobuf base package
src/base/example|examples package
vendor/|third-party packages, used only in make process


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
dump> "util.GetEnv()" "local" 
dump> "util.GetRootPath()" "/path/to/project-name"
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

**src/base/example/gdoclient/base.go**

> Including db CRUD function examples, now support mysql, keep upgrading.

**src/base/example/gmqclient/base.go**

> Including message queue basic examples, now support rabbitmq, keep upgrading.

**src/base/example/gcacheclient/base.go**

> Including cache basic examples, now support redis, keep upgrading.

**src/base/example/etcdclient/base.go**

> Including etcd basic examples, keep upgrading.

**src/base/example/socketserver/base.go**

> A simple socket server demo, including heartbeating implemention and how to use logger library.

**src/base/example/socketclient/base.go**

> A simple socket client demo, including heartbeating implemention.

**src/base/example/httpserver.go**

> A simple http server demo.
