package main

import (
	"fmt"
	"strings"
	"sync"
)

/*
竞争条件

竞争条件是指多个goroutine在交叉操作时，没有给出正确的结果，通常goroutines会并发访问相同的变量并且其中至少一个为写操作。
有三种方式可以避免数据竞争：
 1. 不要去写变量。在创建goroutine之前将数据初始化，然后所有goroutine只读这个变量。
 2. 避免从多个goroutine访问变量。
    由于其它goroutine不能直接访问变量，它们只能通过一个channel给指定的goroutine发送请求来查询或更新变量。
    即“不要使用共享数据来通信，使用通信来共享数据”。提供代理访问受限变量的这个goroutine称之为监听goroutine(monitor goroutine)。
 3. 允许很多goroutine去访问变量，但是同一时刻最多只有一个goroutine在访问。这种方式称之为“互斥”（后续章节）。
*/
var balance int

func Deposit(amount int) { balance += amount }
func Balance() int       { return balance }

// 该程序并发更新同一变量，存在数据竞争，得到的结果不确定。
func concurrent_deposit() {
	var wg sync.WaitGroup

	// Bob
	wg.Add(1)
	go func() {
		defer wg.Done()
		Deposit(200)
		fmt.Printf("Bob balance=%d\n", Balance())
	}()

	// Alice
	wg.Add(1)
	go func() {
		defer wg.Done()
		Deposit(100)
		fmt.Printf("Alice balance=%d\n", Balance())
	}()

	wg.Wait()
}

// 避免竞争的第2种方式，使用monitor goroutine来代理变量的访问，将变量的访问绑定在这个goroutine上。
var deposits = make(chan int)
var balances = make(chan int)

func Deposit2(amount int) { deposits <- amount }
func Balance2() int       { return <-balances }
func teller() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}
func confined_desposit() {
	var wg sync.WaitGroup

	// Bob
	wg.Add(1)
	go func() {
		defer wg.Done()
		Deposit2(200)
		fmt.Printf("Bob desposit: balance=%d\n", Balance2())
	}()

	// Alice
	wg.Add(1)
	go func() {
		defer wg.Done()
		Deposit2(100)
		fmt.Printf("Alice desposit: balance=%d\n", Balance2())
	}()

	wg.Wait()
}

func init() {
	go teller()
}

/*
即使一个变量无法在其整个生命周期内被限制到单个goroutine，加以限制仍然可以是解决并发访问的好方法。
比如一个场景的场景，可以通过借助channel把共享变量的地址从上一步传递到下一步，从而在流水线上的多个goroutine之间共享该变量。
在流水线的每一步，把变量地址传给下一步后就不访问该变量了，这样对所有这个变量的访问都是串行的。
换个说法，这个变量先受限于流水线的一步（step），再受限于下一步，以此类推。这种受限有时也称为串行受限。
*/
type Cake struct{ state string }

func baker(cooked chan<- *Cake) {
	for {
		cake := new(Cake)
		cake.state = "cooked"
		cooked <- cake // baker never touches this cake again
	}
}

func icer(iced chan<- *Cake, cooked <-chan *Cake) {
	for cake := range cooked {
		cake.state = "iced"
		iced <- cake // icer never touches this cake again
	}
}

func main() {
	fmt.Printf("%s\n", strings.Repeat("=", 64))
	concurrent_deposit()

	fmt.Printf("%s\n", strings.Repeat("-", 64))
	confined_desposit()

	fmt.Printf("%s\n", strings.Repeat("=", 64))
}
