package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	now := time.Now()
	fmt.Printf("current time: %v\n", now)
	year := now.Year()
	month := now.Month()

	fmt.Printf("year: %v\n", year)
	fmt.Printf("month: %v\n", month)

	timestamp1 := now.Unix()
	timestamp2 := now.UnixNano()
	fmt.Printf("timestamp1 : %v\n", timestamp1)
	fmt.Printf("timestamp2 : %v\n", timestamp2)
	//定时器
	//ticker := time.Tick(time.Second)
	//for i := range ticker {
	//	fmt.Println(i)
	//}

	//time   format
	// 24小时制
	fmt.Println(now.Format("2006-01-02 15:04:05"))
	//12 小时制
	fmt.Println(now.Format("2006-01-02 03:04:05 PM"))
	//解析字符串格式的时间
	fmt.Println(now)
	//加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	//按照指定时区和指定格式解析字符串时间
	timeObj, err := time.ParseInLocation("2006/01/02 15:04:05", "2019/08/04 14:15:20", loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj)
	fmt.Println(timeObj.Sub(now))

	log.Println("这是一条很普通的日志。")
	v := "很普通的"
	log.Printf("这是一条%s日志。\n", v)
	log.Fatalln("这是一条会触发fatal的日志。")
	log.Panicln("这是一条会触发panic的日志。")
}
