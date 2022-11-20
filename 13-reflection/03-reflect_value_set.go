package main

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

/*
使用reflect.Value来设置值

在Go语言中，有些变量是可寻址，比如x, x.i, x[1]；而有些作为值不能寻址，如x+1, f(2)等等。
对reflect.Value也有类似的区分，某些是可寻址的，而其它的并非如此。
事实上，通过reflect.ValueOf(x)返回的Value都是不可寻址的。
可以通过变量的CanAddr来询问relfect.Value变量是否可寻址。
关于可寻址，参见https://go.dev/ref/spec#Address_operators
https://colobu.com/2018/02/27/go-addressable/
我们可以通过指针间接获取一个可寻址的reflect.Value，即使这个指针是不可寻址的。可寻址的常见规则都在relfect包里有对应项。
比如slice的下标表达式e[i]隐式地使用了指针（follows a pointer），即使e是不可寻址的，但是e[i]是可寻址的。
类似的，reflect.VauleOf(e).Index(i)代表一个变量，尽管reflect.VauleOf(e)是不可寻址的，但是它可以。
但不能通过&操作符直接寻址，可以使用reflect.Value().Addr()来寻址。
*/
func reflect_value_ref() {
	a := reflect.ValueOf(2)
	fmt.Printf("%d\t%p\n", a.Int(), &a)         // 2 xc0000a2078
	fmt.Printf("can a addr: %t\n", a.CanAddr()) // false
	// &reflect.ValueOf(2) // compile error

	x := 2
	fmt.Printf("&x: %p\n", &x) // 0xc0000ac0f0

	// 将接口里的动态值以reflect.Value返回。即x的值
	b := reflect.ValueOf(x)
	// 动态类型大概是这样的： interface { type: int, ptr: null}
	// fmt.Println(b.Type(), b.Pointer()) // b.Pointer panic
	fmt.Println(b.Type())

	// &b 只是取了变量b的内存地址，无法取到x的地址
	fmt.Printf("%d\t%p\n", b.Int(), &b)         // 2 0xc0000a2090
	fmt.Printf("can b addr: %t\n", b.CanAddr()) // false
	// 不能对value使用Elem()，因为它既不是接口也不是指针
	// b.Elem()                                    // runtime panic

	// 将接口里的动态值以reflect.Value返回。即x的地址
	c := reflect.ValueOf(&x)
	// 动态类型大概是这样的： interface { type: *int, ptr: 0xc0000ac0f0}
	fmt.Printf("c type: %s, value: 0x%x\n", c.Type(), c.Pointer())

	// 变量c保存的是x的地址，&c取得的是c的内存地址
	fmt.Printf("%p\t%v\n", &c, c)               // 0xc0000a20a8    0xc0000ac0f0
	fmt.Printf("can c addr: %t\n", c.CanAddr()) // false
	// fmt.Println(c.Addr()) // runtime panic

	//事实上，通过reflect.ValueOf(x)返回的Value都是不可寻址的。
	// 但 d 是通过 c 中的指针解引用（dereferencing）导出的（derived），因此它是可寻址的。
	// 可以通过这个方法，调用reflect.ValueOf(&x).Elem() 来获得任意变量x可寻址的Value值。
	// 为什么Elem可以寻址，源代码显示ValueOf只对Value的flag打上了flagIndir，并没有打上
	// flagAddr标记，在用Addr()寻址时检测到没有这个标记直接panic。而Elem()会根据Value的
	// 动态类型进行处理，对于指针类型，会取出指针并打上flagAddr，这样在Addr()就可以用了。
	d := c.Elem()                               //
	fmt.Printf("%d\t%p\n", d.Int(), &d)         // 2 0xc0000a20d8
	fmt.Printf("can d addr: %t\n", d.CanAddr()) // true
	fmt.Println(d.Addr())                       // 0xc0000ac0f0 变量x的地址

	e := []int{1, 2, 3}
	// fmt.Println(&reflect.ValueOf(e)) // compile error, 不可寻址
	// &reflect.ValueOf(e).Index(1)     // compile error 不能用这种方式来寻址
	// reflect.ValueOf(e).Addr()        // runtime panic，不可对不能寻址的值进行寻址
	reflect.ValueOf(e).Index(1).Addr() // ok
}

/*
从一个可寻址的reflect.Value()获取变量需要三步：
1. 调用Addr()，返回一个Value，其中包含一个指向变量的指针，
2. 在这个Value上调用Interface()，会返回一个包含这个指针的interface{}值
3. 最后，我们知道这个变量的类型，使用类型断言把接口内容转换成普通的指针。
之后就可以通过这个指针更新变量了。
*/
func reflect_value_update() {
	fmt.Println("addressable_value_update")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	x := 2
	d := reflect.ValueOf(&x).Elem()   // d代表变量x，包含变量x的指针
	px := d.Addr().Interface().(*int) // 取出变量x的地址，并通过类型断言转成*int； px := &x
	*px = 3                           // x= 3
	fmt.Println(x)                    // 3
}

/*
但是还可以通过可寻址的reflect.Value来更新变量，不通过指针，而是直接调用reflect.Value.Set方法：
d.Set(reflect.ValueOf(4))
*/
func reflect_value_set() {
	fmt.Println("reflect_value_set")
	x := 2
	d := reflect.ValueOf(&x).Elem()   // d代表变量x，包含变量x的指针
	px := d.Addr().Interface().(*int) // 取出变量x的地址，并通过类型断言转成*int； px := &x
	*px = 3                           // x= 3
	fmt.Println(x)                    // 3

	d.Set(reflect.ValueOf(4)) // x = 4
	fmt.Println(x)            // 4
}

/*
运用Set方法时编译器不会对值的可赋值性进行检查，需要Set方法在运行时检查，因此确保类型的可赋值性很重要：
d.Set(reflect.ValueOf(int64(4))) // runtime paninc 类型不匹配
还有一些基本类型的特化版本，SetInt、SetUint、SetString、SetFloat等：
d.SetInt(5) // x = 5
这些版本具有一定的容错性。
*/
func reflect_value_set_panic() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	x := 2
	d := reflect.ValueOf(&x).Elem() // d代表变量x，包含变量x的指针

	/* 运用Set方法时编译器不会来检查可赋值性条件，由Set在运行时检查。因此确保类型的可赋值性很重要 */
	// d.Set(reflect.ValueOf(int64(4))) // runtime paninc 类型不匹配

	e := reflect.ValueOf(x)
	// e.Set(reflect.ValueOf(5)) // runtime panic 不能对不可寻址的值设置
	fmt.Println(e)

	f := reflect.ValueOf(&x)
	// f.Set(reflect.ValueOf(6)) // runtime panic 不能对不可寻址的值设置
	fmt.Println(f)

	/*
		我们还有为一些基本类型特化的Set变种：SetInt、SetUint、SetString、SetFloat等：
	*/
	d.SetInt(5)
	fmt.Println("SetInt(5)", x) // x = 5

	// d.SetFloat(5.5) // runtime panic

	var y any
	ry := reflect.ValueOf(&y).Elem()
	// ry.SetInt(2) // runtime panic 指向接口的value调用SetInt
	ry.Set(reflect.ValueOf(6))
	fmt.Printf("y\t%d\n", y) // y = 6
	// ry.SetString("hello")    // runtime panic 指向接口的value调用SetString
	ry.Set(reflect.ValueOf("hello")) // y = "hello"

}

/*
我们发现反射可以越过Go语言的导出规则的限制读取结构体中未导出的成员，
比如在类Unix系统上os.File结构体中的fd int成员。
然而，利用反射机制并不能修改这些未导出的成员。
在更新变量前使用CanAddr()来检查并不能确保正确性。
CanSet()方法才能正确的报告一个reflect.Value是否可以被寻址且可改变。
*/
func reflect_value_can_set() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	stdout := reflect.ValueOf(os.Stdout).Elem() // *os.Stdout, an os.File var
	fmt.Println(stdout.Type())                  // "os.File"
	fd := stdout.FieldByName("fd")

	// ⚠️ 这段和书上有出入
	fmt.Printf("can set: %t, can addr: %t\n", fd.CanSet(), fd.CanAddr()) // false false
	// ⚠️ 此处产生panic：call of reflect.Value.Int on zero Value
	fmt.Println(fd.Int()) // "1" --> runtime panic

	fd.SetInt(2) // panic: unexported field
}

func main() {
	fmt.Printf("%s\n", strings.Repeat("=", 64))

	reflect_value_ref()
	fmt.Printf("%s\n", strings.Repeat("-", 64))

	reflect_value_update()
	fmt.Printf("%s\n", strings.Repeat("-", 64))

	reflect_value_set()
	fmt.Printf("%s\n", strings.Repeat("-", 64))

	reflect_value_set_panic()
	fmt.Printf("%s\n", strings.Repeat("-", 64))

	reflect_value_can_set()
	fmt.Printf("%s\n", strings.Repeat("=", 64))
}
