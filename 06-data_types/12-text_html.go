package main

import (
	"html/template"
	"log"
	"os"
	texttemplate "text/template"
	"time"
)

/*
文本和HTML模板

仅仅使用fmt.Printf是不够的，有时候需要复杂的打印格式，这时候一般需要将格式化代码分离出来以便更安全的修改。
text/template和html/template包提供了这些功能，它们提供了一个将变量值填充到一个文本和html格式的模板的机制。

一个模板是一个字符串和一个文件，里面包含了一个或多个由花括号包含的{{action}}对象。
大部分的字符串都是按照字面值打印，但是对于actions部分将触发其它行为。
*/

type IssuesSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*Issue
}
type Issue struct {
	Number    int
	Title     string
	User      string
	CreatedAt time.Time `json:"created_at"`
}

func text_tmplate() {
	items := []*Issue{
		{Number: 1, Title: "A", User: "X1", CreatedAt: time.Now().Local().Add(time.Duration(-12 * time.Hour))},
		{Number: 2, Title: "B", User: "X2", CreatedAt: time.Now().Local().Add(time.Duration(-24 * time.Hour))},
		{Number: 3, Title: "C", User: "X3", CreatedAt: time.Now().Local().Add(time.Duration(-36 * time.Hour))},
		{Number: 4, Title: "D", User: "X4", CreatedAt: time.Now().Local().Add(time.Duration(-48 * time.Hour))},
		{Number: 5, Title: "E", User: "X5", CreatedAt: time.Now().Local().Add(time.Duration(-60 * time.Hour))},
	}
	results := IssuesSearchResult{TotalCount: 5, Items: items}

	const templ = `{{.TotalCount}} issues:
	{{range .Items}}------------------------------------
	Number: {{.Number}}
	User: {{.User}}
	Title: {{.Title | printf "%.64s"}}
	Age: {{.CreatedAt | daysAgo}} days
	{{end}}
	`
	/* 如果模板解析失败将是一个致命的错误，template.Must辅助函数可以简化这个致命错误的处理 */
	reporter := texttemplate.Must(texttemplate.New("report").Funcs(texttemplate.FuncMap{"daysAgo": days_ago}).Parse(templ))
	err := reporter.Execute(os.Stdout, results)
	if err != nil {
		log.Fatal(err)
	}
}

func html_template() {
	items := []*Issue{
		{Number: 1, Title: "A", User: "X1", CreatedAt: time.Now().Local().Add(time.Duration(-12 * time.Hour))},
		{Number: 2, Title: "B", User: "X2", CreatedAt: time.Now().Local().Add(time.Duration(-24 * time.Hour))},
		{Number: 3, Title: "C", User: "X3", CreatedAt: time.Now().Local().Add(time.Duration(-36 * time.Hour))},
		{Number: 4, Title: "D", User: "X4", CreatedAt: time.Now().Local().Add(time.Duration(-48 * time.Hour))},
		{Number: 5, Title: "E", User: "X5", CreatedAt: time.Now().Local().Add(time.Duration(-60 * time.Hour))},
	}
	results := IssuesSearchResult{TotalCount: 5, Items: items}

	var issueList = template.Must(template.New("issuelist").Parse(`
		<h1>{{.TotalCount}} issues</h1>
		<table>
		<tr style='text‐align: left'>
		<th>#</th>
		<th>User</th>
		<th>Title</th>
		</tr>
		{{range .Items}}
		<tr>
		<td><a href='#'>{{.Title}}</a></td>
		</tr>
		{{end}}
		</table>
		`))

	err := issueList.Execute(os.Stdout, results)
	if err != nil {
		log.Fatal(err)
	}
}

func days_ago(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

func main() {
	text_tmplate()
	html_template()
}
