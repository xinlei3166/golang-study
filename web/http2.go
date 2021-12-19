package main

import (
	"fmt"
	"net/http"
)

// 自定义路由器

type MyMux struct {

}

func (p *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sayHelloName(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello myroute!")
}

func main() {
	mux := &MyMux{}
	http.ListenAndServe(":9090", mux)
}
