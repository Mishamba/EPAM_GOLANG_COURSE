package main

import "fmt"

func main() {
	mySlice := []string{
		"misha",
		"dog",
		"cat",
		"string",
		"aaaaaa",
	}
	fmt.Println(max(mySlice))
}

func max(givenSlice []string) (res string) {
	res = givenSlice[0]
	for _, tmp := range givenSlice {
		if len(res) < len(tmp) {
			res = tmp
		}
	}

	return res
}
