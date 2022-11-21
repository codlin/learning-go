package main

import (
	"fmt"
)

type MyError struct {
	Code int
}

func (e *MyError) Error() string {
	if e != nil {
		return fmt.Sprintf("errCode: %v\n", e.Code)
	}
	return ""
}

func returnError1(err bool) error {
	var p *MyError = nil
	if err {
		p = &MyError{Code: 404}
	}
	fmt.Printf("before covert to error, err == nil: %v\n", p == nil)
	return p
}

func returnError2(err bool) error {
	if err {
		return &MyError{404}
	}

	return nil
}

func checkError(err error) {
	if err == nil {
		println("after covert to error, err is nil.")
		return
	}
	println("after covert to error, err is not nil.", err.Error())
}

func main() {
	expectNil := returnError1(false)
	checkError(expectNil)

	actualNil := returnError2(false)
	checkError(actualNil)

	err := returnError2(true)
	checkError(err)
}
