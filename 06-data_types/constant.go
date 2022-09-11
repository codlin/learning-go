package main

import (
	"fmt"
	"math"
	"time"
)

/*
常量

常量表达式的值在编译器计算，而不是在运行期。每种常量的潜在类型都是基础类型：boolean、数字或字符串。
*/
const (
	pi = 3.14
	e  = 2.71
)

/*
常量之间的所有算术运算、逻辑运算和比较运算的结果也是常量，
对常量的类型转换操作或以下函数调用都是返回常量结果：
len, cap, real, imag, complex 和 unsafe.Sizeof
*/

/* 因为它们的值是在编译期就定的，所以常量可以是构成类型的一部分，例如用于指定数组类型的长度 */
const IPv4Len = 4

func parseIPv4(s string) string {
	var p [IPv4Len]byte
	return string(p[:])
}

/* 一个常量的声明可以包含一个类型和一个值，如果没有显式指明类型，那么将从右边的表达式推断类型 */
func const_type() {
	const noDelay time.Duration = 0
	const timeout = 5 * time.Minute
	fmt.Printf("%T %[1]v\n", noDelay)
	fmt.Printf("%T %[1]v\n", timeout)
	fmt.Printf("%T %[1]v\n", time.Minute)
}

/*
如果是批量声明常量，除了第一个常量外其它常量右边的表达式可以省略，
如果省略初始化表达式则表示使用前面常量表达式的初始化表达式写法，对应的常量类型也是一样的。
*/
func const_definition() {
	const (
		a = 1
		b
		c = 2
		d
	)
	fmt.Println(a, b, c, d) // 1 1 2 2
}

/*
常量声明可以使用iota常量生成器初始化，它用于生成一组以相似规则初始化的常量，但是不会每行都写一遍初始化表达式
在一个常量声明语句中，在第一个声明常量所在的行，iota将会被置为0，后续每行加1
*/
type WeekDay int

const (
	Sunday WeekDay = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
)

type Flags uint

const (
	FlagUp Flags = 1 << iota // is UP
	FlagBroadcast
	FlagLoopback
	FlagPointToPoint
	FlagMulticast
)

func IsUp(v Flags) bool     { return v&FlagUp == FlagUp }
func TurnDown(v *Flags)     { *v &^= FlagUp }
func SetBroadcast(v *Flags) { *v |= FlagBroadcast }
func IsCast(v Flags) bool   { return v&(FlagBroadcast|FlagMulticast) != 0 }
func flag() {
	var v Flags = FlagMulticast | FlagUp
	fmt.Printf("%b %t\n", v, IsUp(v)) // "10001 true"
	TurnDown(&v)
	fmt.Printf("%b %t\n", v, IsUp(v)) // "10000 false"
	SetBroadcast(&v)
	fmt.Printf("%b %t\n", v, IsUp(v))   // "10010 false"
	fmt.Printf("%b %t\n", v, IsCast(v)) // "10010 true"
}

const (
	_   = 1 << (10 * iota)
	KiB // 1024
	MiB // 1048576
	GiB // 1073741824
	TiB // 1099511627776 (exceeds 1 << 32)
	PiB // 1125899906842624
	EiB // 1152921504606846976
	ZiB // 1180591620717411303424 (exceeds 1 << 64)
	YiB // 1208925819614629174706176
)

/*不过iota常量生成规则也有其局限性。例如，它并不能用于产生1000的幂（KB、MB等），因为Go
语言并没有计算幂的运算符。*/

/*
无类型常量 Untyped Constants

Go语言的常量有一点不同寻常之处。虽然一个常量可以有任意一个确定的基础类型，但是许多常量并没有一个明确的基础类型。
编译器为这些没有明确基础类型的数字常量提供比基础类型更高的精度算术运算；你可以认为至少有256bit的运算精度。
这里有六种未明确类型的常量类型：
- 无类型的布尔型
- 无类型的整型
- 无类型的浮点数
- 无类型的复数
- 无类型的字符
- 无类型的字符串

通过延迟明确常量的具体类型，无类型的常量不仅可以提供更高的运算精度，而且可以直接用于更多的表达式而无需显示的类型转换。
ZiB 和 YiB 的值已经超过了Go语言中整数类型的表达范围，但是它们依然是合法的常量，而且像下面的常量表达式依然有效：
fmt.Println(YiB/ZiB) // "1024"
*/
func untype_const() {
	fmt.Println(YiB / ZiB) // "1024"

	// math.Pi无类型的浮点数常量，可以直接用于任意需要浮点数或复数的地方
	var x float32 = math.Pi
	var y float64 = math.Pi
	var z complex128 = math.Pi

	/* 如果math.Pi被确定为特定类型，比如float64，那么结果精度可能会不一样，同时对于需要float32
	或complex128类型值的地方则会强制需要一个明确的类型转换：
	*/
	const Pi64 float64 = math.Pi
	var m float32 = float32(Pi64)
	var n float64 = Pi64
	var t complex128 = complex128(Pi64)

	fmt.Println(x, y, z, m, n, t)
}

/*
对于常量面值，不同的写法可能会对应不同的类型。
例如0、0.0、0i和'\u0000'虽然有着相同的常量值，但是它们分别对应无类型的整数、无类型的浮点数、无类型的复数和无类型的字符等不同的常
量类型。
同样，true和false也是无类型的布尔类型，字符串面值常量是无类型的字符串类型。
*/
func untpyed_const_opt() {
	var f float64 = 212
	fmt.Println((f - 32) * 5 / 9)     // "100"; (f ‐ 32) * 5 is a float64
	fmt.Println(5 / 9 * (f - 32))     // "0"; 5/9 is an untyped integer, 0
	fmt.Println(5.0 / 9.0 * (f - 32)) // "100"; 5.0/9.0 is an untyped float

	/* 只有常量可以是无类型的。当一个无类型的常量被赋值给一个变量的时候，就像下面的第一行语
	句，或者出现在有明确类型的变量声明的右边，如下面的其余三行语句，无类型的常量将会被隐式
	转换为对应的类型，如果转换合法的话。
	*/
	var g float64 = 3 + 0i // untyped complex ‐> float64
	g = 2                  // untyped integer ‐> float64
	g = 1e123              // untyped floating‐point ‐> float64
	g = 'a'                // untyped rune ‐> float64
	fmt.Println(g)

	/*无论是隐式或显式转换，将一种类型转换为另一种类型都要求目标可以表示原始值。
	对于浮点数和复数，可能会有舍入处理：
	*/
	const (
		deadbeef = 0xdeadbeef        // untyped int with value 3735928559
		a        = uint32(deadbeef)  // uint32 with value 3735928559
		b        = float32(deadbeef) // float32 with value 3735928576 (rounded up)
		c        = float64(deadbeef) // float64 with value 3735928559 (exact)
		// d        = int32(deadbeef)   // compile error: constant overflows int32
		// e        = float64(1e309)    // compile error: constant overflows float64
		// f        = uint(-1)          // compile error: constant underflows uint
	)

	/* 对于一个没有显式类型的变量声明（包括简短变量声明），常量的形式将隐式决定变量的默认类型 */
	i := 0      // untyped integer; implicit int(0)
	r := '\000' // untyped rune; implicit rune('\000')
	t := 0.0    // untyped floating‐point; implicit float64(0.0)
	s := 0i     // untyped complex; implicit complex128(0i)
	fmt.Println(i, r, t, s)

	/* 注意有一点不同：无类型整数常量转换为int，它的内存大小是不确定的，但是无类型浮点数和复数
	常量则转换为内存大小明确的float64和complex128。 如果不知道浮点数类型的内存大小是很难写
	出正确的数值算法的，因此Go语言不存在整型类似的不确定内存大小的浮点数和复数类型。
	*/

	/* 如果要给变量一个不同的类型，我们必须显式地将无类型的常量转化为所需的类型，
	或给声明的变量指定明确的类型
	*/
	var k = int8(0)
	var h int8 = 0
	fmt.Println(k, h)

	/* 当尝试将这些无类型的常量转为一个接口值时（见第7章），这些默认类型将显得尤为重要，因为
	要靠它们明确接口对应的动态类型。
	*/
	fmt.Printf("%T\n", 0)      // "int"
	fmt.Printf("%T\n", 0.0)    // "float64"
	fmt.Printf("%T\n", 0i)     // "complex128"
	fmt.Printf("%T\n", '\000') // "int32" (rune)
}

func main() {
	const_type()
	const_definition()
	flag()
	untype_const()
	untpyed_const_opt()
}
