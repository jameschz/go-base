package examplehttpserver

import (
	"fmt"
	"github.com/jameschz/go-base/lib/config"
	"io"
	"net/http"
)

// HTTPIndex :
func HTTPIndex(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "<a href='/hello'>hello</a>")
}

// HTTPHello :
func HTTPHello(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, We can load configs:\n")
	io.WriteString(w, fmt.Sprintf("%#v", config.Load("config").GetStringMap("server")))

}
