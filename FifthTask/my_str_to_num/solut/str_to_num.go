package main

import (
	"fmt"
	"strconv"
)

func myStrToNum1(givenString string) (res int, err error) {
	res, err = strconv.Atoi(givenString)
	return
}

func myStrToNum2(givenString string) (res int, err error) {
	_, err = fmt.Sscanf(givenString, "%d", &res)
	return
}
