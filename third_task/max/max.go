package main

import (
	"fmt"
	"unicode/utf8"
)

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
	if len(givenSlice) == 0 {
		return ""
	}

	res = givenSlice[0]
	for _, tmp := range givenSlice {
		if utf8.RuneCountInString(res) < utf8.RuneCountInString(tmp) {
			res = tmp
		}
	}

	return res
}
