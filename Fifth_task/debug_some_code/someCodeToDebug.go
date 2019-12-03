package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type jsStruct struct {
	Data int  `json:"data"`
	OK   bool `json:"ok"`
}

func multiplyByTwo(k *int) error {
	if *k == 0 {
		return errors.New("received zero...")
	}
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

func dejson(out interface{}) error {
	in := []byte(`{"data": 13, "ok": true}`)
	if err := json.Unmarshal(in, out); err != nil {
		return err
	}

	return nil
}
