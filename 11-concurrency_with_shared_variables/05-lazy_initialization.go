package main

import (
	"fmt"
	"image"
	"sync"
)

/*
延迟初始化： sync.Once

延迟一个昂贵的初始化步骤到有实际需要的时刻是一个很好的实践。
预先初始化一个变量会增加程序的启动时延，并且实际执行时有可能根本用不到这个变量，那么初始化也不是必须的。
*/

var icons map[string]image.Image

func loadIcons() {
	icons = map[string]image.Image{
		"spades.png":   loadIcon("spades.png"),
		"hearts.png":   loadIcon("hearts.png"),
		"diamonds.png": loadIcon("diamonds.png"),
		"clubs.png":    loadIcon("clubs.png"),
	}
}

func loadIcon(name string) image.Image {
	fmt.Printf("%s loaded\n", name)
	return nil
}

// 注意： 并发不安全
func Icon(name string) image.Image {
	if icons == nil {
		loadIcons() // 一次性初始化
	}
	return icons[name]
}

/*
上面的例子，在并发调用时是不安全的。
直觉告诉我们，竞态带来的最严重的问题可能就是loadIcons函数会被调用多次。
但这个直觉仍然是错误的（关于并发的直觉都不可靠）。
但这个直觉仍然是错误的（关于并发的直觉都不可靠）。
但这个直觉仍然是错误的（关于并发的直觉都不可靠）。
在缺乏显示同步的情况下，编译器和CPU在保证每个goroutine都满足串行一致的基础上可以自由地重排内存的访问顺序。
它在填充数据之前把一个空map赋值给icons。
*/
func loadIcons2() { // 编译器的可能处理策略
	icons = make(map[string]image.Image)
	icons["spades.png"] = loadIcon("spades.png")
	icons["hearts.png"] = loadIcon("hearts.png")
	icons["diamonds.png"] = loadIcon("diamonds.png")
	icons["clubs.png"] = loadIcon("clubs.png")
}

/*
因此，一个goroutine发现icons不为nil并不意味着变量的初始化肯定已经完成。
保证所有goroutine都能观察到locaIcons效果最简单的正确方法就是用一个互斥锁来同步。
*/
var mu sync.Mutex

// 并发安全
func Icon2(name string) image.Image {
	mu.Lock()
	defer mu.Unlock()
	if icons == nil {
		loadIcons()
	}
	return icons[name]
}

/*
采用互斥锁访问icons的代价是两个goroutine不能并发访问这个变量，即使在这个变量已经安全完成初始化且不再更改的情况下，也会造成这个结果。
使用一个可以并发读的锁就可以改善这个问题：
*/
var rwmu sync.RWMutex

func Icon3(name string) image.Image {
	rwmu.RLock()
	if icons != nil {
		icon := icons[name]
		rwmu.RUnlock()
		return icon
	}
	rwmu.RUnlock()

	// 获取互斥锁
	mu.Lock()
	if icons != nil {
		loadIcons()
	}
	icon := icons[name]
	mu.Unlock()
	return icon
}

/*
上面的例子中，有两个临界区。首先获取一个读锁，查阅map，然后释放这个读锁。如果条目能找到，就返回它。如果没找到，就再获取一个写锁。
由于不先释放一个共享锁就无法直接把它升级到互斥锁，为了避免在过渡期其它的goroutine已经初始化了icons，所以必须重新检查nil值。
这相当于是并发中的双检测锁定模式的一种。
上面的模式具有更好的并发性，但是更加复杂容易出错。幸好的是，
sync包提供了针对一次性初始化问题的特化解决方案：sync.Once
从概念上讲，Once包含一个布尔变量和一个互斥量，布尔变量记录初始化是否完成，互斥量则保护这个布尔变量和客户端的数据结构。
Once的唯一方法Do以初始化函数作为它的参数。每次调用Do时会先锁定互斥量并检查里面的布尔变量。
在第一次调用时，这个布尔变量为false，Do会调用loadIcons后把它设置为true。
后续的调用相当于空操作，只是通过变量的同步来保证loadIcons对内存产生的效果对所有goroutines可见。
以这种方式来使用sync.Once，可以避免变量在正确构造之前就被其他goroutine分享。
（感觉新版本的sync.Once应该是有所改变了）
*/
var loadIconOnce sync.Once

func Icon4(name string) image.Image {
	loadIconOnce.Do(loadIcons)
	return icons[name]
}
