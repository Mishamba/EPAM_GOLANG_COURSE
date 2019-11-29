package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWrongData(t *testing.T) {
	someString := ".bdA562"
	tmpChan := make(chan string)
	receiveType := map[string]bool{
		"wrongData": false,
	}
	wrongDataCheck(receiveType, someString, tmpChan)
	//	<-tmpChan
	assert.True(t, receiveType["wrongData"], "expected \"TRUE\" but received \"nothing\"")

	someString = "B.dA562"
	receiveType = map[string]bool{
		"wrongData": false,
	}
	wrongDataCheck(receiveType, someString, tmpChan)
	//<-tmpChan
	assert.True(t, receiveType["wrongData"], "expected \"TRUE\" but received \"nothing\"")

	someString = "B.562"
	receiveType = map[string]bool{
		"wrongData": false,
	}
	wrongDataCheck(receiveType, someString, tmpChan)
	//<-tmpChan
	assert.True(t, receiveType["wrongData"], "expected \"TRUE\" but received \"nothing\"")

	someString = "14"
	receiveType["wrongData"] = true
	wrongDataCheck(receiveType, someString, tmpChan)
	//<-tmpChan
	assert.True(t, !receiveType["wrongData"], "expected \"FALSE\", but received \"TRUE\"")
}

func TestBinaryCheck(t *testing.T) {
	someString := "111001"
	tmpChan := make(chan string)
	receiveType := map[string]bool{
		"binary": false,
	}
	binaryCheck(receiveType, someString, tmpChan)
	//<-tmpChan
	assert.True(t, receiveType["binary"], "expected \"TRUE\" but received \"nothing\"")

	someString = "101f0201"
	binaryCheck(receiveType, someString, tmpChan)
	//<-tmpChan
	assert.True(t, !receiveType["binary"], "expected \"FALSE\", but received \"TRUE\"")
}

func TestIntCheck(t *testing.T) {
	someString := "1583"
	tmpChan := make(chan string)
	receiveType := map[string]bool{
		"int": false,
	}
	intCheck(receiveType, someString, tmpChan)
	//<-tmpChan
	assert.True(t, receiveType["int"], "expected \"TRUE\" but received \"nothing\"")

	someString = "6.32gs9"
	receiveType["int"] = true
	intCheck(receiveType, someString, tmpChan)
	//<-tmpChan
	assert.True(t, !receiveType["int"], "expected \"FALSE\", but received \"TRUE\"")
}

func TestFloatCheck(t *testing.T) {
	someString := "1286.8643"
	tmpChan := make(chan string)
	receiveType := map[string]bool{
		"float": false,
	}
	var receiveDotPosition int
	expectedDotPosition := 4
	floatCheck(receiveType, someString, &receiveDotPosition, tmpChan)
	//<-tmpChan
	assert.True(t, receiveType["float"], "expected \"TRUE\" but received \"nothing\"")
	assert.Equal(t, expectedDotPosition, receiveDotPosition, "dot expected at \"4\" position, but received another")

	someString = "1.g463A"
	receiveType["float"] = true
	//expectedDotPosition = 1		//if same test passed succesfuly, so we don't need to run it again with another data?
	floatCheck(receiveType, someString, &receiveDotPosition, tmpChan)
	//<-tmpChan
	assert.True(t, !receiveType["float"], "expected \"FALSE\", but received \"TRUE\"")
	//assert.Equal(t, expectedDotPosition, receiveDotPosition, "dot expected at \"1\" position, but received another")

	someString = ".463"
	receiveType["float"] = true
	floatCheck(receiveType, someString, &receiveDotPosition, tmpChan)
	//<-tmpChan
	assert.True(t, !receiveType["float"], "expected \"FALSE\", but received \"TRUE\"")

}

func TestHexadecimalCheck(t *testing.T) {
	someString := "3A632F"
	tmpChan := make(chan string)
	receiveType := map[string]bool{
		"hexadecimal": false,
	}
	hexadecimalCheck(receiveType, someString, tmpChan)
	//<-tmpChan
	assert.True(t, receiveType["hexadecimal"], "expected \"TRUE\" but received \"nothing\"")

	someString = "53Fb32uA"
	receiveType["hexadecimal"] = true
	hexadecimalCheck(receiveType, someString, tmpChan)
	//<-tmpChan
	assert.True(t, !receiveType["hexadecimal"], "expected \"FALSE\", but received \"TRUE\"")
}
