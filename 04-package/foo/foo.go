package foo

import "fmt"

var a = 1
var B = "world"
var C string
var D string

func init() {
	C = "foo"
	fmt.Println("foo init 1 was called")
}

func init() {
	a = 1
	fmt.Println("foo init 2 was called")
}
