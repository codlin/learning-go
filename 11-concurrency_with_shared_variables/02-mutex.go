package main

/*
互斥锁 Mutual Exclusion
*/

/*
1. 什么是可重入锁？
   可重入锁指的是在一个线程中可以多次获取同一把锁，比如：
   一个线程在执行一个带锁的方法，该方法中又调用了另一个需要相同锁的方法，则该线程可以直接执行调用的方法，而无需重新获得锁；
   需要注意的是锁的操作粒度是”线程”，而不是调用，同一个线程再次进入同步代码的时候，可以使用自己已经获取到的锁，这就是可重入锁。
   java里面内置锁(synchronize)和Lock(ReentrantLock)都是可重入。
   可重入锁主要用在线程需要多次进入临界区代码时，需要使用可重入锁。
   为每个锁关联一个获取计数器和一个所有者线程，当计数值为0的时候，这个锁就没有被任何线程持有。
   当线程请求一个未被持有的锁时，将记下锁的持有者，并且将获取计数值置为1，如果同一个线程再次获取这个锁，计数值将递增，
   退出一次同步代码块，计算值递减，当计数值为0时，这个锁就被释放。
   如果是不同的线程获取锁，则这个线程就会休眠，直至获取成功、或立即返回false，或者是超时。

2. Go为什么不支持可重入锁？
   下面是《The Go Programming Language》的中译本给出的解释：
   --------------------------------------------------------------------------------------------------------
  『Go 的互斥锁不能重入是有充分理由的。
   互斥锁的目的是确保共享变量的某些不变量在程序期间的关键点保持不变执行。
   其中一个不变量是“没有goroutine正在访问共享变量”，但可能存在特定于互斥锁保护的数据结构的其他不变量。
   当一个goroutine 获得一个互斥锁，它可以假设不变量成立。虽然它持有锁，它可能会更新共享变量，以便暂时违反不变量。
   但是，当它释放锁时，它必须保证顺序已经恢复，并且不变量再次成立。
   尽管可重入互斥锁将确保没有其它goroutines正在访问共享变量，它不能保护额外的不变量那些变量。』
   --------------------------------------------------------------------------------------------------------
   下面是《The Go Programming Language》原文：
   --------------------------------------------------------------------------------------------------------
   There is a good reason Go’s mutexes are not re-entrant.

   The purpose of a mutex is to ensure that certain invariants of the shared variables are
   maintained at critical points during program execution.

   One of the invariants is "no goroutine is accessing the shared variables", but there may be
   additional invariants specific to the data structures that the mutex guards.

   When a goroutine acquires a mutex lock, it may assume that the invariants hold. While it
   holds the lock, it may update the shared variables so that the invariants are temporarily violated.

   However, when it releases the lock, it must guarantee that order has been restored
   and the invariants hold once again.

   Although a re-entrant mutex would ensure that no other goroutines are accessing the shared variables,
   it cannot protect the additional invariants of those variables.
   --------------------------------------------------------------------------------------------------------
   不管是中文还是英文，这段话都很难理解：
   a. 如何理解不变量 invariants ？
   b. additional invariants 具体指的是什么样的场景？
   c. shared variables 和 invariants 是什么关系？
   d. "...guarantee that order has been restored..." 是什么意思？
*/

/*
这是在网上找到的关于为什么互斥锁不支持可重入的资料：
--------------------------------------------------------------------------------------------------------
考虑如下代码：

func F() {
   mu.Lock()
   ... do some stuff ...
   G()
   ... do some more stuff ...
   mu.Unlock()
}

func G() {
   mu.Lock()
   ... do some stuff ...
   mu.Unlock()
}

如果是可重入锁则不会死锁，进入F时Lock一次锁的可重入次数+1，进入G时又Lock一次又+1，对应Unlock时再-1，正常执行结束。
而且可重入锁可在用户进程空间实现，避免了额外开销，又能解决死锁，性能也比重量级锁好，为什么还说其设计“不太好”呢？

可以看到，作为为并发而生的Go语言一直没有引入可重入锁，就是因为Russ Cox大佬坚持可重入锁的设计理念不好，他认为，既然你加了锁，
那你就要保证锁住的这部分资源/区域具有“状态不变性”，具有状态不变性的集合称为状态不变量（Russ Cox称为“the invariants”），
而可重入锁破坏了这种不变性状态。

状态不变量（“the invariants”）简单理解就是一个同步的范围，或者一个同步的元素集合，显然在上述代码中，如果是可重入锁，
直接执行F是没问题，但如果F、G都是对外暴露的方法呢？那么外界在单独调用G的时候，很有可能有其他线程（用户）在调用F，
这时，两种不同性质的线程就会用到同一个mutex，这就会出问题。这里的不同性质的线程是指不同业务的线程，如果多个线程只执行F函数，
那么这些线程是相同性质的线程。
一个mutex理应当只保护一个同步范围，对同一性质的多个线程进行同步，正如Russ Cox说的：
“Recursive mutexes do not protect invariants. Mutexes have only one job, and recursive mutexes don't do it.”
（普通mutex保证单一职责，但可重入mutex做不到），所以出问题的原因就在于这里异化了mutex的职责————本来只用保护一块同步范围，
现在要保护两块，即mutex出现了两种场景下的语义，混用了语义，当然会容易出错。

有人会问，你这不太对啊，怎么感觉怪怪的，你上面代码mutex换成普通的锁，一样会出问题呀，这和可重入锁有啥关系？
没错，如果F、G都是对外暴露的方法，在多线程环境下这么写确实也会出问题，比如F如果执行过长，会导致本来很快执行完的G超时，
这是最简单的情况，属于程序员的低级错误。但问题的本质不在这里，问题的本质在于是否混淆了语义。
如果mutex上的是普通锁，那么从宏观角度看，要么是G执行，要么是F执行，两者是互斥的语义，不存在G和F的交集。
如果mutex上的是可重入锁，那么从宏观角度看，同样是G执行时，F不能执行，F执行时，G不能执行。

但若将视角切细点，从F的角度看，那么会发现，F执行时，G也可以执行，那么我到底听谁的？这就好像你新来的接手一个复杂项目，
文档告诉你G和F的调用方式是互斥的，但实际上他又可以不是完全互斥的—————其实这就是混淆了mutex的语义。
而正是这种看似不起眼的语义混淆，在复杂的代码环境下，就是滋生bug的温床。这也是Russ Cox亲自说的：
“Recursive mutexes are just a mistake, nothing more than a comfortable home for bugs.”
--------------------------------------------------------------------------------------------------------
上面这段代码说的还算不错，但后面的感觉解释的有些牵强，还是有种隔着迷雾的感觉。
下面是Russ Cox在Google Group里的相关内容的回复：
--------------------------------------------------------------------------------------------------------
Recursive (aka reentrant) mutexes are a bad idea.
The fundamental reason to use a mutex is that mutexes protect invariants,
perhaps internal invariants like `p.Prev.Next == p` for all elements of the ring,
or perhaps external invariants like `my local variable x is equal to p.Prev'.

Locking a mutex asserts `I need the invariants to hold`
and perhaps `I will temporarily break those invariants.`
Releasing the mutex asserts `I no longer depend on those invariants`
and `If I broke them, I have restored them.`

Understanding that mutexes protect invariants is essential to
identifying where mutexes are needed and where they are not.
For example, does a shared counter updated with atomic
increment and decrement instructions need a mutex?
It depends on the invariants. If the only invariant is that
the counter has value i - d after i increments and d decrements,
then the atmocity of the instructions ensures the
invariants; no mutex is needed. But if the counter must be
in sync with some other data structure (perhaps it counts
the number of elements on a list), then the atomicity of
the individual operations is not enough. Something else,
often a mutex, must protect the higher-level invariant.
This is the reason that operations on maps in Go are not
guaranteed to be atomic: it would add expense without
benefit in typical cases.

Let's take a look at recursive mutexes.
Suppose we have code like this:

func F() {
mu.Lock()
... do some stuff ...
G()
... do some more stuff ...
mu.Unlock()
}

func G() {
mu.Lock()
... do some stuff ...
mu.Unlock()
}

Normally, when a call to mu.Lock returns, the calling code
can now assume that the protected invariants hold, until
it calls mu.Unlock.

A recursive mutex implementation would make G's mu.Lock
and mu.Unlock calls be no-ops when called from within F
or any other context where the current thread already holds mu.
If mu used such an implementation, then when mu.Lock
returns inside G, the invariants may or may not hold. It depends
on what F has done before calling G. Maybe F didn't even realize
that G needed those invariants and has broken them (entirely
possible, especially in complex code).

Recursive mutexes do not protect invariants.
Mutexes have only one job, and recursive mutexes don't do it.

There are simpler problems with them, like if you wrote

func F() {
mu.Lock()
... do some stuff
}

you'd never find the bug in single-threaded testing.
But that's just a special case of the bigger problem,
which is that they provide no guarantees at all about
the invariants that the mutex is meant to protect.

If you need to implement functionality that can be called
with or without holding a mutex, the clearest thing to do
is to write two versions. For example, instead of the above G,
you could write:

// To be called with mu already held.
// Caller must be careful to ensure that ...
func g() {
... do some stuff ...
}

func G() {
mu.Lock()
g()
mu.Unlock()
}

or if they're both unexported, g and gLocked.

I am sure that we'll need TryLock eventually; feel free to
send us a CL for that. Lock with timeout seems less essential
but if there were a clean implementation (I don't know of one)
then maybe it would be okay. Please don't send a CL that
implements recursive mutexes.

Recursive mutexes are just a mistake, nothing more than
a comfortable home for bugs.

Russ
--------------------------------------------------------------------------------------------------------
下面是另一位网友的解释：
--------------------------------------------------------------------------------------------------------
Say the mutex protects the invariant that a == b, as in these functions:

func up() {
   mutex.Lock()
   a++
   // See below
   b++
   mutex.Unlock()
}


func down() {
   mutex.Lock()
   if a != b { panic("HELP!") }
   a--
   b--
   mutex.Unlock()
}

What happens if down() is called on the line marked "See below" in f?
(this could happen through subtle paths in realistic code, as you know.)
The invariant is not true when down() is called *even though it holds the mutex*, and down() will panic.

That is what it means to say that reentrant mutexes don't preserve invariants.
It's because they don't preserve invariants when called recursively, which defeats the purpose of having a mutex.

-rob
--------------------------------------------------------------------------------------------------------
rob的意思应该是说，如果up和down都是保护不变量(the invariant) a 和 b的一致性，如果是不可重入锁，自然没问题。
如果是可重入锁的话，则会无法保证同步性，down()将会panic。譬如：

func up() {
   mutex.Lock()
   a++

   // See below
   down()

   b++
   mutex.Unlock()
}
*/
