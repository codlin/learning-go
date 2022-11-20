package main

/*
包的导入

包的导入类型：
1. 直接导入
2. 分组导入
3. 嵌套导入
4. 别名导入
5. 点导入
6. 空白导入
7. 绝对路径导入
8. 相对路径导入
9. 循环导入 （错误）
*/

/*
直接导入

单行导入：
import "fmt"

上面fmt是Go语言的标准库，他其实是去$GOROOT下去加载该模块

多行导入：
import "fmt"
import "math"
*/

/*
分组导入

import (
	"fmt"
	"math"
)

分组导入中间也可以插入空行，以对不同类型的包隔开
import (
	fmt"
	"math"

	"third-part/package/a"

	"my/local/package/b"
)
*/

/*
嵌套导入

指的是导入包的路径嵌套，即可以导入一个路径下的包以及子路径下的包

import (
	"my/loca/package"
	"my/loca/package/subpackge/a"
	"my/loca/package/subpackge/b"
)
*/

/*
别名导入

当导入的包名字很长或者和其它包有冲突时，可以为包指定别名予以区别。

import (
	crypto/rand"
	mrand "math/rand"

	pb "my/loca/package/protobuffer"
)
*/

/*
点导入 Dot Import

import (
	. "io"
)

点导入主要是用于测试目的，测试者用这种类型的导入只要是用来测试它们导出的数据结构，函数，包等功能是否正确。
这种导入方式的好处就是不用写报名就可以直接使用包中导出的元素。相反，带来的坏处就是可能造成命名空间冲突。
*/

/*
空白导入

如果导入的名字没有在文件中引用，就会产生一个编译错误。
但有时候我们必须导入一个包，这仅仅是为了利益其副作用：对包级别的变量进行初始化表达式求值，并执行它的init函数。
为了防止“未使用的导入“错误，我们必须使用一个重命名导入，它使用一个替代的名字_，这表示导入的内容未空白标识符。
通常情况下，空白标识符不可能被引用。
import _ "image/png"
*/

/*
相对路径导入

import   "./test_model"  //当前文件同一目录的test_model目录，但是不建议这种方式import
*/
/*
绝对路径导入

import "mygoproject/test_model"  //加载$GOPATH/src/mygoproject/test_model模块
*/

/*
循环导入

这种方式在Go语言里会导致编译错误。
在A包里导入了B包，在B包里导入了A包；
或者是间接相互引用，在A包里导入了B包，在B包里导入了C包，C包里面又引入了A包。
*/
