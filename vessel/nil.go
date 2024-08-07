package vessel

import "fmt"

// NilLearn nil 是一个预先定义好的标识符，它是许多类型的零值表示
func NilLearn() {

	// nil 标识符是不能比较的
	//invalid operation: nil == nil
	//fmt.Println(nil == nil)

	// nil 没有默认类型
	fmt.Println("%T", nil)

	// 不同类型的 nil 也不能比较
	var m map[string]int
	var ptr *int
	fmt.Println("m==nil", m == nil)
	fmt.Println(ptr == nil)
	//invalid operation: m == ptr
	//fmt.Println(m == ptr)

	// 两个相同类型的 nil 值也可能无法比较
	//var s1 s2 []int
	//fmt.Printf(s1 == s2)

	//nil 是 map、slice、pointer、channel、func、interface 的零值

	//不同类型的 nil 值占用的内存大小可能是不一样的
}
