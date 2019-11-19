package main

import "fmt"

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
	res := []string{}
	keys := []int{}
	for k := range givenMap {
		keys = append(keys, k)
	}

	fmt.Print(keys)

	for i := 0; i < len(keys); i++ {
		for j := i + 1; j < len(keys); j++ {
			if keys[i] > keys[j] {
				keys[i], keys[j] = keys[j], keys[i]
			}
		}
	}

	for k := range keys {
		res = append(res, givenMap[k]) //не записывает последнее значение.
	}

	fmt.Println(res)
}
