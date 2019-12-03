package main

import (
	"fmt"
	"strconv"
)

func myStrToNum1(givenString string) (res int, err error) {
	return strconv.Atoi(givenString)
}

func myStrToNum2(givenString string) (res int, err error) {
	_, err = fmt.Sscanf(givenString, "%d", &res)
	return
}
