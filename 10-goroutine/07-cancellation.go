package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

/*
并发的取消

有时候我们需要通知goroutine停止它正在干的事，但Go语言并没有提供在一个goroutine中终止另一个goroutine的方法
由于这样会导致goroutine之间的共享变量落在未定义的状态。
*/

// 这个程序可能还存在问题

var sema = make(chan struct{}, 20)
var done = make(chan struct{})

// 增加信号量控制并发
func walkDir3(dir string, n *sync.WaitGroup, fileSizes chan<- int64) {
	defer n.Done()
	if canceled() {
		return
	}

	for _, entry := range dirents2(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			n.Add(1)
			go walkDir3(subdir, n, fileSizes)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

func canceled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

// 增加信号量控制并发
func dirents2(dir string) []os.FileInfo {
	sema <- struct{}{}
	select {
	case sema <- struct{}{}:
	case <-done:
		return nil
	}
	defer func() { <-sema }()

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

// du5版本增加并发的退出
func du5() {
	var verbose = flag.Bool("v", false, "show verbose progress message")
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}

	var wg sync.WaitGroup
	fileSizes := make(chan int64)
	for _, root := range roots {
		wg.Add(1)
		go walkDir3(root, &wg, fileSizes)
	}
	go func() {
		wg.Wait()
		close(fileSizes)
	}()

	go func() {
		os.Stdin.Read(make([]byte, 1))
		fmt.Println("user abort")
		close(done)
	}()

	var tick <-chan time.Time
	if *verbose {
		// log.Print("create ticker")
		ticker := time.NewTicker(500 * time.Millisecond)
		defer ticker.Stop()
		// defer log.Print("stop ticker")
		tick = ticker.C
	}

	var nfiles, nbytes int64
loop:
	for {
		select {
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		case <-done:
			// drain channel
			for range fileSizes {
			}
			return
		}
	}
	printDiskUsage(nfiles, nbytes)
}

func main() {
	du5()
}
