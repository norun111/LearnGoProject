package main

import (
	"errors"
	"fmt"
)

//func example(code string) (int, error) {
//	if code == "hoge" {
//		return 1, nil
//	}
//
//	return 0, errors.New("code must be hoge")
//}

type Hex int

func (h Hex) String() string {
	return fmt.Sprintf("%x", int(h))
}
