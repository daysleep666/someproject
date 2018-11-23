package main

import "net/http"

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`hello world`))
	})
	http.ListenAndServe(":2234", nil) // <-今天讲的就是这个ListenAndServe是如何工作的
}
