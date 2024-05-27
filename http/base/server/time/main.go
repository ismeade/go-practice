package main

import (
	"fmt"
	"net/http"
	"time"
)

func defult(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, time.Now().Unix())
}

func main() {
	http.HandleFunc("/", defult)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("http server failed, err:%v\n", err)
		return
	}
}
