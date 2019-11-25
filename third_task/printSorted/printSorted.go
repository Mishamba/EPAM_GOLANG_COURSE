package main

import (
	"fmt"
	"sort"
)

func main() {
	myMap := map[int]string{
		5: "think",
		2: "is",
		3: "very",
		1: "Go",
		4: "strange",
	}
	printSorted(myMap)
}

func printSorted(givenMap map[int]string) {
	keys := []int{}
	for k := range givenMap {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	for _, v := range keys {
		fmt.Println(givenMap[v])
	}
}
