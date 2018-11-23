package main

import (
	"fmt"

	resty "gopkg.in/resty.v0"
)

func main() {
	r := resty.R()
	r.SetHeaders(map[string]string{
		"Cache-Control": "max-age=0",
	})
	res, err := r.Get("http://127.0.0.1:2234/hello/hello.html")
	if err != nil {
		panic(err)
	}
	fmt.Println(res.StatusCode())
}
