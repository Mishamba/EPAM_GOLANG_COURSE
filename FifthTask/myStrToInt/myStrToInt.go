package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"unicode/utf8"
)

const (
	hexadecimalNotation = "ABCDEF"
)

var (
	varNumber   = [5]string{"hexadecimal", "binary", "float", "int", "wrongData"}
	uft8Numbers = [10]int{48, 49, 50, 51, 52, 53, 54, 55, 56, 57}
)

func main() {
	a, err := myStrToInt("5")
	if err != nil {
		panic(err)
	}
	fmt.Println(a)
}

func myStrToInt(givenString string) (result float64, err error) {
	numberType := make(chan string)
	go wrongDataCheck(numberType, givenString)
	go binaryCheck(numberType, givenString)
	go intCheck(numberType, givenString)
	var dotPosition int
	go floatCheck(numberType, givenString, dotPosition)
	go hexadecimalCheck(numberType, givenString)
	myType := <-numberType
	fmt.Println(myType)
	if myType == "wrong Data" {
		err = errors.New("u gave wrong data")
		return 0.0, err
	}
	if myType == "binaty notation" {
		for i, _ := range givenString {
			result += math.Pow(2.0, float64(i))
		}
		return result, nil
	}
	if myType == "int number" {
		res, err := strconv.Atoi(givenString)
		if err != nil {
			panic(err)
		}
		return float64(res), nil
	}
	if myType == "float number" {
		intPart := givenString[:dotPosition]
		fractPart := givenString[dotPosition+1:]
		intP, err := strconv.Atoi(intPart)
		if err != nil {
			panic(err)
		}
		fractP, err := strconv.Atoi(fractPart)
		if err != nil {
			panic(err)
		}
		result = float64(intP) + float64(fractP)/math.Pow(10.0, float64(utf8.RuneCountInString(fractPart)))
	}
	if myType == "hexadecimal notation" {
		//тута надо будет такое писать, что офигеть можно
		fmt.Println("i've been here...")
	}

	err = errors.New("smth went wrong. u need to debug this code...")
	return 0.0, err
}

func wrongDataCheck(result chan string, given string) {
	for _, v := range given {
		if v < 46 || v == 47 || (v > 57 && v < 65) || v > 70 {
			fmt.Println("wr")
			result <- "wrong Data"
		}
	}
}

func binaryCheck(result chan string, given string) {
	key := true
	for _, v := range given {
		if !(v == 48 || v == 49) {
			key = false
		}
	}
	if key {
		fmt.Println("b")
		result <- "binary notation"
	}

}

func intCheck(result chan string, given string) {
	key := true
	for _, v := range given {
		if !(v >= 47 && v <= 57) {
			key = false
		}
	}
	if key {
		fmt.Println("int")
		result <- "int number"
	}
}

func floatCheck(result chan string, given string, dotPosition int) {
	key := true
	for i, v := range given {
		if !((v >= 47 && v <= 57) || v != 46) {
			key = false
		}
		if v == 46 {
			dotPosition = i
		}
	}
	if key {
		fmt.Println("float")
		result <- "float number"
	}
}

func hexadecimalCheck(result chan string, given string) {
	key := true
	for _, v := range given {
		if !((v >= 47 && v <= 57) || (v >= 65 && v <= 70)) {
			key = false
		}
	}
	if key {
		fmt.Println("float")
		result <- "hexadecimal notation"
	}
}
