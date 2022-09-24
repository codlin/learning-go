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

	"golang.org/x/net/html"
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
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
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
释放资源的defer应该直接跟在请求资源的语句后。
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
	doc, err := html.Parse(resp.Body)
	if err != nil {
		return fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
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
