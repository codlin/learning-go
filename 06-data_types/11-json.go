package main

import (
	"encoding/json"
	"fmt"
	"log"
)

/*
JSON

Java Script Object Notation(JSON)是一种用于发送和接收结构化信息的标准协议。
Go语言中由encoding/json包对JSON提供支持。

JSON是对JavaScript中的各种数据类型--字符串、数字、布尔和对象以Unicode文本进行编码的方式。
基本的JSON数据类型有数字、布尔、字符串，其中字符串以双引号包含的Unicode字符序列，支持和Go语言类似的反斜杠转义特性，
不过JOSN使用的是\Uhhhh转义数字来表示一个UTF-16编码，而不是Go语言的rune类型。
这些基础对象可以通过JSON的数组和对象类型进行递归组合。
一个JSON数组是一个有序的值序列，写在一个方括号中并以逗号分隔；一个JSON数组可以用于编码Go语言的数组和slice。
一个JOSN对象是一个字符串到值的映射，写成以系列的name:value对形式，以花括号包含并以逗号分隔；JSON对象可以用于编码Go语言的map类型和结构体。
*/
type Movie struct {
	Title    string // 如果json中字段和Go语言中一致，则可以不使用tag
	subtitle string // 如果字段未导出，则不会生成json字段
	Year     int    `json:"release"`
	Color    bool   `json:"color,omitempty"`
	Actors   []string
}

/* 将一个Go语言的结构体转换成JSON的过程叫编组Marshaling。编组通过json.Marshal函数完成 */
func json_marshal() {
	var movies = []Movie{
		{Title: "Casablanca", Year: 1942, Color: false,
			Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
		{Title: "Cool Hand Luke", Year: 1967, Color: true,
			Actors: []string{"Paul Newman"}},
		{Title: "Bullitt", Year: 1968, Color: true,
			Actors: []string{"Steve McQueen", "Jacqueline Bisset"}},
	}

	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)
}

func main() {
	json_marshal()
}
