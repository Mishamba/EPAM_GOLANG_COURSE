//	This package can translate your number, given as a string, to NUMBER(int, float, ...)
package main

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"unicode/utf8"
)

func myStrToNum(givenString string, expectedType string) (result float64, err error) {
	numberType := map[string]bool{}

	wrongDataCheck(numberType, givenString)
	binaryCheck(numberType, givenString)
	intCheck(numberType, givenString)
	var dotPosition int
	floatCheck(numberType, givenString, &dotPosition)
	hexadecimalCheck(numberType, givenString)

	myType, err := typeDefine(numberType, expectedType)

	fmt.Println("your type is", myType)

	switch myType {
	case "binary":
		{
			return binaryConverter(givenString)
		}

	case "int":
		{
			return intConverter(givenString)
		}

	case "float":
		{
			return floatConverter(givenString, dotPosition)
		}

	case "hexadecimal":
		{
			return hexadecimalConverter(givenString)
		}

	default:
		return 0.0, errors.New("u gave wrong data")
	}
}

func typeDefine(numberType map[string]bool, expectedType string) (string, error) {
	if numberType["wrongData"] {
		return "wrongData", errors.New("we received wrong data in typeDefine")
	}
	var typeCount int
	var trueState string

	for key, state := range numberType {
		if state {
			typeCount++
			trueState = key
		}
	}

	if typeCount == 0 {
		return "", errors.New("no TRUE types defined")
	}

	if typeCount > 1 {
		if numberType["float"] {
			return "", errors.New("received float with another types")
		}
		/*
			if _, err := fmt.Println("we not sure which number type it is. here u can see our ideas, check one u like the most"); err != nil { //in this comment stays code, which works, but can't pass tests
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

			if numberType[usersChoose] {
				return usersChoose, nil
			} else {
				return "", errors.New("u made a mistake, during entering one of the variants")
			}*/

		if numberType[expectedType] {
			return expectedType, nil
		} else {
			return "", errors.New("didn't recognized expectedType")
		}
	}

	return trueState, nil
}

func binaryConverter(givenString string) (result float64, err error) {
	for i := utf8.RuneCountInString(givenString) - 1; i > -1; i-- {
		tmpString := givenString[i : i+1] //looks like bad idea. if it can become better, so report
		if tmp, err := strconv.Atoi(string(tmpString)); err != nil {
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
	if string(givenString[dotPosition]) != "." {
		return 0.0, errors.New("on dot position received smth difference")
	}
	fractPart := givenString[dotPosition+1:]
	var intP int
	if intP, err = strconv.Atoi(intPart); err != nil {
		return 0.0, err
	}
	var fractP int
	if fractP, err = strconv.Atoi(fractPart); err != nil {
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
		if !((givenString[i] >= 48 && givenString[i] <= 57) || (givenString[i] >= 65 && givenString[i] <= 70)) {
			return 0.0, errors.New("in hexadecimalconverter was given not hexidecimal number")
		}
		tmp := lettersCost[string(tmpString[0])]
		result += float64(tmp) * math.Pow(16.0, float64(utf8.RuneCountInString(givenString)-i-1))
	}
	fmt.Println()
	return result, nil

}

func wrongDataCheck(result map[string]bool, given string) {
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

	if (dotCheck && letterCheck) || (dotPosition == 0 && dotCheck) { //if in string we have dots and letter together, so it's a mistake
		result["wrongData"] = true
		return
	}

	result["wrongData"] = false
}

func binaryCheck(result map[string]bool, given string) {
	for i := 0; i < utf8.RuneCountInString(given); i++ {
		if !(given[i] == 48 || given[i] == 49) {
			result["binary"] = false
			return
		}
	}

	result["binary"] = true
}

func intCheck(result map[string]bool, given string) {
	for i := 0; i < utf8.RuneCountInString(given); i++ {
		if !(given[i] >= 48 && given[i] <= 57) {
			result["int"] = false
			return
		}
	}

	result["int"] = true
}

func floatCheck(result map[string]bool, given string, dotPosition *int) {
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
	for i := 0; i < utf8.RuneCountInString(given); i++ {
		if !((given[i] >= 48 && given[i] <= 57) || (given[i] >= 65 && given[i] <= 70)) {
			result["hexadecimal"] = false
			return
		}
	}

	result["hexadecimal"] = true
}
