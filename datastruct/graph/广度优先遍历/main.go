package main

import (
	"fmt"

	"github.com/daysleep666/someproject/datastruct/graph/graph"
)

func BFS() {

}

func main() {
	dg := graph.Init()
	dg.Add("a", "b")
	dg.Add("a", "c")
	dg.Add("b", "c")
	var list []string
	var tmpList []string
	list = append(list, "a")

	for {
		if len(list) == 0 {
			break
		}
		for _, v := range list {
			tmpList = append(tmpList, dg.GetNodes(v)...)
			fmt.Printf("%v->", v)
		}
		list = tmpList
		tmpList = tmpList[:0]
	}
}
