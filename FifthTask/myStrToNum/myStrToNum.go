//	This package can translate your number, given as a string, to NUMBER(int, float, ...)
//	This task is complited, but i need some help. u can see, that i tried to do int with goroutined. Now routines are really helpful. and i need your help. show, how to start it there please.
package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"sync"
	"unicode/utf8"
)

var wg sync.WaitGroup

func main() {
	a, err := myStrToInt("10") //set value here(expected 1000000000)
	if err != nil {
		panic(err)
	}
	fmt.Println(a)
}

func myStrToInt(givenString string) (result float64, err error) {
	numberType := map[string]bool{}

	wg.Add(5)

	go wrongDataCheck(numberType, givenString)
	go binaryCheck(numberType, givenString)
	go intCheck(numberType, givenString)
	var dotPosition int
	go floatCheck(numberType, givenString, &dotPosition)
	go hexadecimalCheck(numberType, givenString)

	wg.Wait()

	var myType string
	if myType, err = typeDefine(numberType); err != nil {
		return 0.0, err
	}

	if myType == "wrongData" {
		return 0.0, errors.New("u gave wrong data")
	}

	fmt.Println("your type is", myType)

	if myType == "binary" {
		return binaryConverter(givenString)
	}

	if myType == "int" {
		return intConverter(givenString)
	}

	if myType == "float" {
		return floatConverter(givenString, dotPosition)
	}

	if myType == "hexadecimal" {
		return hexadecimalConverter(givenString)
	}

	return 0.0, errors.New("smth went wrong. u need to debug this code...")
}

func typeDefine(numberType map[string]bool) (string, error) {
	var count int
	var trueState string

	for key, state := range numberType {
		if state {
			count++
			trueState = key
		}
	}

	if count > 1 {
		if _, err := fmt.Println("we not sure which number type it is. here u can see our ideas, check one u like the most"); err != nil {
			return "", err
		}

		fmt.Println()

		for i, v := range numberType {
			if v {
				fmt.Println(i)
			}
		}

		fmt.Println("Choose one (u need to enter one variant as a string)")

		var usersChoose string
		if _, err := fmt.Scan(&usersChoose); err != nil {
			return "", err
		}

		return usersChoose, nil

	}

	return trueState, nil
}

func binaryConverter(givenString string) (result float64, err error) {
	for i := utf8.RuneCountInString(givenString) - 1; i > -1; i-- {
		tmpString := givenString[i : i+1] //looks like bad idea. if it can become better, so report
		if tmp, err := strconv.Atoi(tmpString); err != nil {
			return 0.0, err
		} else {
			result += float64(tmp) * math.Pow(2.0, float64(utf8.RuneCountInString(givenString)-i-1))
		}
	}
	fmt.Println()
	return result, nil
}

func intConverter(givenString string) (result float64, err error) {
	if result, err := strconv.Atoi(givenString); err != nil {
		return 0.0, err
	} else {
		return float64(result), nil
	}
}

func floatConverter(givenString string, dotPosition int) (result float64, err error) {
	intPart := givenString[:dotPosition]
	fractPart := givenString[dotPosition+1:]
	intP, err := strconv.Atoi(intPart)
	if err != nil {
		return 0.0, err
	}
	fractP, err := strconv.Atoi(fractPart)
	if err != nil {
		return 0.0, err
	}
	result = float64(intP) + float64(fractP)/math.Pow(10.0, float64(utf8.RuneCountInString(fractPart)))

	return result, nil

}

func hexadecimalConverter(givenString string) (result float64, err error) {
	lettersCost := map[string]int{
		"0": 0, //if i can use here iota, please show how, because i couldn't find any exaple of using it it maps
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
		"A": 10,
		"B": 11,
		"C": 12,
		"D": 13,
		"E": 14,
		"F": 15,
	}
	for i := utf8.RuneCountInString(givenString) - 1; i > -1; i-- {
		tmpString := givenString[i : i+1]
		tmp := lettersCost[string(tmpString[0])]
		result += float64(tmp) * math.Pow(16.0, float64(utf8.RuneCountInString(givenString)-i-1))
	}
	fmt.Println()
	return result, nil

}

func wrongDataCheck(result map[string]bool, given string) {
	defer wg.Done()
	var dotCheck bool
	var letterCheck bool
	var dotPosition int
	for i := 0; i < utf8.RuneCountInString(given); i++ {
		if given[i] < 46 || given[i] == 47 || (given[i] >= 58 && given[i] <= 64) || given[i] > 70 {
			result["wrongData"] = true
			return
		}

		if given[i] == 46 {
			dotPosition = i
			dotCheck = true
		}
		if given[i] >= 65 && given[i] <= 70 {
			letterCheck = true
		}
	}

	if dotCheck && letterCheck { //if in string we have dots and letter together, so it's a mistake
		result["wrongData"] = true
		return
	}

	if dotPosition == 0 && dotCheck {
		result["wrongData"] = true
		return
	}

	result["wrongData"] = false
}

func binaryCheck(result map[string]bool, given string) {
	defer wg.Done()
	for i := 0; i < utf8.RuneCountInString(given); i++ {
		if !(given[i] == 48 || given[i] == 49) {
			result["binary"] = false
			return
		}
	}

	result["binary"] = true
}

func intCheck(result map[string]bool, given string) {
	defer wg.Done()
	for i := 0; i < utf8.RuneCountInString(given); i++ {
		if !(given[i] >= 48 && given[i] <= 57) {
			result["int"] = false
			return
		}
	}

	result["int"] = true
}

func floatCheck(result map[string]bool, given string, dotPosition *int) {
	defer wg.Done()
	for i := 0; i < utf8.RuneCountInString(given); i++ {
		if !((given[i] >= 48 && given[i] <= 57) || given[i] == 46) {
			result["float"] = false
			return
		}
		if given[i] == 46 {
			*dotPosition = i
		}
	}

	if *dotPosition == 0 {
		result["float"] = false
		return
	}

	result["float"] = true
}

func hexadecimalCheck(result map[string]bool, given string) {
	defer wg.Done()
	for i := 0; i < utf8.RuneCountInString(given); i++ {
		if !((given[i] >= 48 && given[i] <= 57) || (given[i] >= 65 && given[i] <= 70)) {
			result["hexadecimal"] = false
			return
		}
	}

	result["hexadecimal"] = true
}
