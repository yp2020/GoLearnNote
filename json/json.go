package json1

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name string
	Age  int
	Sex  string
}

type Class struct {
	Id       string // 名字小写的时候，序列化会将其去掉。
	Students []Student
}

func LearnJson() {
	Student1 := Student{
		Name: "张三",
		Age:  18,
		Sex:  "男",
	}
	Student2 := Student{
		Name: "李四",
		Age:  10,
		Sex:  "女",
	}
	Student3 := Student{
		Name: "王五",
		Age:  12,
		Sex:  "男",
	}
	c1 := Class{
		Id:       "1(2)班级",
		Students: []Student{Student1, Student2, Student3},
	}
	marshal, err := json.Marshal(c1)
	if err != nil {
		fmt.Println("marshal error", err)
	}
	fmt.Println("marshal result", string(marshal))
	fmt.Println("========================================")
	var c2 Class
	err = json.Unmarshal(marshal, &c2)
	if err != nil {
		fmt.Println("unmarshal error", err)
	}
	fmt.Printf("%+v", c2)
}

// 序列化，就是将 结构体转为 2 进制流
//反序列化，就是将 2 进制流 转为 结构体,用于存储和传输 必须要传指针

//sonic json 序列化快
