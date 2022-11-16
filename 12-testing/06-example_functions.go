package main

/*
示例函数

被go test特殊对待的第三种函数是示例函数，它的名字以Example开头。它既没有参数也没有结果。例如：
func ExampleIsPalindrome() {
    fmt.Println(IsPalindrome("A man, a plan, a canal: Panama"))
    fmt.Println(IsPalindrome("palindrome"))
    // Output:
    // true
    // false
}
示例函数有三个目的：
1. 首要的目的是作为文档；比起乏味的描述，举一个好的例子是描述库函数功能最简洁和直观的方式。
   和带注释的例子不同，示例函数是真实的Go代码，必须通过编译时检查，所以随着代码的进化它们也不会过时。
2. 第二个目的是它们是可以通过go test运行的可执行测试。
3. 第三个目的是提供手动实验代码。在golang的文档服务器可以使用Go Playground来让用户在web浏览器
   上面编辑和运行函数示例。这通常是了解特定函数功能或者了解语言特性最快捷的方法。
*/
