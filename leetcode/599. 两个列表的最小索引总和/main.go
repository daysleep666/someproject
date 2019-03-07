package main

import "fmt"

func findRestaurant(list1 []string, list2 []string) []string {
	result := make([]string, 0)
	m := make(map[string]int)
	min := 99999
	for i, s1 := range list1 {
		m[s1] = i
	}
	for j, s2 := range list2 {
		if i, isExist := m[s2]; isExist {
			v := i + j
			fmt.Println(v, min)
			if v == min {
				result = append(result, s2)
			} else if v < min {
				result = []string{s2}
				min = v
			}
		}
	}
	return result
}

func main() {
	fmt.Println(findRestaurant([]string{"Shogun", "Tapioca Express", "Burger King", "KFC"}, []string{"KFC", "Shogun", "Burger King"}))
}
