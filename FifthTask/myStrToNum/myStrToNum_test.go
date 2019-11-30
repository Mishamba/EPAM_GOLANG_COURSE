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
	wg.Add(1)
	wrongDataCheck(receiveType, someString)
	wg.Wait()
	assert.True(t, receiveType["wrongData"], "expected \"TRUE\" but received \"nothing\"")

	someString = "B.dA562"
	receiveType = map[string]bool{
		"wrongData": false,
	}
	wg.Add(1)
	wrongDataCheck(receiveType, someString)
	wg.Wait()
	assert.True(t, receiveType["wrongData"], "expected \"TRUE\" but received \"nothing\"")

	someString = "B.562"
	receiveType = map[string]bool{
		"wrongData": false,
	}
	wg.Add(1)
	wrongDataCheck(receiveType, someString)
	wg.Wait()
	assert.True(t, receiveType["wrongData"], "expected \"TRUE\" but received \"nothing\"")

	someString = "14"
	receiveType["wrongData"] = false
	wg.Add(1)
	wrongDataCheck(receiveType, someString)
	wg.Wait()
	assert.True(t, !receiveType["wrongData"], "expected \"FALSE\", but received \"TRUE\"")

	someString = ".463"
	receiveType["wrongData"] = true
	wg.Add(1)
	wrongDataCheck(receiveType, someString)
	wg.Wait()
	assert.True(t, receiveType["wrongData"], "expected \"FALSE\", but received \"TRUE\" 1")
}

func TestBinaryCheck(t *testing.T) {
	someString := "111001"
	receiveType := map[string]bool{
		"binary": false,
	}
	wg.Add(1)
	binaryCheck(receiveType, someString)
	wg.Wait()
	assert.True(t, receiveType["binary"], "expected \"TRUE\" but received \"nothing\"")

	someString = "101f0201"
	wg.Add(1)
	binaryCheck(receiveType, someString)
	wg.Wait()
	assert.True(t, !receiveType["binary"], "expected \"FALSE\", but received \"TRUE\"")
}

func TestIntCheck(t *testing.T) {
	someString := "1583"
	receiveType := map[string]bool{
		"int": false,
	}
	wg.Add(1)
	intCheck(receiveType, someString)
	wg.Wait()
	assert.True(t, receiveType["int"], "expected \"TRUE\" but received \"nothing\"")

	someString = "6.32gs9"
	receiveType["int"] = true
	wg.Add(1)
	intCheck(receiveType, someString)
	wg.Wait()
	assert.True(t, !receiveType["int"], "expected \"FALSE\", but received \"TRUE\"")
}

func TestFloatCheck(t *testing.T) {
	someString := "1286.8643"
	receiveType := map[string]bool{
		"float": false,
	}
	var receiveDotPosition int
	expectedDotPosition := 4
	wg.Add(1)
	floatCheck(receiveType, someString, &receiveDotPosition)
	wg.Wait()
	assert.True(t, receiveType["float"], "expected \"TRUE\" but received \"nothing\"")
	assert.Equal(t, expectedDotPosition, receiveDotPosition, "dot expected at \"4\" position, but received another")

	someString = "1.g463A"
	receiveType["float"] = true
	//expectedDotPosition = 1		//if same test passed succesfuly, so we don't need to run it again with another data?
	wg.Add(1)
	floatCheck(receiveType, someString, &receiveDotPosition)
	wg.Wait()
	assert.True(t, !receiveType["float"], "expected \"FALSE\", but received \"TRUE\"")
	//assert.Equal(t, expectedDotPosition, receiveDotPosition, "dot expected at \"1\" position, but received another")

	someString = ".463"
	receiveType["float"] = true
	wg.Add(1)
	floatCheck(receiveType, someString, &receiveDotPosition)
	wg.Wait()
	assert.True(t, !receiveType["float"], "expected \"FALSE\", but received \"TRUE\"")

}

func TestHexadecimalCheck(t *testing.T) {
	someString := "3A632F"
	receiveType := map[string]bool{
		"hexadecimal": false,
	}
	wg.Add(1)
	hexadecimalCheck(receiveType, someString)
	wg.Wait()
	assert.True(t, receiveType["hexadecimal"], "expected \"TRUE\" but received \"nothing\"")

	someString = "53Fb32uA"
	receiveType["hexadecimal"] = true
	wg.Add(1)
	hexadecimalCheck(receiveType, someString)
	wg.Wait()
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

	someString = "6332a853"
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

	receivedType, err := typeDefine(someTypes)
	expectedType := "wrongData"
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
	receivedType, err = typeDefine(someTypes)
	expectedType = ""
	assert.Equal(t, receivedType, expectedType, "expected \"nothing\", but received smth different")
	assert.NotNil(t, err, "expected \"smth\", but received \"nil\"")

	someTypes["binary"] = true
	someTypes["int"] = true
	someTypes["hexadecimal"] = true
	someTypes["float"] = true
	receivedType, err = typeDefine(someTypes)
	expectedType = ""
	assert.Equal(t, receivedType, expectedType, "expected \"nothing\", but received smth different")
	assert.NotNil(t, err, "expected \"smth\", but received \"nil\"")

	for k, _ := range someTypes {
		someTypes[k] = false
	}

	someTypes["int"] = true
	receivedType, err = typeDefine(someTypes)
	expectedType = "int"
	assert.Equal(t, receivedType, expectedType, "expected \"int\", but received smth different")
	assert.Nil(t, err, "expected \"nil\", but received smth different")
	someTypes["int"] = false

	someTypes["float"] = true
	receivedType, err = typeDefine(someTypes)
	expectedType = "float"
	assert.Equal(t, receivedType, expectedType, "expected \"float\", but received smth different")
	assert.Nil(t, err, "expected \"nil\", but received smth different")
	someTypes["float"] = false

	someTypes["binary"] = true
	receivedType, err = typeDefine(someTypes)
	expectedType = "binary"
	assert.Equal(t, receivedType, expectedType, "expected \"binary\", but received smth different")
	assert.Nil(t, err, "expected \"nil\", but receivec smth different")
	someTypes["binary"] = false

	someTypes["hexadecimal"] = true
	receivedType, err = typeDefine(someTypes)
	expectedType = "hexadecimal"
	assert.Equal(t, receivedType, expectedType, "expected \"hexadecimal\", but receibed smth different")
	assert.Nil(t, err, "expected\"nil\", but received smth different")

	someTypes["int"] = true //pass thi test from time to time
	receivedType, err = typeDefine(someTypes)
	expectedType = "hexadecimal"
	assert.Equal(t, receivedType, expectedType, "expected \"int\", but received smth different")
	assert.Nil(t, err, "expected \"nil\", but received smth different")
}

func TestMyStrNum(t *testing.T) {
	someString := "fdks351A"
	receivedNumber, err := myStrToNum(someString)
	var expectedNumber float64 = 0.0
	assert.Equal(t, receivedNumber, expectedNumber, "expected \"0.0\", but received smth different")
	assert.NotNil(t, err, "expected \"smth\", but received \"nil\"")

	/*someString = "100101"
	receivedNumber, err = myStrToNum(someString)
	expectedNumber = //smth
	assert.Equal(t, receivedNumber, expectedNumber, "expected \"351\", but received smth different")
	assert.Nil(t, err, "expected \"nil\", but received \"smth\"")*/

	someString = "1.6"
	receivedNumber, err = myStrToNum(someString)
	expectedNumber = 1.6
	assert.Equal(t, receivedNumber, expectedNumber, "expected \"1.632\", but received smth different")
	assert.Nil(t, err, "expected \"nil\", but received \"smth\"")

	someString = "F38A"
	receivedNumber, err = myStrToNum(someString)
	expectedNumber = 62346
	assert.Equal(t, receivedNumber, expectedNumber, "expected \"62346\", but received smth different")
	assert.Nil(t, err, "expected \"nil\", but received \"smth\"")
}
