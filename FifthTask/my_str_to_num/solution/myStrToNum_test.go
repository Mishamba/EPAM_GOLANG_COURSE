package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrongData(t *testing.T) {
	someString := ".bdA562"
	receiveType := map[string]bool{
		"wrongData": false,
	}
	wrongDataCheck(receiveType, someString)
	assert.True(t, receiveType["wrongData"], "expected \"TRUE\" but received \"nothing\"")

	someString = "B.dA562"
	receiveType = map[string]bool{
		"wrongData": false,
	}
	wrongDataCheck(receiveType, someString)
	assert.True(t, receiveType["wrongData"], "expected \"TRUE\" but received \"nothing\"")

	someString = "B.562"
	receiveType = map[string]bool{
		"wrongData": false,
	}
	wrongDataCheck(receiveType, someString)
	assert.True(t, receiveType["wrongData"], "expected \"TRUE\" but received \"nothing\"")

	someString = "14"
	receiveType["wrongData"] = false
	wrongDataCheck(receiveType, someString)
	assert.True(t, !receiveType["wrongData"], "expected \"FALSE\", but received \"TRUE\"")

	someString = ".463"
	receiveType["wrongData"] = true
	wrongDataCheck(receiveType, someString)
	assert.True(t, receiveType["wrongData"], "expected \"FALSE\", but received \"TRUE\" 1")
}

func TestBinaryCheck(t *testing.T) {
	someString := "111001"
	receiveType := map[string]bool{
		"binary": false,
	}
	binaryCheck(receiveType, someString)
	assert.True(t, receiveType["binary"], "expected \"TRUE\" but received \"nothing\"")

	someString = "101f0201"
	binaryCheck(receiveType, someString)
	assert.True(t, !receiveType["binary"], "expected \"FALSE\", but received \"TRUE\"")
}

func TestIntCheck(t *testing.T) {
	someString := "1583"
	receiveType := map[string]bool{
		"int": false,
	}
	intCheck(receiveType, someString)
	assert.True(t, receiveType["int"], "expected \"TRUE\" but received \"nothing\"")

	someString = "6.32gs9"
	receiveType["int"] = true
	intCheck(receiveType, someString)
	assert.True(t, !receiveType["int"], "expected \"FALSE\", but received \"TRUE\"")
}

func TestFloatCheck(t *testing.T) {
	someString := "1286.8643"
	receiveType := map[string]bool{
		"float": false,
	}
	var receiveDotPosition int
	expectedDotPosition := 4
	floatCheck(receiveType, someString, &receiveDotPosition)
	assert.True(t, receiveType["float"], "expected \"TRUE\" but received \"nothing\"")
	assert.Equal(t, expectedDotPosition, receiveDotPosition, "dot expected at \"4\" position, but received another")

	someString = "1.g463A"
	receiveType["float"] = true
	floatCheck(receiveType, someString, &receiveDotPosition)
	assert.True(t, !receiveType["float"], "expected \"FALSE\", but received \"TRUE\"")

	someString = ".463"
	receiveType["float"] = true
	floatCheck(receiveType, someString, &receiveDotPosition)
	assert.True(t, !receiveType["float"], "expected \"FALSE\", but received \"TRUE\"")

}

func TestHexadecimalCheck(t *testing.T) {
	someString := "3A632F"
	receiveType := map[string]bool{
		"hexadecimal": false,
	}
	hexadecimalCheck(receiveType, someString)
	assert.True(t, receiveType["hexadecimal"], "expected \"TRUE\" but received \"nothing\"")

	someString = "53Fb32uA"
	receiveType["hexadecimal"] = true
	hexadecimalCheck(receiveType, someString)
	assert.True(t, !receiveType["hexadecimal"], "expected \"FALSE\", but received \"TRUE\"")
}

func TestBinaryConverteer(t *testing.T) {
	someString := "11111010110"
	receivedNumber, err := binaryConverter(someString)
	expectedNumber := 2006.0
	assert.Equal(t, expectedNumber, receivedNumber, "expected \"2006\", but received smth different")
	assert.Nil(t, err, "expected \"nil\", but received \", err, \"")

	someString = "1003b01"
	receivedNumber, err = binaryConverter(someString)
	expectedNumber = 0.0
	assert.Equal(t, expectedNumber, receivedNumber, "expected \"0.0\", but received smth different")
	assert.NotNil(t, err, "expected \"smth\", but received\"nil\"")
}

func TestIntConverter(t *testing.T) {
	someString := "2006"
	receivedNumber, err := intConverter(someString)
	expectedNumber := 2006.0
	assert.Equal(t, expectedNumber, receivedNumber, "expected \"2006\", but received smth different")
	assert.Nil(t, err, "expected \"nil\", but received smth different")

	someString = "200g5"
	receivedNumber, err = intConverter(someString)
	expectedNumber = 0.0
	assert.Equal(t, expectedNumber, receivedNumber, "expected \"0.0\", but received smth different")
	assert.NotNil(t, err, "expected \"smth\", but received \"nil\"")
}

func TestFloatConverter(t *testing.T) { //finish IT
	someString := "6332.853"
	dotPosition := 4
	receivedNumber, err := floatConverter(someString, dotPosition)
	expectedNumber := 6332.853
	assert.Equal(t, expectedNumber, receivedNumber, "expected \"6332.853\", but received smth different")
	assert.Nil(t, err, "expected \"nil\", but received smth different")

	someString = "63f2.853"
	dotPosition = 4
	receivedNumber, err = floatConverter(someString, dotPosition)
	expectedNumber = 0.0
	assert.Equal(t, expectedNumber, receivedNumber, "expected \"0.0\", but received smth different")
	assert.NotNil(t, err, "expected \"smth\", but received smth \"nil\"")

	someString = "63.32853"
	dotPosition = 4
	receivedNumber, err = floatConverter(someString, dotPosition)
	expectedNumber = 0.0
	assert.Equal(t, expectedNumber, receivedNumber, "expected \"0.0\", but received smth different")
	assert.NotNil(t, err, "expected \"smth\", but received \"nil\"")

	someString = "!@!@!@!@!"
	dotPosition = 4
	receivedNumber, err = floatConverter(someString, dotPosition)
	expectedNumber = 0.0
	assert.Equal(t, expectedNumber, receivedNumber, "expected \"0.0\", but received smth different")
	assert.NotNil(t, err, "expected \"smth\", but received smth \"nil\"")

	someString = "6432.8g53"
	dotPosition = 4
	receivedNumber, err = floatConverter(someString, dotPosition)
	expectedNumber = 0.0
	assert.Equal(t, expectedNumber, receivedNumber, "expected \"0.0\", but received smth different")
	assert.NotNil(t, err, "expected \"smth\", but received smth \"nil\"")

}

func TestHexadecimalConverter(t *testing.T) {
	someString := "7D6"
	receivedNumber, err := hexadecimalConverter(someString)
	expectedNumber := 2006.0
	assert.Equal(t, expectedNumber, receivedNumber, "expected \"2006\", but received smth different")
	assert.Nil(t, err, "expected \"nil\", but received smth different")

	someString = "7Df6"
	receivedNumber, err = hexadecimalConverter(someString)
	expectedNumber = 0.0
	assert.Equal(t, expectedNumber, receivedNumber, "expected \"0.0\", but received smth different")
	assert.NotNil(t, err, "expected \"smth\", but received \"nil\"")
}

func TestTypeDefine(t *testing.T) {
	someTypes := map[string]bool{
		"binary":      true,
		"int":         true,
		"float":       true,
		"wrongData":   true,
		"hexadecimal": true,
	}

	expectedType := "wrongData"
	receivedType, err := typeDefine(someTypes, expectedType)
	assert.Equal(t, receivedType, expectedType, "expected \"wrongData\", but received smth different")
	assert.NotNil(t, err, "expected \"smth\", but received \"nil\"")

	/*someTypes["binary"] = false
	someTypes["int"] = false
	someTypes["float"] = false
	someTypes["wrongData"] = false
	someTypes["hexadecimal"] = false*/
	for k, _ := range someTypes {
		someTypes[k] = false
	}
	expectedType = ""
	receivedType, err = typeDefine(someTypes, expectedType)
	assert.Equal(t, receivedType, expectedType, "expected \"nothing\", but received smth different")
	assert.NotNil(t, err, "expected \"smth\", but received \"nil\"")

	someTypes["binary"] = true
	someTypes["int"] = true
	someTypes["hexadecimal"] = true
	someTypes["float"] = true
	expectedType = ""
	receivedType, err = typeDefine(someTypes, expectedType)
	assert.Equal(t, receivedType, expectedType, "expected \"nothing\", but received smth different")
	assert.NotNil(t, err, "expected \"smth\", but received \"nil\"")

	for k, _ := range someTypes {
		someTypes[k] = false
	}

	someTypes["int"] = true
	expectedType = "int"
	receivedType, err = typeDefine(someTypes, expectedType)
	assert.Equal(t, receivedType, expectedType, "expected \"int\", but received smth different")
	assert.Nil(t, err, "expected \"nil\", but received smth different")
	someTypes["int"] = false

	someTypes["float"] = true
	expectedType = "float"
	receivedType, err = typeDefine(someTypes, expectedType)
	assert.Equal(t, receivedType, expectedType, "expected \"float\", but received smth different")
	assert.Nil(t, err, "expected \"nil\", but received smth different")
	someTypes["float"] = false

	someTypes["binary"] = true
	expectedType = "binary"
	receivedType, err = typeDefine(someTypes, expectedType)
	assert.Equal(t, receivedType, expectedType, "expected \"binary\", but received smth different")
	assert.Nil(t, err, "expected \"nil\", but receivec smth different")
	someTypes["binary"] = false

	someTypes["hexadecimal"] = true
	expectedType = "hexadecimal"
	receivedType, err = typeDefine(someTypes, expectedType)
	assert.Equal(t, receivedType, expectedType, "expected \"hexadecimal\", but receibed smth different")
	assert.Nil(t, err, "expected\"nil\", but received smth different")

	someTypes["int"] = true //pass thi test from time to time
	receivedType, err = typeDefine(someTypes, expectedType)
	assert.Equal(t, receivedType, expectedType, "expected \"hexadecimal\", but received smth different")
	assert.Nil(t, err, "expected \"nil\", but received smth different")

	receivedType, err = typeDefine(someTypes, "gexadecimol")
	expectedType = ""
	assert.Equal(t, receivedType, expectedType, "expected \"nothing\", but received smth different")
	assert.NotNil(t, err, "expected \"smth\", but received \"nil\"")
}

func TestMyStrNum1(t *testing.T) {
	someString := "fdks351A"
	receivedNumber, err := myStrToNum1(someString, "wrongData")
	var expectedNumber float64 = 0.0
	assert.Equal(t, receivedNumber, expectedNumber, "expected \"0.0\", but received smth different")
	assert.NotNil(t, err, "expected \"smth\", but received \"nil\"")

	someString = "11101010"
	receivedNumber, err = myStrToNum1(someString, "binary")
	expectedNumber = 234.0
	assert.Equal(t, receivedNumber, expectedNumber, "expected \"234\", but received smth different")
	assert.Nil(t, err, "expected \"nil\", but received different")

	someString = "321"
	receivedNumber, err = myStrToNum1(someString, "int")
	expectedNumber = 321
	assert.Equal(t, receivedNumber, expectedNumber, "expected \"321\", but received smth different")
	assert.Nil(t, err, "expected \"nil\", but received smth different")

	someString = "1.6"
	receivedNumber, err = myStrToNum1(someString, "float")
	expectedNumber = 1.6
	assert.Equal(t, receivedNumber, expectedNumber, "expected \"1.632\", but received smth different")
	assert.Nil(t, err, "expected \"nil\", but received \"smth\"")

	someString = "F38A"
	receivedNumber, err = myStrToNum1(someString, "hexadecimal")
	expectedNumber = 62346
	assert.Equal(t, receivedNumber, expectedNumber, "expected \"62346\", but received smth different")
	assert.Nil(t, err, "expected \"nil\", but received \"smth\"")
}

func TestMyStrNum2(t *testing.T) {
	someString := "fdks351A"
	receivedNumber, err := myStrToNum2(someString)
	var expectedNumber float64 = 0.0
	assert.Equal(t, receivedNumber, expectedNumber, "expected \"0.0\", but received smth different")
	assert.NotNil(t, err, "expected \"smth\", but received \"nil\"")

	someString = "11101010"
	receivedNumber, err = myStrToNum2(someString)
	expectedNumber = 234.0
	assert.Equal(t, receivedNumber, expectedNumber, "expected \"234.0\", but received smth different")
	assert.Nil(t, err, "expected \"nil\", but received different")

	someString = "321"
	receivedNumber, err = myStrToNum2(someString)
	expectedNumber = 321.0
	assert.Equal(t, receivedNumber, expectedNumber, "expected \"5321\", but received smth different")
	assert.Nil(t, err, "expected \"nil\", but received smth different")

	someString = "1.6"
	receivedNumber, err = myStrToNum2(someString)
	expectedNumber = 1.6
	assert.Equal(t, receivedNumber, expectedNumber, "expected \"1.6\", but received smth different")
	assert.Nil(t, err, "expected \"nil\", but received \"smth\"")

	someString = "F38A"
	receivedNumber, err = myStrToNum2(someString)
	expectedNumber = 62346
	assert.Equal(t, receivedNumber, expectedNumber, "expected \"62345\", but received smth different")
	assert.Nil(t, err, "expected \"nil\", but received \"smth\"")

}