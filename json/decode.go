package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type Message struct {
	Name   string  `json:"name"`
	Text   string  `json:"text"`
	Age    int     `json:"age"`
	Tall   float32 `json:"tall,omitempty"`
	Weight float32 `json:"-"`
}

type Message2 struct {
	Name   *string  `json:"name"`
	Text   *string  `json:"text"`
	Age    *int     `json:"age"`
	Tall   *float32 `json:"tall,omitempty"`
	Weight *float32 `json:"-"`
}

func decodeStruct(msg string) {
	dec := json.NewDecoder(strings.NewReader(msg))

	// read open bracket
	t, err := dec.Token()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T: %v\n", t, t)

	// while the array contains values
	for dec.More() {
		var m Message
		// decode an array value (Message)
		err := dec.Decode(&m)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v: %v - %v - %v - %v\n", m.Name, m.Text, m.Age, m.Tall, m.Weight)
	}

	// read closing bracket
	t, err = dec.Token()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T: %v\n", t, t)
}

func decodeStructPoint(msg string) {
	dec := json.NewDecoder(strings.NewReader(msg))

	// read open bracket
	t, err := dec.Token()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T: %v\n", t, t)

	// while the array contains values
	for dec.More() {
		var m Message2
		// decode an array value (Message)
		err := dec.Decode(&m)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v: %v - %v - %v - %v\n", m.Name, m.Text, m.Age, m.Tall, m.Weight)
	}

	// read closing bracket
	t, err = dec.Token()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T: %v\n", t, t)
}

func main() {
	const jsonStream = `
	[
		{"name": "Ed", "text": "Knock knock.", "weight": 65},
		{"name": "Sam", "text": "Who's there?"},
		{"name": "Ed", "text": "Go fmt.", "tall": 175},
		{"name": "Sam", "text": "Go fmt who?", "weight": 80},
		{"name": "Ed", "text": "Go fmt yourself!"}
	]`

	decodeStruct(jsonStream)

	decodeStructPoint(jsonStream)

	/*
	 * 可以看的到：
	 * 如果结构体中定义的字段在json中找不到，则用该字段类型的默认值填充
	 * 如果结构体中定义的字段有注释"-"，则在decode时忽略该字段的值，即使传入的json中该字段有值
	 * 如果结构体中是指针类型，则在decode后存储对应值的指针。这很适合用来判断json中有某个字段
	 */
}
