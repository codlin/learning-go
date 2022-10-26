package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

/*
基于select的多路复用（Multiplexing with Select）

当我们需要监听多个channel时，我们无法循环去轮询，因为当我们在某个channel上接收时，
如果该channel中没有数据，那么程序就会阻塞，这样我们就无法收到第二个channel上发送的数据。
这时我们需要使用多路复用（multiplex），类似于linux里的多路复用select，Go语言使用select-case提供多路复用机制。
```go

	select {
	    case <‐ch1:
	    // ...
	    case x := <‐ch2:
	    // ...use x...
	    case ch3 <‐ y:
	    // ...
	    default:
	    // ...
	}

```
select和switch语言有点类似，也有几个case和default分支。每一个case代表一个通信操作（在某个channel上发送或者接收），
并且会包含一些语句组成一个语句块。一个接收表达式可能只包含接收表达式自身（即不把接收的值赋给变量），就像上面第一个case；
或者包含在一个**简短的变量声明**中，像第二个case里一样；第二种形式让我们可以引用到接收到的值。

select会等待所有case中有能够执行的case时去执行。当条件满足时select才会去通信并执行case之后的语句；这时候其它通信时不会被执行的。
一个没有任何case的select语句写作select{}，会永远等待下去。
*/
func rocket_lunch() {
	fmt.Printf("%s\r", "rocket lunching")
}

// ！！！该程序有问题，会造成goroutine泄露！！！
func rocket_countdown1() {
	fmt.Println("Commencing countdown.")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Printf("%2d\r", countdown)
		<-tick
	}
	rocket_lunch()
}

// 增加倒计时过程中用户按下return键时中断发射流程的功能
// ！！！该程序有问题，会造成goroutine泄露！！！
func rocket_countdown2() {
	fmt.Println("Commencing countdown. Press return to abort")
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		select {
		case <-tick:
			fmt.Printf("%-2d\r", countdown)
		case <-abort:
			fmt.Println("lunching was abort")
			return
		}
	}
	rocket_lunch()

}

/*
time.Tick会返回一个<- time.Time的接收类型channel，我们无法在接收的goroutine中关闭它。
上面的两个rocket_countdown例子中，该函数表现的好像它创建了一个在循环中调用time.Sleep的goroutine，
每次被唤醒时发送一个事件。当countdown函数返回时，它会停止从channel中接收事件，但是tick这个gorouteine依然存活，
继续徒劳的尝试向channel发送值，然而这时候已没有其它的goroutine会从该channel中接收值了---因此造成了goroutine泄露。

time.Tick挺方便，但是只有当程序整个生命周期都需要这个事件时我们用它才比较合适。否则的话我们应该使用time.NewTicker
*/
func rocket_countdown3() {
	fmt.Println("Commencing countdown. Press return to abort")
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	tick := time.NewTicker(1 * time.Second)
	defer tick.Stop() // 返回去关闭channel

	for countdown := 10; countdown > 0; countdown-- {
		select {
		case <-tick.C:
			fmt.Printf("%-2d\r", countdown)
		case <-abort:
			fmt.Println("lunching was abort")
			return
		}
	}
	rocket_lunch()

}

/*
channel的零值时nil。nil的channel有时候也是有一些用处的。因为对一个nil的channel发送或接收操作会永远阻塞，
在select中操作nil的channel将永远都不会被select到。这使得我们可以用nil来激活或禁用case，来达成处理其它输入或
输出事件时超时或取消的逻辑。具体请参考例子du3()
*/
// walkDir 以给定的目录作为根目录递归遍历文件树，并且发送每个文件的大小给fileSizes
func walkDir(dir string, fileSizes chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

// dirents 返回目录下的条目信息
func dirents(dir string) []os.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du: %v\n", err)
		return nil
	}
	return entries
}

func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files %.1f MB\n", nfiles, float64(nbytes)/1e6)
}

// 计算给定目录的大小
func du() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	fileSizes := make(chan int64)
	go func() {
		for _, root := range roots {
			walkDir(root, fileSizes)
		}
		close(fileSizes)
	}()

	var nfiles, nbytes int64
	for size := range fileSizes {
		nfiles++
		nbytes += size
	}
	printDiskUsage(nfiles, nbytes)
}
func main() {
	// rocket_countdown3()

	du()
}
