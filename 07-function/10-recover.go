package main

import "fmt"

/*
Recover

通常不应该对panic做任何处理，但是有时候我们也许可以从异常中恢复，至少我们可以在程序崩溃前做一些操作。
*/

/*
如果在deferred函数中定义了内置函数recover，并且该deferred函数发生了panic异常，recover会使程序从panic中恢复，并返回panic value。
导致panic异常的函数不会继续运行，但能正常返回。在未发生panic时调用recover，recover会返回nil
*/
func f_recover() (err error) {
	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("internal error: %v", p)
		}
	}()

	panic("test panic")
}

func test_recover() {
	err := f_recover()
	fmt.Println(err)
}

func panic_wrapper() {
	panic("panic wrapper")
}

/* 外部的recover可以捕获子函数中的panic */
func recover_wrappered_panic() (err error) {
	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("internal error: %v", p)
		}
	}()
	panic_wrapper()
	return nil
}

func test_external_recover() {
	err := recover_wrappered_panic()
	fmt.Println(err)
}

/*
不加区分的恢复所有的panic异常，是不可取的做法：因为panic之后，无法保证包级变量的状态仍然和我们预期一致。
对panic的处理都集中在一个包下，有助于简化对复杂和不可预料问题的处理，但作为被广泛遵守的规范，我们不应该取试图恢复其它包引起的panic。
共有API应该将函数的运行失败作为error返回，而不是panic。同样的，你也不应该恢复一个由他人开发的函数引起的panic。
有时我们很难完全遵循规范，安全的做法是有选择的recover。只恢复应该被恢复的panic异常，同时这个比例应该尽可能的低。
为了表标识某panic是否应该被恢复，我们可以将panic value设置成特殊类型。在recover时对panic value进行检查，如果是特殊类型，就将这个
类型当作error处理，如果不是，则按照正常的panic进行处理。
*/
func recover_selectively() (err error) {
	defer func() {
		switch p := recover(); p {
		case nil: // no panic
		case "ABC":
			err = fmt.Errorf("panic: %v", "ABC")
		default:
			panic(p)
		}
	}()

	panic("ABC")
}

func main() {
	test_recover()
	test_external_recover()
}
