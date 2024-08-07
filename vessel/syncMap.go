package vessel

import (
	"fmt"
	"sync"
)

// SyncMapLearn  map 在并发情况下，只读是线程安全的，同时读写是线程不安全的
// sync.Map 为了保证并发安全有一些性能损失，因此在非并发情况下，使用 map 相比使用 sync.Map 会有更好的性能。
func SyncMapLearn() {
	//unSafeMapTest()
	safeMapTest()
}

/*
sync.Map 有以下特性
1. 无须初始化,直接声明即可
2.sync.Map 不能使用 map 的方式进行取值和设置等操作，
  而是使用 sync.Map 的方法进行调用，Store 表示存储，Load 表示获取，Delete 表示删除。
3.使用 Range 配合一个回调函数进行遍历操作，通过回调函数返回内部遍历出来的值，
  Range 参数中回调函数的返回值
   在需要继续迭代遍历时，返回 true，
   终止迭代遍历时，返回 false。
*/

// 最后报错 fatal error: concurrent map read and map write
func unSafeMapTest() {
	m := make(map[int]int)
	// 不停的对 map 进行写入
	go func() {
		for {
			m[1] = 1
		}
	}()
	// 不停的对 map 进行读取
	go func() {
		for {
			_ = m[1]
		}
	}()
	// 无限程序，阻塞程序运行
	for {

	}

}

func safeMapTest() {
	var m1 sync.Map
	m1.Store(1, 1)
	m1.Store(2, 2)
	m1.Store(3, 3)
	// 遍历 sync.Map 中的键值对
	m1.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})

	// 从 sync.Map 中根据键取值
	value, ok := m1.Load(1)
	fmt.Println("m1.Load(1) = ", value, ok)

	// 从 sync.Map 中根据键删除值
	m1.Delete(1)

	// 再遍历 sync.Map
	m1.Range(func(key, value interface{}) bool {
		fmt.Println(key, value)
		return true
	})
	//sync.Map 还支持其他的方法
}
