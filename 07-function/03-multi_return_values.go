package main

/*
函数的多返回值

调用多返回值函数时，返回给调用者的是一组值，调用者必须显示的将这些值分配给变量：
links, err := findLinks(url)
如果某个值不被使用，可以分配给blank identifier：
links, _ := findLinks(url)

一个函数内部可以将另一个有多返回值的函数作为返回值返回。
func findLinksLog(url string) ([]string, error) {
	log.Printf("findLinks %s", url)
	return findLinks(url)
}

当你调用接受多个参数的函数时，可以将一个返回多参数的函数作为该函数的参数。虽然这很少出现在实际的生产代码中，但是
这个特性在debug时很方便。
log.Println(findLinks(url))

准确的命名可以传达函数返回值的含义，尤其在返回值类型都相同时。
func Size(rect image.Rectangle) (width, height int)
按照惯例，函数的最后一个bool返回值代表函数是否运行成功， error类型的返回值代表函数的错误信息，对于这些类似的惯例，
我们无需思考合适的命名，它们都无需解释。
func Size(rect image.Rectangle) (width, height int, err error)

如果一下函数的所有返回值都有显式的变量名，那么该函数的return函数的返回语句可以省略操作数。这称之为bare return。
func Size(rect image.Rectangle) (width, height int, err error) {
	width = 2
	height = 4
	err = nil
	return
}
当一个函数有多处return语句以及很多返回值时，bare语句可以减少代码的重复，但也使得代码难以被理解。
基于以上原因，不宜过度使用bare return。
*/
