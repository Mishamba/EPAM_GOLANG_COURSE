package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

func multiplyByTwo(k *int) error {
	*k *= 2
	return nil
}

func printMoreTen(k int64) error {
	if k < 10 {
		return errors.New("k must be > 10")
	}
	fmt.Println(k)
	return nil
}

func dejson() error {
	type jsStruct struct {
		Data int  `json:"data"`
		OK   bool `json:"ok"`
	}
	in := []byte(`{"data": 13, "ok": true}`)
	var out jsStruct
	if err := json.Unmarshal(in, out); err != nil {
		panic(err)
	}

	return nil
}

func main() {
	var r int = 11
	//var k = 10
	err := multiplyByTwo(&r)
	if err != nil {

	}

	if err := printMoreTen(int64(r)); err != nil {
		panic(err)
	}
}
