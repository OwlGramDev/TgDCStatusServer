package webserver

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"log"
)

func (g Context) Run() {
	fmt.Println("Running WebServer!")
	log.Fatal(fasthttp.ListenAndServe(fmt.Sprintf("127.0.0.1:%d", 8143), g.handler))
}
