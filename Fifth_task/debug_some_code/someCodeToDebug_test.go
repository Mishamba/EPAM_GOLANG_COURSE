package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMulriplyByTwo(t *testing.T) {
	var number int = 5
	err := multiplyByTwo(&number)
	assert.Equal(t, 10, number, "test 1. value check")
	assert.Nil(t, err, "test 1. err check")

	number = 0
	err = multiplyByTwo(&number)
	assert.Equal(t, 0, number, "test2. value check")
	assert.NotNil(t, err, "test 2. err check")
}

func TestPrintMoreTen(t *testing.T) {
	var number int64 = 5
	err := printMoreTen(number)
	assert.NotNil(t, err, "test 3. err check")

	number = 10
	err = printMoreTen(number)
	assert.Nil(t, err, "test 4. err check")

	number = 15
	err = printMoreTen(number)
	assert.Nil(t, err, "test 5. err check")
}

func TestDejson(t *testing.T) {
	var smth interface{}
	smth = &jsStruct{}
	err := dejson(smth)
	assert.Nil(t, err, "test 6. err check")

	smth = 5
	err = dejson(smth)
	assert.NotNil(t, err, "test 7. err check")
}
