package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"
)

/*
包含defer的函数
*/

/*
下面的函数title在请求http后，需要在后续的处理中关闭resp.Body资源。
resp.Body.close被调用了多次，这是为了确保title在所有执行路径下（即使函数运行失败）都关闭了
网络连接。随着函数变得复杂，需要处理的错误也变多，维护清理逻辑变得越来越困难。
而Go语言独有的defer机制可以让事情变得简单。
*/
func title(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	// Check Content‐Type is HTML (e.g., "text/html;charset=utf‐8").
	ct := resp.Header.Get("Content‐Type")
	if ct != "text/html" && !strings.HasPrefix(ct, "text/html;") {
		resp.Body.Close()
		return fmt.Errorf("%s has type %s, not text/html", url, ct)
	}
	// doc, err := html.Parse(resp.Body)
	// resp.Body.Close()
	// if err != nil {
	// 	return fmt.Errorf("parsing %s as HTML: %v", url, err)
	// }
	return nil
}

/*
通过defer机制可以在函数返回前执行定义的操作，譬如释放占有的资源等。
这样不必在每个返回错误的分支条件中释放资源。
只需要在调用普通函数或方法前加上关键字defer，就完成了defer所需要的语法。
当defer语句被执行时，跟在defer后面的函数会被延迟执行。
直到包含该defer语句的函数执行完毕时，defer后的函数才会被执行，不论包含defer语句的函数是通过return正常结束，还是由于panic导致的异常结
束。
你可以在一个函数中执行多条defer语句，它们的执行顺序与声明顺序相反。

defer语句经常被用于处理成对的操作，如打开、关闭、连接、断开连接、加锁、释放锁。
通过defer机制，不论函数逻辑多复杂，都能保证在任何执行路径下，资源被释放。
释放资源的defer应该直接跟在请求资源的语句后。需要注意一点：不能忘记defer语句后面的圆括号。
*/
func title_2(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	ct := resp.Header.Get("Content‐Type")
	if ct != "text/html" && !strings.HasPrefix(ct, "text/html;") {
		return fmt.Errorf("%s has type %s, not text/html", url, ct)
	}
	// doc, err := html.Parse(resp.Body)
	// if err != nil {
	// 	return fmt.Errorf("parsing %s as HTML: %v", url, err)
	// }
	// ...print doc's title element…
	return nil
}

/* 对文件的操作，释放资源也可以使用defer */
func ReadFile(filename string) ([]byte, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return io.ReadAll(f)
}

/* 处理互斥锁也可以用defer */
var mu sync.Mutex
var m = make(map[string]int)

func lookup(key string) int {
	mu.Lock()
	defer mu.Unlock()

	return m[key]
}

/* 调试复杂程序时，defer机制也常被用于记录何时进入和退出函数 */
func bigSlowOperation() {
	defer trace("bigSlowOperation")() // don't forget the extra parentheses
	// ...lots of work…
	time.Sleep(10 * time.Second) // simulate slow operation by sleeping
}
func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}

/*
我们知道，对defer语句的执行是在return语句更新返回值变量之后，又因为在函数中定义的匿名函数可以访问该函数包括返回值在内的所有变量，
所以，对于匿名函数采用defer机制，可以使其观察函数的返回值
*/
func double(x int) (result int) {
	defer func() { fmt.Printf("double(%d) = %d\n", x, result) }()
	result = x + x
	return
}

/* 被延迟执行的函数甚至可以修改函数返回给调用者的返回值 */
func triple(x int) (result int) {
	defer func() {
		result += x
	}()
	return double(x)
}

/*
在循环体中的语句需要特别的注意，因为是在循环体完成后在执行。
下面的函数导致系统的文件描述符耗尽。
*/
func exhaust_fd(filenames []string) error {
	for _, filename := range filenames {
		f, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer f.Close()
	}
	return nil
}

/*
一种解决办法就是将循环体中的defer移至另外一个函数。每次循环时都调用这个函数。
*/
func doFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	return nil
}

func exhaust_fd_2(filenames []string) error {
	for _, filename := range filenames {
		return doFile(filename)
	}
	return nil
}

func main() {
	fmt.Println(double(4))
}
