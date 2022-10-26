package main

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

/*
Looping in Parallel

在loop中创建并发goroutine时需要注意的地方
*/

// ImageFile reads an image from infile and writes
// a thumbnail‐size version of it in the same directory.
// It returns the generated file name, e.g., "foo.thumb.jpg".
func ImageFile(infile string) (string, error) {
	fmt.Printf("process file: %s\n", infile)
	time.Sleep(time.Duration(rand.Intn(3)) * time.Second)
	fmt.Printf("process file: %s, done\n", infile)

	return infile, nil
}

// 串行版本
func makeThumbnails(filenames []string) {
	for _, f := range filenames {
		if _, err := ImageFile(f); err != nil {
			log.Println(err)
		}
	}
}

// 并发版本，但是有错误，因为在goroutines还没有完成时makeThumbnails2已经返回了
func makeThumbnails2(filenames []string) {
	for _, f := range filenames {
		go ImageFile(f) // NOTE: ignoring errors
	}
}

/*
没有什么直接的办法能够等待goroutine，但是我们可以改变goroutine里的代码让其能够对完成的情况报告给外部goroutine知晓，
使用的方式是向一个共享的channel中发送事件。
下面的例子，我们已经知道了goroutine的数量为len(filenames)，所以makeThumbnails3只要在返回对这些事件计数。
*/
func makeThumbnails3(filenames []string) {
	ch := make(chan struct{})
	for _, f := range filenames {
		go func(f string) {
			ImageFile(f) // NOTE: ignoring errors
			ch <- struct{}{}
		}(f) // NOTE: 显式传入f（匿名函数中的循环变量快照问题）
	}
	// Wait for goroutines to complete.
	for range filenames {
		<-ch
	}
}

// 如果想在goroutine里返回错误，可以创建一个元素类型是error的channel
func makeThumbnails4(filenames []string) error {
	errors := make(chan error)
	for _, f := range filenames {
		go func(f string) {
			_, err := ImageFile(f)
			errors <- err
		}(f)
	}

	for range filenames {
		if err := <-errors; err != nil {
			return err // Error! 此处造成了goroutine泄露
			/* 因错误返回后，channel里的数据无法被排空(drain)，这样其它goroutines在发送值时，都会永远的阻塞下去，并且永远都不会退出 */
		}
	}

	return nil
}

// 针对上面的错误，最简单的办法时创建一个带buffer的channel，这样即使因错误退出，也不会造成goroutines泄露了
func makeThumbnails5(filenames []string) (thumbfiles []string, err error) {
	type item struct {
		thumbfile string
		err       error
	}
	ch := make(chan item, len(filenames))
	for _, f := range filenames {
		go func(f string) {
			var it item
			it.thumbfile, it.err = ImageFile(f)
			ch <- it
			fmt.Printf("wrote [%s] to channel\n", it.thumbfile)
		}(f)
	}
	/* 这个版本用了buffered channel，因此各个goroutines可以把处理的结果并发的写入到channel里，
	所以可能第一次迭代后实际结果已经完成了，因此再使用filenames循环迭代可能不会出现问题，
	但是用这个循环次数来判断是否goroutines是否结束已经不准确了。
	*/
	for _, f := range filenames {
		fmt.Println(f)

		it := <-ch
		if it.err != nil {
			return nil, it.err
		}

		fmt.Printf("got [%s] from channel\n", it.thumbfile)
		thumbfiles = append(thumbfiles, it.thumbfile)
	}
	return thumbfiles, nil
}

/*
为了知道最后一个goroutine什么时间结束，我们需要一个递增的计数器，在每一个goroutine启动时加1，在goroutine退出时减1。
这种计数器需要在多个goroutine操作时做到安全并且提供在其减为0之前一直等待的一直方法。
Go语言提供了这种计数类型称之为sync.WaitGroup
*/
func makeThumbnails6(filenames []string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup
	for _, f := range filenames {
		wg.Add(1)
		// worker
		go func(f string) {
			defer wg.Done()

			_, err := ImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}
			// mock
			sizes <- int64(rand.Intn(10000))
		}(f)
	}

	// closer
	go func() {
		wg.Wait()
		close(sizes)
	}()

	var total int64
	for size := range sizes {
		fmt.Printf("got size: %d\n", size)
		total += size
	}
	return total
}
func test_makeThumbnails6() {
	filenames := []string{"/tmp/a", "/tmp/b", "/tmp/c", "/tmp/d"}
	sizes := makeThumbnails6(filenames)
	fmt.Printf("total size: %d\n", sizes)
}

func main() {
	// makeThumbnails5([]string{"a", "b", "c", "d"})
	test_makeThumbnails6()
}
