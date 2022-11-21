package main

import (
	"fmt"
	"unsafe"
)

func side_effort_1() {
	a := []byte{'h', 'e', 'l', 'l', 'o'}
	b := a[2:3]
	fmt.Println(string(a), string(b)) // 输出 hello l

	b = append(b, 'w')
	fmt.Println(string(a), string(b)) // 输出 helwo lw

	// 避免切片副作用的一个方法是在reslice时指定max的边界
	c := a[2:3:3]
	c = append(c, 'r')
	fmt.Println(string(a), string(c)) // 输出 helwo lr

	// 避免切片副作用的一个方法是使用 copy 函数
	d := make([]byte, 1)
	copy(d, a[2:3])
	d = append(d, 'd')
	fmt.Println(string(a), string(d)) // 输出 helwo ld
}

func side_effort_2() {
	// 指针类型变量引用切片产生副作用
	type sliceHeader struct {
		array unsafe.Pointer // 底层数据数组的指针
		len   int            // 切片长度
		cap   int            // 切片容量
	}

	type User struct {
		Likes int
	}

	users := make([]User, 1)

	ptrUsers := (*sliceHeader)(unsafe.Pointer(&users))
	fmt.Printf("切片%s: 底层数组地址=0x%x, 长度=%d, 容量=%d\n", "a", ptrUsers.array, ptrUsers.len, ptrUsers.cap)

	pFirstUser := &users[0]
	pFirstUser.Likes++
	fmt.Println("所有用户：")
	for i := range users {
		fmt.Printf("User: %d Likes: %d\n\n", i, users[i].Likes)
	}

	users = append(users, User{})
	ptrUsersNew := (*sliceHeader)(unsafe.Pointer(&users))
	fmt.Printf("切片%s: 底层数组地址=0x%x, 长度=%d, 容量=%d\n", "a", ptrUsersNew.array, ptrUsersNew.len, ptrUsersNew.cap)

	// 此处 pFirstUser.Likes++ 对第一个用户的加1操作失效，这是因为 users 经过 append 后，底层数组指针以及发生了改变，
	// 所以对原来的指针变量操作，相当于访问了脏数据
	pFirstUser.Likes++
	fmt.Println("所有用户：")
	for i := range users {
		fmt.Printf("User: %d Likes: %d\n\n", i, users[i].Likes)
	}
}

func main() {
	side_effort_1()

	side_effort_2()

	/**
	 * 避免切片副作用黄金法则
	 * 黄金法则1：在边界处拷贝切片，这里面的边界指的是函数接受切片参数或返回切片的时候。
	 * 黄金法则2：永远不要使用一个变量来引用切片数据
	**/
}
