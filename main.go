package main

import (
	"fmt"

	"github.com/daysleep666/someproject/datastruct/queue/queuesturct"
)

func main() {
	q := queuesturct.NewQueue()
	for i := 0; i < 10; i++ {
		q.Push(i)
	}
	for v, err := q.Pop(); err == nil; v, err = q.Pop() {
		fmt.Println(v.(int))
	}
}
