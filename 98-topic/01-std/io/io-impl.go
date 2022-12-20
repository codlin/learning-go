package main

/*
标准库io接口的实施练习
*/

/*
对于io库中Reader接口描述的翻译：
type Reader interface {
	Read(p []byte) (n int, err error)
}

Reader 是包装基本 Read 方法的接口。

Read 将最多 len(p) 个字节读入 p。它返回字节数read (0 <= n <= len(p)) 和遇到的任何错误。
即使阅读返回 n < len(p)，它可以在调用期间使用所有 p 作为暂存空间。
如果有一些数据可用但没有 len(p) 字节，则按常规读取返回可用的而不是等待更多。

当 Read 之后遇到错误或文件结束条件时成功读取 n > 0 个字节，它返回的数量字节读取。
它可能会从同一调用返回（非零）错误或者从后续调用中返回错误（并且 n == 0）。
这种一般情况的一个例子是 Reader 返回输入流末尾的非零字节数可能返回 err == EOF
或 err == nil。下次阅读应该返回 0，结束。

Callers should always process the n > 0 bytes returned before
considering the error err. Doing so correctly handles I/O errors
that happen after reading some bytes and also both of the
allowed EOF behaviors.
调用方在考虑错误err前应始终处理返回的 n > 0 个字节。
这样做可以正确处理在读取了一些字节之后产生的I/O 错误，允许的 EOF 行为。

不鼓励 Read 的实现返回一个零字节计数和nil错误，除非 len(p) == 0。
调用者应该将 0 和 nil 的返回值视为表示什么也没发生；特别是它不表示 EOF。
*/
