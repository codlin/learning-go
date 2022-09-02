package main

import "fmt"

const defaultPath = "/usr/bin/"

type Config struct {
	path string
}

func (c *Config) Path() string {
	if c == nil {
		return defaultPath
	}
	return c.path
}

func main() {
	// 值为nil的变量可以作为函数的接收者：
	var c1 *Config
	var c2 = &Config{
		path: "/usr/local/bin/",
	}
	fmt.Println(c1.Path(), c2.Path())
}
