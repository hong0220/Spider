package main

import (
	"fmt"
	"os"
)

type SpiderSlaver struct {
}

func (spiderSlaver SpiderSlaver) work(job <-chan UrlSeed, result chan int) {
	for urlSeed := range job {
		fmt.Printf("当前goroutineID=%d,url=%s\n", GoroutineUti{}.GetGoroutineID(), urlSeed.url)

		// 发起网络请求
		var content = HttpUtil{}.Get(urlSeed.url)
		// fmt.Println(content)

		// 持久化
		f, err := os.Create("/Users/hongduoduo/Desktop/source/go/go_work/src/fast-boot-go/project/spider/data.txt")
		if err != nil {
			panic(err)
		}
		n3, err := f.WriteString(content)
		if err != nil {
			panic(err)
		}
		fmt.Printf("write %d bytes\n", n3)

		result <- 1
	}
}
