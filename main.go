package main

import (
	"fmt"
	"net/http"
	"os"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("hello world")
}

func main() {
	http.HandleFunc("/heartbeat", IndexHandler)
	ip := "0.0.0.0:" + os.Args[1]
	fmt.Println("listen in", ip)
	http.ListenAndServe(ip, nil)
}
