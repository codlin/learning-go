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

在编码时，默认使用Go语言结构体的成员名字（struct field names）作为JSON的对象，只有导出的结构体成员（field）才会被编码。
如果json字段的名字和Go语言字段不一样，可以使用结构体成员tag（field tag）。
一个field tag是在编译阶段关联到该结构体成员的元信息字符串。A field tag is a string of metadata
associated at compile time with the field of a struct.
一个成员tag可以时任意的字符串字面值，但通常是一系列用空格分隔的key:"value"键值对序列；
因为成员tag值中含有双引号字符，所以成员tag一般用原生字符串面值的形式书写。
json开头键名对应的值用于控制encoding/json包的编码和解码行为，并且encoding/...下面的包也遵循这个约定。
成员tag中json对应值的第一部分用于指定JSON对象的名字，如下面例子，release代表Go成员Year对应的JSON对象的名字
```go

	type Movie struct {
		Title    string // 如果json中字段和Go语言中一致，则可以不使用tag
		subtitle string // 如果字段未导出，则不会生成json字段
		Year     int    `json:"release"`
		Color    bool   `json:"color,omitempty"`
		Actors   []string
	}

```
如果名字后面跟了omitempty，则指定在结构体成员值为零值时不生成JSON对象。
*/
type Movie struct {
	Title    string // 如果JSON对象中该字段和Go语言中一致，则可以不使用tag
	subtitle string // 如果字段未导出，则不会生成json字段
	Year     int    `json:"release"`         // 使用json tag指定任意名字
	Color    bool   `json:"color,omitempty"` // 如果为零值，则在json中不输出该字段
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

	/* Marshal后的结果很紧凑，没有空白字符，打印在一行
	为了生成编译阅读的格式，MarshalIndent函数将产生整齐缩进的输出
	*/
	data, err = json.MarshalIndent(movies, "", "  ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)
}

/* 解码是编码的逆操作，对应将JSON数据解码为Go语言的数据结构，Go语言中一般叫unmarshaling，通过json.Unmarshal函数完成 */
func json_unmarshal() {
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

	/* 通过定义合适的Go语言数据结构，可以选择性的解码JSON中感兴趣的成员 */
	var titles []struct{ Title string }
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s", err)
	}
	fmt.Println(titles)
}

func main() {
	json_marshal()
	json_unmarshal()
}
