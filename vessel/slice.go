package vessel

import "fmt"

// SliceLearn 切片的使用,底层结构后续补上
func SliceLearn() {
	//SliceFromArray()
	//GenerateSlice()
	//generateSliceUseMake()
	//appendItem()
	//copySlice()
	//referencePass()
	//deleteItem()
	//forRangeTraverse()
	multiSlice()
}

func sliceFromArray() {
	// 从数组生成切片
	var a = [5]int{1, 2, 3, 4, 5}
	fmt.Println("arr: ", a)
	// 取出位置1到2的元素 不包括结束位置的索引
	fmt.Println("s1: ", a[1:3])
	// 缺省开始位置,表示从连续区域开头到结束位置
	fmt.Println("s2: ", a[:2])
	// 缺省结束位置,表示从开始位置到整个连续区域末尾
	fmt.Println("s3: ", a[1:])
	// 缺省开始和结束位置,表示整个连续区域,
	fmt.Println("s4: ", a[:])
	// 两者同时为 0,等效于 空切片，一般用于切片复位,清空所有的元素
	fmt.Println("s5: ", a[0:0])
}

func generateSlice() {
	// 声明切片
	// var name []Type
	// 声明字符串切片
	var strList []string
	// 声明整型切片
	var numList []int
	fmt.Println("strList: ", strList)
	fmt.Println("numList: ", numList)
	// 声明但是为使用的切片的默认值为 nil
	fmt.Println("strList == nil: ", strList == nil)
	fmt.Println("numList == nil: ", numList == nil)
	// 声明一个空的切片,切片为空，但是此时的 numListEmpty 已经被分配了内存，只是还没有元素。
	var numListEmpty = []int{}
	fmt.Println("numListEmpty: ", numListEmpty)
	fmt.Println("numListEmpty == nil: ", numListEmpty == nil)
}

func generateSliceUseMake() {
	// make([]Type, size, cap)
	// Type 指的是切片的类型
	// size 指分配的元素个数
	// cap 为预分配的元素数量,这个值设定后不影响 size，
	//只是能提前分配空间，降低多次分配空间造成的性能问题。
	// 使用 make() 生成的切片一定发生了内存分配操作
	// 但是给定开始与结束位置或者复位并不会发生内存分配操作,只是将新的切片结构指向已经分配好的内存区域，设定开始与结束位置
	a := make([]int, 2)
	b := make([]int, 2, 4)
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	fmt.Printf("len(a):%d,cap(a):%d\n", len(a), cap(a))
	fmt.Printf("len(b):%d,cap(b):%d\n", len(b), cap(b))
}

func appendItem() {
	var s1 []int
	fmt.Println("s1: ", s1)
	fmt.Println(len(s1))
	fmt.Println(cap(s1))
	// 追加一个元素
	s1 = append(s1, 1)
	fmt.Println("s1: ", s1)
	//追加一个切片 使用 ... 表示解包成单个元素
	s1 = append(s1, []int{1, 2, 3}...)
	fmt.Println("s1: ", s1)
	//在使用 append() 函数为切片动态添加元素时，
	//如果空间不足以容纳足够多的元素，切片就会进行“扩容”，
	//此时新切片的长度会发生改变。
	fmt.Println(len(s1))
	fmt.Println(cap(s1))
	// 切片在扩容时,容量会翻倍增长
	var numbers []int
	for i := 0; i < 10; i++ {
		numbers = append(numbers, i)
		fmt.Printf("len: %d  cap: %d numbers: ", len(numbers), cap(numbers))
		fmt.Println(numbers)
	}
}

// copy(dest,src,[]T) int
// 解释:  将 src 复制到 dest 中,返回实际发生复制的元素个数
// 要求:  dest 与 src 的类型必须一致 && 目标切片必须分配过空间且足够承载复制的元素个数
// 如果切片 dest src 不一样大，就会按照其中较小的那个数组切片的元素个数进行复制
func copySlice() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := []int{11, 12, 13}
	fmt.Println("before copy: ")
	fmt.Println("s1: ", s1)
	fmt.Println("s2: ", s2)
	cnt1 := copy(s1, s2)
	fmt.Println("after copy s2 to s1")
	fmt.Println("s1: ", s1)
	fmt.Println("s2: ", s2)
	fmt.Println("cnt: ", cnt1)

	s1 = []int{1, 2, 3, 4, 5}
	s2 = []int{11, 12, 13}
	cnt2 := copy(s2, s1)
	fmt.Println("after copy s1 to s2")
	fmt.Println("s1: ", s1)
	fmt.Println("s2: ", s2)
	fmt.Println("cnt: ", cnt2)
}

// slice 是引用传递
func referencePass() {
	const elementCount = 1000
	// 预分配足够多的元素切片
	srcData := make([]int, elementCount)
	// 将切片赋值
	for i := 0; i < elementCount; i++ {
		srcData[i] = i
	}
	// 引用切片数据
	refData := srcData

	// 预分配足够多的元素切片
	copyData := make([]int, elementCount)
	// 将数据复制到新的切片空间中
	copy(copyData, srcData)

	// 修改原始数据的第一个元素
	srcData[0] = 999
	// 打印引用切片的第一个元素
	fmt.Println(refData[0])
	// 打印复制切片的第一个和最后一个元素
	fmt.Println(copyData[0], copyData[elementCount-1])

	// 复制原始数据从4到6(不包含)
	copy(copyData, srcData[4:6])

	for i := 0; i < 5; i++ {
		fmt.Printf("%d ", copyData[i])
	}
}

// 从切片中删除元素，
// Go 语言本身就没有为删除元素提供专用的语法或者接口
// 做法是以被删除元素为分界点，将前后两个部分的内存重新连接起来
// 连续的容器在删除元素时都涉及到内存的重新分配和移动，随着元素的增加,
// 这个 过程变的极为耗时，因此当业务需要大量,频繁的从一个切片中删除元素时,
// 如果对性能有要求,就考虑更换其他的容器
func deleteItem() {
	// 删除开头位置的元素
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(a)
	index := 4
	// 删除a[4]
	a = append(a[:index], a[index+1])

}

// for-range 遍历切片
func forRangeTraverse() {
	slice1 := []int{1, 2, 3, 4, 5}

	for idx, val := range slice1 {
		fmt.Printf("idx:%d,val:%d\n", idx, val)
	}
	// range 创建的是每个元素的副本, 而不是元素的引用
	//如果想获得每个元素的地址，需要使用 &slice[idx]
	for idx, val := range slice1 {
		fmt.Printf("idx:%d,val:%d,value-addr:%X,elem-addr:%X\n", idx, val, &val, &slice1[idx])
	}
}

// 多维切片
func multiSlice() {
	slice := [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(slice)
	slice[0] = append(slice[0], 7)
	fmt.Println(slice)
}
