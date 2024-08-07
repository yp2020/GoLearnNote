package vessel

import "fmt"

func ArrayLearn() {
	// 声明数组
	// var 数组变量名 [元素数量]Type
	var a1 [3]int
	// 默认情况下，数组的每个元素都会被初始化成对应的 0 值
	fmt.Println(a1[0])
	fmt.Println(a1[len(a1)-1])
	// 声明数组并初始化
	var a = [3]int{1, 2, 3}
	fmt.Println(a[1])
	// 长度可以省略
	var b = [...]int{1, 2, 3}
	fmt.Println("len(b) = ", len(b))
	// 数组长度也是类型的一部分，
	//且长度必须是常量表达式，因为数组的长度需要在编译阶段确定
	fmt.Printf("%T\n", b)
	// 比较两个数组是否相等
	//数组的长度,数组中元素的类型都相同才可以认为数组的类型相同
	// 数组类型相同才可以比较是否相等
	var c = [3]int{1, 2, 3}
	d := [...]int{1, 2, 3}
	e := [3]int{2, 3, 4}
	fmt.Println(c == d, c == e, d == e)
	// 编译错误，无法比较
	//f := [...]int{1, 2}
	//fmt.Println(c==f)

	// 遍历数组 -- 访问数组中的每一个数组元素
	for i := 0; i < len(a); i++ {
		fmt.Printf("%d ", a[i])
	}
	fmt.Println()
	// for-range 遍历
	for k, v := range a {
		fmt.Printf("%d:%d\n", k, v)
	}
}

func MultidimensionalArray() {
	// 声明二维数组
	var a1 [3][2]int
	fmt.Println("a1: ", a1)
	// 使用数组字面量来声明并且初始化一个二维整形数组
	a2 := [4][2]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}
	fmt.Println("a2: ", a2)
	// 声明并初始化部分元素,比如某个索引下
	a3 := [4][2]int{1: {20, 21}, 2: {40, 41}}
	fmt.Println("a3: ", a3)
	a4 := [4][2]int{1: {1: 21}, 2: {1: 22}}
	fmt.Println("a4: ", a4)

	// 访问多维数组中的元素
	a5 := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("a5[1][1]:", a5[1][1])

	// 同样类型的多维数组赋值
	a6 := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("a6: ", a6)
	var a7 [2][3]int
	fmt.Println("a7 before: ", a7)
	a7 = a6
	fmt.Println("a7 after: ", a7)
	// 单独赋某一个维度的值
	var a8 [3]int
	fmt.Println("a8 before: ", a8)
	a8 = a6[1]
	fmt.Println("a8 after: ", a8)
	var value int = a6[1][1]
	fmt.Println("value: ", value)
}
