package server

import (
	"fmt"
	"net/http"
	"time"
)

func postMethod(w http.ResponseWriter, r *http.Request) {
	method := r.Method
	if method != "POST" {
		fmt.Fprintln(w, "only post.")
	} else {
		fmt.Fprintln(w, "success.")
	}
}

func defult(w http.ResponseWriter, r *http.Request) {
	header := r.Header
	fmt.Println(header.Get("User-Agent"))
	fmt.Fprintln(w, time.Now())
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello!")
}

func startUp() {
	http.HandleFunc("/hello", sayHello) // 注册路由 当访问/ 就执行sayHello方法
	http.HandleFunc("/post", postMethod)
	http.HandleFunc("/", defult)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		fmt.Printf("http server failed, err:%v\n", err)
		return
	}
}
