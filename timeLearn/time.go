package timeLearn

import (
	"fmt"
	"time"
)

const (
	DATE = "2006-01-02"
	TIME = "2006-01-02 15:04:05"
)

func LearnTimeFunction() {

	// 获取当前时间
	now := time.Now()
	// 打印当前时间
	fmt.Println("now:", now)
	// 转时间戳
	nowUnix := now.Unix()
	fmt.Println("nowUnix", nowUnix)

	//让程序在某个地方暂停多长时间
	time.Sleep(time.Second * 3)
	// 求两个时间的差值 结果为 Duration 类型
	sub := time.Now().Sub(now)
	// Duration 类型可以有不同的精度显示
	// 差多少分钟
	fmt.Println("sub ", sub.Minutes())
	// 差多少秒
	fmt.Println("sub ", sub.Seconds())
	// 差多少毫秒
	fmt.Println("sub ", sub.Milliseconds())

	// 计算时间的加法 time.Time + duration 类型
	now2 := now.Add(sub)
	// Time 转成时间戳
	fmt.Println("now2: ", now2.Unix())
	//时间 + 时间格式 转成字符串的可读形式 Time -> string
	dateFormat := now.Format(DATE)
	fmt.Println("date string", dateFormat)
	timeFormat := now.Format(TIME)
	fmt.Println("time string", timeFormat)
	// Parse 函数 string 转 Time 类型
	parseDate, err := time.Parse(DATE, dateFormat)
	if err != nil {
		fmt.Println("parse error", err)
	}
	fmt.Println("parseDate", parseDate)

	parseTime, err := time.Parse(TIME, timeFormat)
	if err != nil {
		fmt.Println("parse error", err)
	}
	fmt.Println("parseTime", parseTime)
	// Parse 的时候确定在哪一个时区
	east8, _ := time.LoadLocation("Asia/Shanghai")
	parseTimeInLocation, err := time.ParseInLocation(TIME, timeFormat, east8)
	if err != nil {
		fmt.Println("parse error", err)
	}
	fmt.Println("parseTimeInLocation", parseTimeInLocation)
}
