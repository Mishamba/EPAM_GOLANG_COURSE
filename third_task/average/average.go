package main

import "fmt"

func main() {
	myArray := [5]int{1, 2, 3, 4, 5}
	fmt.Println(average(myArray[:]))
	myArray2 := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(average(myArray2[:]))
}

func average(givenArray []int) float64 {
	var Res int
	for _, data := range givenArray {
		Res += data
	}

	return float64(Res) / float64(len(givenArray))
}
