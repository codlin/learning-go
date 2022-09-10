package main

import (
	"fmt"
	"reflect"
	"unicode/utf8"
	"unsafe"
)

/*
string
一个字符串是一个不可改变的字节序列。一个字符串可以包含任意的数据，包括byte值0.
文本字符串通常被解释为采用UTF8编码的Unicode码点（rune）序列。
*/

/*
内置的函数len可以返回一个字符串的字节数目（不是rune字符数目），
索引s[i]返回第i个z字节的字节值，i必须满足0<=i<len(s)条件约束
*/
func strlen() {
	s := "hello world，你好，世界！" // 每个中文字符占3个字节
	fmt.Printf("%s, len: %d\n", s, len(s))
}

/* 子字符串操作s[i:j]基于原始的s字符串的第i个字节开始到第j个字节（不包含j本身）生成一个新字符串。新生成的字符串将包含j-i个字节 */
func substr() {
	s := "hello world，你好，世界！" // 每个中文字符占3个字节

	fmt.Println(s[:5])
	fmt.Println(s[13:])
	fmt.Println(s[:])

	/* +操作符将两个字符串连接构造一个新的字符串 */
	fmt.Println(s[14:23] + "中国！")
}

/* 字符串可以用==和<=进行比较；比较通过逐个字节比较完成，因此比较的结果是字符串自然编码的顺序*/
func strcmp() {
	s, t := "你好，世界！", "你好，中国！"
	if s < t {
		fmt.Println("s<t")
	} else {
		fmt.Println("s>=t")
	}
}

/* 字符串的值是不可变的：一个字符串包含的字节序列永远不会被改变；当然我们可以给一个字符串变量分配一个新字符串 */
func str_immut() {
	s := "hello, world"
	// s[0] = 'f' // compile error
	s = "foo bar"
	fmt.Println(s)
}

/* 字符串不变性意味着两个字符串可以安全的共享相同的底层数据，这使得复制任何长度的字符串代价是低廉的 */
func str_share() {
	s := "hello, world"
	fmt.Println((*reflect.StringHeader)(unsafe.Pointer(&s)))

	// m 和 s 共享了相同的底层数据
	m := s[:5]
	fmt.Println((*reflect.StringHeader)(unsafe.Pointer(&m)))

	// n 指向 s 内部数据
	n := s[7:]
	fmt.Println((*reflect.StringHeader)(unsafe.Pointer(&n)))

	// t 和 s 具有相同的首地址和长度，即变量t作为指针变量，内部指向了s
	t := s[:]
	fmt.Println((*reflect.StringHeader)(unsafe.Pointer(&t)))

	/* 对字符串切片，生成的仍旧是字符串，不可改变 */
	// m = append(m, "China") // 编译不通过，m 需要为 slice
}

/*
字符串字面值
只需要用英文双引号将字符串包括即可。如"hell, world"。还可以使用反引号包括字符串，这样得到的将是原始字符串，没有转义操作。
因为GO语言的源文件总是用UTF8编码，并且Go语言的文本字符串也可以用UTF8编码的方式处理，因此我们可以将Unicode码点也写到字符串字面值中，如"你好"。
*/
func str_literal() {
	s := "hello world，你好，世界！" // 因采用UTF8编码方式，每个中文字符占3个字节
	fmt.Println(s)
}

/* 可以在双引号包括的字符串中用反斜杠 \ 插入转义字符 */
func escap_sequence() {
	s := "hello world!\n你好，世界！" // 插入了转义序列\n，使字符串分两行打印
	fmt.Println(s)

	/* 可以通过十六进制或八进制转义在字符串字面值里包含任意的字节。八进制数字最大不能超过377 */
	s = "hello \xF7 world \377"
	fmt.Println(s)

	/* 可以使用反引号包括字符串，这样得到的将是原始字符串，没有转义操作 */
	s = `hello \xF7 world \377`
	fmt.Println(s)
}

/*
Unicode & UTF-8

起始，计算机世界还是比较简单的。
世界只有一个字符集：ASCII(American Standard Code for Information Interchange，米国信息交换码)。
它使用7bit来表示128个字符：包括英文字母的大小写、数字、各种标点符号和设备控制符。
后来，世界上很多地区的用户无法用字节的语言符号来使用计算机。

神说，要有光。
于是，Unicode应运而生。它收集了世界上所有的符号系统，每个符号都分配一个唯一的ID，称之为Unicode码点（Unicode Pinpoint）。
UTF-8、UTF-16、UTF-32是把Unicode的码点表示为二进制的方法，这样才能够被计算机识别和理解，我们称之为编码方式。

通用的表示一个Unicode码点的数据类型是int32，可以采用多种编码方式，下面是三种UTF编码方式的简单介绍：
- UTF-8
  把Unicode字符表示为变长的比特，和ASCII码完美兼容。也就是说对于ASCII码，用一个字节表示，其他字符用2个或者更多个字节表示。
  具体编码方式为：
  每个字符编码的第一个字节高位用于指示后面有几个编码字节。
  如果第一个字节的高位为0，则表示对应7bit的ASCII字符，和传统的ASCII编码兼容。
  如果第一个字节的高位时110，则说明需要两个字节；后续的字节高位都以10开始。
  更大的Unicode码点也是采用类似的策略处理。
  0xxxxxxx runes    0‐127                (ASCII)
  110xxxxx 10xxxxxx 128‐2047             (values <128 unused)
  1110xxxx 10xxxxxx 10xxxxxx             2048‐65535 (values <2048 unused)
  11110xxx 10xxxxxx 10xxxxxx 10xxxxxx    65536‐0x10ffff (other values unused)
  变长的编码无法直接通过索引来访问第n个字符，但是UTF-8获得了很多额外的优点：
  1. UTF-8 比较紧凑，占用空间少；
  2. UTF-8 完全兼容ASCII码；
  3. UTF-8 编码可以自动同步：它可以通过向前回溯最多2个字节就可以确定当前字符编码的开始字节的位置（通过前缀）。
  4. UTF-8 是一个前缀编码，所以当从左往右解码时不会有任何歧义，也不需要向前查看。
  5. UTF-8 没有任何字符编码是其它字符编码的子串，或其它序列的子串，因此字符串搜索时只需要搜索它的字节编码序列即可。不需要担心前后的上下文会对搜索结果造成干扰。
  6. UTF-8 的编码顺序和Unicode的码点序列一致，因此可以直接排序UTF-8的编码序列（码点可以直接按照大小写入上面的xxx中，所以在排序时得到的顺序也和原字符一致）。
  7. UTF-8 因为没有嵌入NUL(0)字节，所以可以很好的兼容那些使用NUL作为字符串结尾的编程语言。
- UTF-16
  常用与存储空间和获取字符效率需要取得平衡的场景。常用的字符用16bit（一个code unit）编码，其他的用一对code unit编码(pairs of 16-bit code units)。
- UTF-32
  用户存储空间不是问题，需要等宽的code unit的场景。每一个Unicode字符被表示为32bit。

Uincode码点对应Go语言中的rune整数类型，rune（符文）也是int32等价类型。

Go语言源文件采用UTF-8编码，并且处理UTF-8编码的文本也很出色。
unicode包通过了很多处理rune字符相关的函数，unicode/utf8包则提供了用于rune字符序列的UTF-8编码和解码的功能。

Go语言字符串字面值中的Unicode转义字符让我们可以通过Unicode码点输入特殊的字符。
有两种形式：\uhhhh对应16bit的码点值，\Uhhhhhhhh对应32bit的码点值，其中h是一个十六进制数字；
一般很少需要使用32bit的形式。上面两种每一种都是指定码点的编码形式。
例如：下面的字母串面值都表示相同的值：
"世界"
"\xe4\xb8\x96\xe7\x95\x8c"
"\u4e16\u754c"
"\U00004e16\U0000754c"
上面三个转义序列都为第一个字符串提供替代写法，但是它们的值都是相同的，其中第一种是UTF8编码后的字节序列。

Unicode转义也可以使用在rune字符中。下面三个字符是等价的：
'世界' '\u4e16界' '\U00004e16界'
*/

func unicode_utf8() {
	s := "hello, 世界"
	fmt.Printf("s length: %d\n", len(s))
	for i := 0; i < len(s); i++ {
		fmt.Printf("%d, %[2]q, %[2]b\n", i, s[i])
	}

	fmt.Printf("s rune length: %d\n", utf8.RuneCountInString(s))
	for i, v := range s { /* for range 在处理字符串的时候，会自动隐式解码UTF8字符串 */
		fmt.Printf("%d %q\n", i, v)
	}

	// for 循环中采用解码的方式遍历所有 rune。这段程序和 for range 等价
	for i := 0; i < len(s); {
		// DecodeRuneInString 返回解码后的 rune 符文以及该符文的编码长度
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}

	count := 0
	for range s {
		count++
	}
	fmt.Printf("s length: %d\n", count)

	/* 每一个UTF8字符解码，不管是显式的调用 utf8.RuneCountInString 解码或是在range循环中隐式解码，如果遇到一个错误的UTF8编码输入，
	将生成一个特别的Unicode字符'\uFFFD', 在打印时这个字符为一个菱形里面有个问号
	*/
	s = "hello, 世\xF7"
	fmt.Println(s)

	/* UTF8 字符串作为交换格式是非常方便的，但是在程序内部使用rune序列可能更方便，因为rune大小一致，可以支持数组索引和方便切割
	string 接收到 []rune 的类型转换，可以将一个UTF8编码的字符串解码为Unicode字符序列
	*/
	s = "プログラム"
	fmt.Printf("% x\n", s) // "e3 83 97 e3 83 ad e3 82 b0 e3 83 a9 e3 83 a0"
	r := []rune(s)
	fmt.Printf("%x\n", r) // "[30d7 30ed 30b0 30e9 30e0]"

	/* 如果将一个 []rune 类型的Unicode字符slice或数组转换为string，则对它们进行UTF8编码 */
	fmt.Println(string(r))

	/* 将一个整型转换为字符串的意思是生成只包含对应Unicode码点字符的UTF8字符串 */
	fmt.Println(string(65))
	fmt.Println(string(0x4eac))

	/* 如果对应码点的字符串是无效的，则用"\uFFFD"无效字符作为替换 */
	fmt.Println(string(12345678))
}

func str_slice() {

}

func main() {
	strlen()
	substr()
	strcmp()
	str_share()
	escap_sequence()
	unicode_utf8()
}
