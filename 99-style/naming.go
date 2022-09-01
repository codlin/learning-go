package main

/*
Go语言推荐驼峰式名命，当名字有几个单词组成时，优先使用首字母大小写分割，而是不用下划线分割。
对于缩略词则避免混合的写法。
*/
func main() {

	// good style
	thisVar := ""
	ThisVar := ""

	// bad style
	this_var := ""
	this_Var := ""

	print(thisVar, ThisVar, this_var, this_Var)

	/* 对于缩略词则避免混合的写法 */
	// good style
	escapeHTML := "<span>go</span>"
	htmlEscape := "<span>go</span>"
	HTMLEscape := "<span>go</span>"

	// bad style
	escapeHtml := "<span>go</span>"
	print(escapeHTML, htmlEscape, HTMLEscape, escapeHtml)
}
