package main

import (
	"fmt"

	"github.com/daysleep666/someproject/datastruct/graph/graph"
)

func BFS() {

}

func main() {
	dg := graph.Init()
	dg.Add("0", "1")
	dg.Add("0", "3")
	dg.Add("1", "2")
	dg.Add("1", "4")
	dg.Add("3", "4")
	dg.Add("4", "5")
	dg.Add("4", "6")
	dg.Add("6", "7")
	var list []string
	var tmpList []string
	list = append(list, "0")
	var visisted = make(map[string]struct{})
	// var a struct{}
	for {
		if len(list) == 0 {
			break
		}
		for i := 0; i < len(list); i++ {
			tmpList = append(tmpList, dg.GetNodes(list[i])...)
			fmt.Println(list, len(list), i, list[i])
			// if list[i] == "3" {
			// 	fmt.Println("-------", list)
			// }
			// if _, isExist := visisted[v]; !isExist {
			// 	// fmt.Printf("%v->", v)
			// 	visisted[v] = a
			// }
		}
		// fmt.Printf("\ntmplist:%v....\n", tmpList)
		// copy(list, tmpList)
		list = tmpList[:]
		tmpList = tmpList[:0]
	}
	fmt.Println(visisted)
}
