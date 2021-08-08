package main

import "fmt"

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

func mergeTwoString(string1, string2 string) {
	var result []string
	for i := 0; i < len(string1) || i < len(string1); i++ {
		if i < len(string1) {
			result = append(result, string(string1[i]))
		}

		if i < len(string2) {
			result = append(result, string(string2[i]))
		}
	}
	var word string
	for _, val := range result {
		word += val
	}
	fmt.Println(word)
}

func jackDaniel(a string, b string) {
	t := len(a) + len(b)
	i, j := 0, 0
	ret := make([]byte, t)
	for i+j < t && i < len(a) && j < len(b) {
		if b[j] < a[i] {
			ret[i+j] = b[j]
			j += 1
			continue
		}
		ret[i+j] = a[i]
		i += 1
	}
	ret = []byte(fmt.Sprintf("%s%s%s", ret, a[i:], b[j:]))
	fmt.Println(string(ret))
	//return string(ret)
}

func minMax(array []int) (int, int) {
	var max int = array[0]
	var min int = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return min, max
}

func mOperation(input [][]int) {
	initial := []int{0, 0, 0, 0, 0}
	for i := 0; i <= len(input)-1; i++ {
		for j := input[i][0] - 1; j <= input[i][1]-1; j++ {
			initial[j] += input[i][2]
		}
	}
	fmt.Println(initial)
	_, max := minMax(initial)
	fmt.Println(max)
}

func main() {
	fmt.Println(".:Start:.")
	fmt.Print("\n")

	fmt.Println("No.1")
	swapVariable(3, 5)
	fmt.Print("\n")

	fmt.Println("No.2")
	howManyTimeAppears([]int{4, 6, 3, 5, 4, 6, 7, 8, 3, 4, 6, 7, 5, 4, 6, 4, 4, 5, 6})
	fmt.Print("\n")

	fmt.Println("No.3")
	mergeTwoString("saya", "kamu")
	mergeTwoString("kosong", "ada")
	mergeTwoString("ada", "kosong")
	fmt.Print("\n")

	fmt.Println("No.4")
	jackDaniel("JACK", "DANIEL")
	jackDaniel("ABACABA", "ABACABA")
	fmt.Print("\n")

	fmt.Println("No.5")
	var inputListNumber = [][]int{
		{1, 2, 100},
		{2, 5, 100},
		{3, 4, 100},
	}
	mOperation(inputListNumber)

	fmt.Print("\n")
	fmt.Println(".:Finished:.")
}
