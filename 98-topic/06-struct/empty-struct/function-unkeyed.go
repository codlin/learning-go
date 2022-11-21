package main

import "fmt"

/**
# 用途
由于空结构体占用的空间大小为零，我们可以利用这个特性，完成一些功能，却不需要占用额外空间。
**/

/**
# 阻止 unkeyed 方式初始化结构体
**/
type MustKeyedStruct struct {
	Name string
	Age  int
	_    struct{}
}

func main() {
	person := MustKeyedStruct{Name: "Yu", Age: 18}
	fmt.Println(person)

	// person2 := MustKeyedStruct{"hello", 10} //编译失败，提示： too few values in MustKeydStruct{...}
	// fmt.Println(person2)
}
