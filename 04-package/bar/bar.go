package bar

import (
	"fmt"

	"mypackage/foo"
)

var a = 1
var B = "hello"
var C string

func init() {
	C = "bar"
	fmt.Println("bar init 1 was called")
}

func init() {
	a = 1
	foo.D = "modified"

	fmt.Println("bar init 2 was called")
}
