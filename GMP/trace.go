package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}

	fmt.Println("hello GMP")

	var a uint = 1
	var b uint = 2
	fmt.Println(a - b)

	trace.Stop()
}
