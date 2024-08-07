package vessel

import (
	"fmt"
)

// MapLearn map 是key-value 的无序集合
func MapLearn() {
	//stateMap()
	//forRangeTraverseMap()
	deleteItemFromMap()
}

// map 的声明和初始化
// 获取 map 的容量
func stateMap() {
	// 声明 map
	//var mapName map[keyType]valueType
	// 不需要提供 map 的长度，map 在使用过程中会自动扩充
	// 未初始化的 map 的值是 nil
	// 使用 len() 可以获取 map 中 pair 的数目
	var mapList map[string]int
	fmt.Println("mapList== nil: ", mapList == nil)
	// 初始化 map
	mapList = map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Println(mapList)
	// 使用 make 创建 map
	mapAssign := make(map[string]int)
	fmt.Println("mapAssign== nil: ", mapAssign == nil)
	mapAssign["one"] = 1
	mapAssign["two"] = 2
	fmt.Println(mapAssign)
	// 这种方式等价于 make(map[string]int)
	mapCreated := map[string]int{}
	fmt.Println("mapCreated== nil: ", mapCreated == nil)
	mapCreated["one"] = 1
	fmt.Println(mapCreated)
	// map 不存在固定长度或者最大限制, 可以随时增加新的键值对 动态伸缩
	// 但是也可以选择标明 map 的初始容量 capacity
	// 出于性能的考虑，对于大的 map 或者会快速扩张的 map，即使只是大概知道容量，也最好先标明。
	m2 := make(map[string]int, 100)
	fmt.Println(m2)
	// map 的值也可以是 切片
	m3 := make(map[string][]int)
	m3["one"] = []int{1, 2, 3}
	fmt.Println(m3)
}

// 遍历 map
func forRangeTraverseMap() {
	m1 := make(map[string]int, 100)
	m1["one"] = 1
	m1["two"] = 2
	m1["three"] = 3
	// 遍历 map 时，遍历输出元素的顺序与填充顺序无关，不能期望 map 在遍历时返回某种期望顺序的结果。的
	// 如果要按顺序遍历，需要手动对 key 进行排序
	for k, v := range m1 {
		fmt.Printf("key:%s,value:%d\n", k, v)
	}

	// 判断 某个 key 是否存在 map 中
	if _, ok := m1["one"]; ok {
		fmt.Println("one 存在")
	}
}

// 使用 delete() 函数从 map 中删除键值对
func deleteItemFromMap() {
	// 删除 map 中的一个元素
	m1 := make(map[string]int, 100)
	m1["one"] = 1
	m1["two"] = 2
	m1["three"] = 3
	delete(m1, "one")
	fmt.Println(m1)
	// 如果要删除的键不存在，不会报错
	delete(m1, "four")
	fmt.Println(m1)
	// 清空 map 中的所有元素
	//清空 map 的唯一办法就是重新 make 一个新的 map，
	//不用担心垃圾回收的效率，Go语言中的并行垃圾回收效率比写一个清空函数要高效的多。
}
