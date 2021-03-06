package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// http包建立web服务器
func sayhelloName(w http.ResponseWriter, r * http.Request) {
	r.ParseForm()	// 解析参数，默认是不会解析的
	fmt.Println("Form: ", r.Form)	// 这些信息是输出到服务器端的打印信息
	fmt.Println("path: ", r.URL.Path)
	fmt.Println("scheme: ", r.URL.Scheme)
	fmt.Println("url_long: ", r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key: ", k)
		fmt.Println("val: ", strings.Join(v, ""))
	}
	fmt.Fprintln(w, "Hello junxi")
}

func main()  {
	http.HandleFunc("/", sayhelloName)	// 设置访问的路由
	err := http.ListenAndServe(":9090", nil)	// 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

