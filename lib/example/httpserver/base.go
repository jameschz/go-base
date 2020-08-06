package httpserver

import (
	"go-base/lib/config"
	"fmt"
	"io"
	"net/http"
)

func HttpIndex(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "<a href='/hello'>hello</a>")
}

func HttpHello(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, We can load configs:\n")
	io.WriteString(w, fmt.Sprintf("%#v", config.Load("config").GetStringMap("server")))

}
