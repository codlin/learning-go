package main

import (
	"fmt"
	"math"
	"strings"
)

/*
基于指针的接收器

当调用一个函数时，会对每一个参数进行值拷贝，如果一个函数需要更新一个变量，或者一个函数的参数实在太大我们希望避免这种默认的拷贝，
这种情况下我们就需要用到指针。
对应到这里用来更新接收器对象的方法，当这个接收者变量本身比较大时，我们就可以用其指针而不是对象声明方法。
*/
type Point struct {
	X, Y float64
}

// method of the Point type
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// method of the Point type, use *Point as receiver
func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

/*
在现实的程序里，一般会约定，如果类型有一个方法使用了指针作为接收器，那么所有的方法都必须使用指针作为接收器。
即使那些并不需要指针作为接收器的函数。

只有类型或只向它们的指针，才能可能是会出现在接收器声明里的两种接收器。
在声明一个方法是，如果一个类型名本身是一个指针的话，是不允许出现在接收器中的。
type P *int

func (p P) f() {} // invalid receiver type P (pointer or interface type)
*/

/*
如果类型T的所有方法都是用T类型自己来做接收器（而不是*T），那么copy这种类型的实例就是安全的；调用它的任何一个方法就会产生一个值的拷贝。
（注：此处我认为不是绝对安全，如果结构体中有指针，如果拷贝了这个变量，那么它们指向同一块内存。(x)
更新：上面我的认识是不对的
即使结构体里面的字段是指针类型，在变量调用方法时，因为形参receiver是值，所以会从调用方法的变量拷贝出一些新变量作为receiver，
对于这个新变量来说，结构体内的指针成员也会指向新的内存，即编译器/运行时做了深拷贝，详细见后面例子。
但是，对于外部变量来说，如果通过值复制的方式生成了新变量，修改新变量的指针指向的内容的值，会影响原有变量，这和C语言是一致的。）

如果一个方法使用指针作为接收器，那么在调用方法时，实际上接收器指向的是同一个指针，因为我们应该避免对其拷贝，以免产生错误。
*/
type watch struct {
	brand  string
	model  string
	weight int
}

func (w watch) show() {
	fmt.Printf("type pointer: %p, brand: %s, model:%s, weight: %d\n", &w, w.brand, w.model, w.weight)
}

type phone struct {
	brand  string
	model  string
	weight int
}

func (p *phone) show() {
	fmt.Printf("in show: type pointer: %p, brand: %s, model:%s, weight: %d\n", p, p.brand, p.model, p.weight)
}

func (p *phone) update(weight int) {
	fmt.Printf("modify weight in method: %d\n", weight)
	p.weight = weight
	fmt.Printf("in update: type pointer: %p, brand: %s, model:%s, weight: %d\n", p, p.brand, p.model, p.weight)
}

func (p *phone) update_mode(model string) {
	fmt.Printf("modify model in method: %s\n", model)
	p.model = model
	fmt.Printf("in update_mode: type pointer: %p, brand: %s, model:%s, weight: %d\n", p, p.brand, p.model, p.weight)
}

type pad struct {
	brand  string
	model  string
	weight *int
}

func (p pad) update(weight int) {
	fmt.Printf("modify weight in method: %d\n", weight)
	p.weight = &weight
	fmt.Printf("in update: type pointer: %p, brand: %s, model:%s, weight[0x%p]: %d\n", &p, p.brand, p.model, p.weight, *p.weight)
}

func (p pad) show() {
	fmt.Printf("in show: %p\tbrand:%s\tmodel:%s\tweight[0x%p]:%d\n", &p, p.brand, p.model, p.weight, *p.weight)
}

func show_pad(p *pad) {
	fmt.Printf("out show: %p\tbrand:%s\tmodel:%s\tweight[0x%p]:%d\n", p, p.brand, p.model, p.weight, *p.weight)
}

type notebook struct {
	brand  string
	model  string
	weight *int
}

func (n *notebook) update(weight int) {
	fmt.Printf("modify weight in method: %d\n", weight)
	n.weight = &weight
	fmt.Printf("in update, type pointer: %p, brand: %s, model:%s, weight: %d\n", n, n.brand, n.model, *n.weight)
}

func (n *notebook) show() {
	fmt.Printf("in show: %p\tbrand:%s\tmodel:%s\tweight[0x%p]:%d\n", n, n.brand, n.model, n.weight, *n.weight)
}

func show_notebook(n *notebook) {
	fmt.Printf("out show: %p\tbrand:%s\tmodel:%s\tweight[0x%p]:%d\n", n, n.brand, n.model, n.weight, *n.weight)
}

type itv struct {
	brand  string
	model  string
	weight *int
}

func (n *itv) update_weight(weight int) {
	fmt.Printf("modify weight in method: %d\n", weight)
	n.weight = &weight
	fmt.Printf("in update_weight, type pointer: %p, brand: %s, model:%s, weight: %d\n", n, n.brand, n.model, *n.weight)
}
func (n itv) update_model(model string) {
	fmt.Printf("modify model in method: %s\n", model)
	n.model = model
	fmt.Printf("in update_model, type pointer: %p, brand: %s, model:%s, weight: %d\n", &n, n.brand, n.model, *n.weight)
}

func (n *itv) show() {
	fmt.Printf("in show: %p\tbrand:%s\tmodel:%s\tweight[0x%p]:%d\n", n, n.brand, n.model, n.weight, *n.weight)
}

func show_itv(n *itv) {
	fmt.Printf("out show: %p\tbrand:%s\tmodel:%s\tweight[0x%p]:%d\n", n, n.brand, n.model, n.weight, *n.weight)
}

// 测试类型的方法接收器是值、指针以及混合情况下的初始化方式
/*
在每一个合法的方法调用表达式中，下面三种情况的任意一种都是可以的：
1. 实际变量的类型和接收器形参的类型相同
2. 接收器形参类型是T，但接收器实参类型是*T，这种情况下编译器会隐式的为我们解引用，取到指针指向的实际变量。
   如func (w watch) show()，对于指针变量 pwatch，pwatch.show() 相当于 (*pwatch).show()
3. 接收器形参类型是*T，但接收器实参类型是T，这种情况下编译器会隐式的为我们取变量的地址。
   如func (p *phone) show()，对于变量 iphone6s， iphone6s.show() 相当于 (&iphone6s).show()
*/
func test_init() {
	fmt.Printf("%s\n", strings.Repeat("=", 64))
	fmt.Println("情景1： 方法接收器为值类型，通过值类型的变量来调用方法")
	watch3 := watch{brand: "iwatch", model: "v3", weight: 40}
	watch3.show()

	fmt.Println("情景2： 方法接收器为值类型，通过指针类型的变量来调用方法")
	watch4 := &watch{brand: "iwatch", model: "v4", weight: 45}
	watch4.show()
	fmt.Println("watch4.show()这种情况，编译器会自动对指针变量解引用后调用方法，相当于：(*watch4).show()")
	(*watch4).show()
	fmt.Println("对于方法接收器为值类型，我们可以通过直接声明变量的方式来调用方法，\n如：watch{brand: \"iwatch\", model: \"v7\", weight: 60}.show()")
	watch{brand: "iwatch", model: "v7", weight: 60}.show()

	fmt.Printf("%s\n", strings.Repeat("=", 64))
	fmt.Println("情景3： 方法接收器为指针类型，通过值类型的变量来调用方法")
	iphone6s := phone{brand: "iphone", model: "6s", weight: 180}
	iphone6s.show()
	fmt.Println("iphone6s.show()这种情况，编译器会自动取变量地址后调用方法，相当于：(&iphone6s).show()")
	(&iphone6s).show()

	fmt.Println("情景4： 方法接收器为指针类型，通过指针类型的变量来调用方法")
	iphone6sp := &phone{brand: "iphone", model: "6sp", weight: 200}
	iphone6sp.show()
	fmt.Println("(x)对于方法接收器为指针类型，我们**无法**通过直接声明变量取地址的方式来调用方法，\n如：&phone{brand: \"iphone\", model: \"6sp\", weight: 200}.show()")

	fmt.Printf("%s\n", strings.Repeat("=", 64))
	fmt.Println("情景5： 方法接收器为既有值也有指针的混合类型，通过值类型的变量来调用方法")
	weight := 800
	itv_3 := itv{brand: "iTV", model: "3", weight: &weight}
	itv_3.update_weight(805) // 相当于 (&itv_3).update_weight(805)
	itv_3.update_model("4")
	itv_3.show() // 相当于 (&itv_3).show()

	fmt.Printf("%s\n", strings.Repeat("=", 64))
	fmt.Println("情景5： 方法接收器为既有值也有指针的混合类型，通过指针类型的变量来调用方法")
	weight = 810
	itv_4 := &itv{brand: "iTV", model: "4", weight: &weight}
	itv_4.update_weight(815)
	itv_4.update_model("4") // 相当于 （*itv_4）.update_model("4")
	itv_4.show()
}

func test_copy() {
	iwatch := watch{brand: "iwatch", model: "v4", weight: 40}
	fmt.Printf("%p\t%v\n", &iwatch, iwatch)
	iwatch.show() // type pointer: 0xc000100510

	fmt.Printf("%s\n", strings.Repeat("=", 64))
	fmt.Println("情景1： 方法接收器为值类型，变量通过拷贝值生成")
	another_iwatch := iwatch
	another_iwatch.weight = 50
	another_iwatch.model = "v5"
	fmt.Printf("%p\t%v\n", &another_iwatch, another_iwatch)
	another_iwatch.show()

	/* 哪怕是同一个变量，如果方法的receiver是值，在调用方法时也会先拷贝一份变量值，然后再调用方法，因此上面打印显示类型的指针不同 */
	fmt.Println("show original var")
	iwatch.show() // type pointer: 0xc000100570
	fmt.Println("哪怕是同一个变量，如果方法的receiver是值，在调用方法时也会先拷贝一份变量值，然后再调用方法，因此上面打印显示类型的指针不同")

	fmt.Printf("%s\n", strings.Repeat("-", 64))
	fmt.Println("情景2： 方法接收器为值类型，变量通过拷贝指针生成")
	third_iwatch := &iwatch
	third_iwatch.weight = 60
	third_iwatch.model = "v6"
	third_iwatch.show()

	fmt.Println("show original var")
	iwatch.show()

	/* 如果方法的receiver是指针变量，则不产生拷贝。打印变量的指针，和方法中打印变量的指针相同。*/
	fmt.Printf("%s\n", strings.Repeat("-", 64))
	iphone_6s := phone{brand: "iphone", model: "6s", weight: 180}
	fmt.Printf("%p\t%v\n", &iphone_6s, iphone_6s)
	iphone_6s.update(182)
	iphone_6s.update_mode("6")
	iphone_6s.show()
	fmt.Println("如果方法的receiver是指针变量，则不产生拷贝。上面打印变量的指针，和方法中打印变量的指针相同。")

	fmt.Printf("%s\n", strings.Repeat("=", 64))
	fmt.Println("情景3： 方法接收器为指针类型，变量通过拷贝指针生成")
	iphone_6sp := &iphone_6s
	iphone_6sp.model = "6sp"
	iphone_6sp.weight = 200
	fmt.Printf("%p\t%v\n", &iphone_6sp, iphone_6sp)
	iphone_6sp.show()
	/* 上面修改的是同一个变量，因为show方法的接收者是指针，因此在调用时不会产生拷贝，打印出的类型指针和iphone_6s相同 */

	fmt.Println("show original var")
	iphone_6s.show()
	fmt.Println("上面修改的是同一个变量，因为show方法的接收者是指针，因此在调用时不会产生拷贝，打印出的类型指针和iphone_6s相同")

	fmt.Printf("%s\n", strings.Repeat("-", 64))
	fmt.Println("情景4： 方法接收器为指针类型，变量通过拷贝值生成")
	iphone_x := iphone_6s
	iphone_x.model = "x"
	iphone_x.weight = 220
	fmt.Printf("%p\t%v\n", &iphone_x, iphone_x)
	iphone_x.show()
	/* 变量iphone_x对iphone_6s进行了值copy，因此产生了新的变量，所以调用方法的类型指针也发生了变化，但是和新变量的指针一样 */
	fmt.Println("变量iphone_x对iphone_6s进行了值copy，因此产生了新的变量，所以调用方法的类型指针也发生了变化，但是和新变量的指针一样")

	/* 不会影响原变量的值 */
	fmt.Println("show original var")
	iphone_6s.show()

	fmt.Printf("%s\n", strings.Repeat("-", 64))
	weight := 400
	ipad_air2 := pad{brand: "ipad", model: "air2", weight: &weight}
	show_pad(&ipad_air2)
	ipad_air2.show()

	fmt.Printf("%s\n", strings.Repeat("=", 64))
	fmt.Println("情景5： 方法接收器为值类型，变量通过拷贝值生成。变量中存在指针类型的字段")
	ipad_air3 := ipad_air2
	ipad_air3.model = "ari3"
	weight = 405
	ipad_air3.weight = &weight
	show_pad(&ipad_air3)
	ipad_air3.update(406)
	ipad_air3.show()
	show_pad(&ipad_air3)
	fmt.Printf("show ipad_air2: ")
	show_pad(&ipad_air2)
	/* ↑因为update方法的接收者是值，所以在调用时会拷贝一个变量，因此在方法内修改weight的值不会影响到外面。
	但是外面的变量ipad_air3修改了weight指针的值，ipad_air2和ipad_air2底层指向同一个对象，所以ipad_air2的weight也被修改了。
	*/
	fmt.Println("↑因为update方法的接收者是值，所以在调用时会拷贝一个变量，因此在变量内修改weight的值不会影响到外面。\n但是外面的变量ipad_air3修改了weight指针的值，ipad_air2和ipad_air2底层指向同一个对象，所以ipad_air2的weight也被修改了。")

	fmt.Printf("%s\n", strings.Repeat("-", 64))
	fmt.Println("情景6： 方法接收器为值类型，变量通过拷贝指针生成。变量中存在指针类型的字段")
	ipad_air4 := &ipad_air2
	ipad_air4.model = "ari4"
	weight = 410
	ipad_air4.weight = &weight
	show_pad(ipad_air4)
	ipad_air4.update(415)
	ipad_air4.show()
	show_pad(ipad_air4)
	fmt.Printf("show ipad_air2: ")
	show_pad(&ipad_air2)
	/* 外部复制指针变量，对新变量的修改行为和C语言一致，因为底层指向的是同一个对象 */
	fmt.Println("外部复制指针变量，对新变量的修改行为和C语言一致，因为底层指向的是同一个对象")

	fmt.Printf("%s\n", strings.Repeat("-", 64))
	weight = 1000
	notebook_air := notebook{brand: "iMac", model: "bookair", weight: &weight}
	show_notebook(&notebook_air)
	notebook_air.update(1401)
	notebook_air.show()
	show_notebook(&notebook_air)

	fmt.Printf("%s\n", strings.Repeat("=", 64))
	fmt.Println("情景7： 方法接收器为指针类型，变量通过拷贝值生成。变量中存在指针类型的字段")
	notebook_air2 := notebook_air
	notebook_air2.model = "ari2"
	weight = 1405
	notebook_air2.weight = &weight
	show_notebook(&notebook_air2)
	notebook_air2.update(1406)
	notebook_air2.show()
	show_notebook(&notebook_air2)

	fmt.Printf("%s\n", strings.Repeat("-", 64))
	fmt.Println("情景8： 方法接收器为指针类型，变量通过拷贝指针生成。变量中存在指针类型的字段")
	notebook_pro := notebook_air
	notebook_pro.model = "pro"
	weight = 1410
	notebook_pro.weight = &weight
	show_notebook(&notebook_pro)
	notebook_pro.update(1415)
	notebook_pro.show()
	show_notebook(&notebook_pro)

	fmt.Printf("%s\n", strings.Repeat("=", 64))
	fmt.Println("情景9： 类型的多个方法接收器有值类型和指针类型的混合形式")
	fmt.Printf("%s\n", strings.Repeat("-", 64))
	weight = 800
	itv_3 := itv{brand: "iTV", model: "3", weight: &weight}
	show_itv(&itv_3)
	itv_3.update_weight(850)
	itv_3.update_model("4")
	itv_3.show()     // model还是3
	show_itv(&itv_3) // model还是3
	fmt.Println("由上面的例子可见，尽量不要在类型的各种方法里混合Mix不同的接收器类型，这样会有意想不到的效果。")
}

func main() {
	test_init()
	test_copy()
}
