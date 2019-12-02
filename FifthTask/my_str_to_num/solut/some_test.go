package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test1(t *testing.T) {
	a := "15"
	receivedNumber, err := myStrToNum1(a)
	assert.Equal(t, 15, receivedNumber, "test1, equal numbers")
	assert.Nil(t, err, "test 2, nil check")

	a = "5g"
	receivedNumber, err = myStrToNum1(a)
	assert.NotNil(t, err, "test3, nil check")
}

func Test2(t *testing.T) {
	a := "15"
	receivedNumber, err := myStrToNum2(a)
	assert.Equal(t, 15, receivedNumber, "test1, equal numbers")
	assert.Nil(t, err, "test 2, nil check")

	a = "g"
	receivedNumber, err = myStrToNum2(a)
	fmt.Println(receivedNumber)
	assert.NotNil(t, err, "test3, nil check")
}
