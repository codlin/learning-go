package main

func pointer_var() {
	x := 1

	// &x表达式可以取变量x的内存地址
	p := &x
	print(p)

	// 使用*p可以读取指针变量p指向的内容
	print(*p)

	// *p表达式可以出现在赋值语句的左边，表示更新指针所指向内存变量的值
	*p = 2
	print(2)
}

/*
在GO中，函数返回局部变量的指针也是安全的。（函数逃逸，变量将会分配在堆上）
*/
func pointer_f() *int {
	v := 1
	return &v
}

func main() {
	pointer_var()
	p := pointer_f()
	print(p)
}
