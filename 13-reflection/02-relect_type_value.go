package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"time"
)

/*
reflect.Type和reflect.Value

反射功能由reflect包提供，它定义了两个重要的类型：Type和Value。
Type表示Go语言的一个类型，它是一个有很多方法的接口，这些方法可以用来识别类型以及透视类型的组成部分，
比如一个结构的组成部分或者一个函数的各个参数。
reflect.Type接口只要一个实现，即类型描述符。接口值中的动态类型也是类型描述符。
reflect.TypeOf函数接受任何类型的interface{}参数，并且把接口中的动态类型以reflect.Type的形式返回。
reflect.ValueOf函数接受任何类型的interface{}参数，并且把接口中的动态值以reflect.Value的形式返回。
*/

func reflect_demo() {
	/* 把一个具体值赋给一个接口类型时会发生一个隐式类型转换，转换会生成一个包含两部分内容的接口值：
	动态类型和动态值
	此处，动态类型是int，动态值是3
	*/
	t := reflect.TypeOf(3)  // 一个reflect.Type
	fmt.Println(t.String()) // "int"
	fmt.Println(t)          // "int"
}

/*
因为reflect.TypeOf返回一个接口的动态类型，所以它返回的总是具体类型（而不是接口类型）。
与reflect.TypeOf类似，reflect.ValueOf返回的也是一个具体值，但它也可以包含一个接口值。
*/
func reflect_stdout() {
	var w io.Writer = os.Stdout
	fmt.Println(reflect.TypeOf(w)) // "*os.File"

	v := reflect.ValueOf(3) // 3, 一个reflect.Value
	fmt.Println(v)          // "3"
	/*
		另一个与reflect.Type类似的是，reflect.Value也满足
		fmt.Stringer，但除非Value包含的是一个字符串，否则String方法的
		结果仅仅暴露了类型。通常，你需要使用fmt包的%v功能，它对
		reflect.Value会进行特殊处理。
	*/
	fmt.Println(v.String()) // 注意: "<int Value>"
	fmt.Printf("%v\n", v)   // "3"
}

/*
reflect.ValueOf的逆操作是reflect.Value.Interface方法。它返回一个interface{}接口值，与reflect.Value包含同一个具体值。
*/
func reflect_value_interface() {
	v := reflect.ValueOf(3) // a reflect.Value
	x := v.Interface()      // an interface{}
	i := x.(int)            // an int
	fmt.Printf("%d\n", i)   // 3
}

// Any formats any value as a string.
func Any(value interface{}) string {
	return formatAtom(reflect.ValueOf(value))
}

// formatAtom formats a value without inspecting its internal structure.
func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	// ...floating-point and complex cases omitted for brevity...
	case reflect.Bool:
		return strconv.FormatBool(v.Bool())
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr, reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" +
			strconv.FormatUint(uint64(v.Pointer()), 16)
	default: // reflect.Array, reflect.Struct, reflect.Interface
		return v.Type().String() + " value"
	}
}

func formatAtom_test() {
	var x int64 = 1
	var d time.Duration = 1 * time.Nanosecond
	fmt.Println(Any(x))                  // "1"
	fmt.Println(Any(d))                  // "1"
	fmt.Println(Any([]int64{x}))         // "[]int64 0x8202b87b0"
	fmt.Println(Any([]time.Duration{d})) // "[]time.Duration 0x8202b87e0"
}

func Display(name string, x interface{}) {
	fmt.Printf("Display %s (%T):\n", name, x)
	display(name, reflect.ValueOf(x))
}

/*
reflect.Value.Kind()返回值的类型
reflect.Value.NumField()返回结构体字段的下标
reflect.Value.Field(i)返回结构体的第i个字段
reflect.Value.MapKeys()返回一个包含map的key值的slice，这些key值和顺序无关。
*/
func display(path string, v reflect.Value) {
	switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s = invalid\n", path)
	case reflect.Slice, reflect.Array:
		for i := 0; i < v.Len(); i++ {
			display(fmt.Sprintf("%s[%d]", path, i), v.Index(i))
		}
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
			display(fieldPath, v.Field(i))
		}
	case reflect.Map:
		for _, key := range v.MapKeys() {
			display(fmt.Sprintf("%s[%s]", path,
				formatAtom(key)), v.MapIndex(key))
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			display(fmt.Sprintf("(*%s)", path), v.Elem())
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			fmt.Printf("%s.type = %s\n", path, v.Elem().Type())
			display(path+".value", v.Elem())
		}
	default: // basic types, channels, funcs
		fmt.Printf("%s = %s\n", path, formatAtom(v))
	}
}

type Movie struct {
	Title, Subtitle string
	Year            int
	Color           bool
	Actor           map[string]string
	Oscars          []string
	Sequel          *string
}

func display_test() {
	strangelove := Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		Year:     1964,
		Color:    false,
		Actor: map[string]string{
			"Dr. Strangelove":            "Peter Sellers",
			"Grp. Capt. Lionel Mandrake": "Peter Sellers",
			"Pres. Merkin Muffley":       "Peter Sellers",
			"Gen. Buck Turgidson":        "George C. Scott",
			"Brig. Gen. Jack D. Ripper":  "Sterling Hayden",
			`Maj. T.J. "King" Kong`:      "Slim Pickens",
		}, Oscars: []string{
			"Best Actor (Nomin.)",
			"Best Adapted Screenplay (Nomin.)",
			"Best Director (Nomin.)",
			"Best Picture (Nomin.)",
		},
	}
	Display("strangelove", strangelove)
}

func main() {
	reflect_demo()
	reflect_stdout()
	reflect_value_interface()
	formatAtom_test()
	display_test()
}
