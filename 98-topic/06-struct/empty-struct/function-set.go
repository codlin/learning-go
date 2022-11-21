package main

/**
# 实现集合数据结构
集合数据结构我们可以使用map来实现：只关心key，不必关心value，我们就可以值设置为空结构体类型变量（或者底层类型是空结构体的变量）。
**/

type Set struct {
	items map[interface{}]emptyItem
}

type emptyItem struct{}

var itemExists = emptyItem{}

func NewSet() *Set {
	set := &Set{items: make(map[interface{}]emptyItem)}
	return set
}

// 添加元素到集合
func (set *Set) Add(item interface{}) {
	set.items[item] = itemExists
}

// 从集合中删除元素
func (set *Set) Remove(item interface{}) {
	delete(set.items, item)
}

// 判断元素是否存在集合中
func (set *Set) Contains(item interface{}) bool {
	_, contains := set.items[item]
	return contains
}

// 返回集合大小
func (set *Set) Size() int {
	return len(set.items)
}

func main() {
	set := NewSet()
	set.Add("hello")
	set.Add("world")
	println(set.Contains("hello"))
	println(set.Contains("Hello"))
	println(set.Size())
}
