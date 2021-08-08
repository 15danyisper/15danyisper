package main

import (
	"fmt"
)

var listNumber = []int{4, 6, 3, 5, 4, 6, 7, 8, 3, 4, 6, 7, 5, 4, 6, 4, 4, 5, 6}

type countMap struct {
	num   int
	count int
}

func swapVariable(a, b int) {
	a, b = b, a
	fmt.Println(a, b)
}

func unique(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}

func howManyTimeAppears(n []int) {
	isUnique := unique(n)
	listMapUnique := []map[int]int{}

	for _, val := range isUnique {
		mapUnique := map[int]int{}
		mapUnique[val] = 0
		listMapUnique = append(listMapUnique, mapUnique)
	}

	for i := 0; i < len(isUnique); i++ {
		var counter int
		for _, val := range n {
			if isUnique[i] == val {
				counter++
				listMapUnique[i][isUnique[i]] = counter
			}
		}
	}

	for i := 0; i < len(listMapUnique); i++ {
		fmt.Println(isUnique[i], "=>", listMapUnique[i][isUnique[i]])
	}

}

func main() {
	fmt.Println("Hello world...")
	swapVariable(3, 5)
	howManyTimeAppears(listNumber)
}
