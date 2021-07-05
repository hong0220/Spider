package main

import (
	"fmt"
	"time"
)

func main() {
	// 启动时间
	fmt.Printf("spider start %s\n", time.Now())

	jobs := make(chan UrlSeed, 100)
	result := make(chan int, len(jobs))

	// master启动
	count := 2
	SpiderMaster{}.start(jobs, result, count)

	// 等待
	for i := 0; i < count; i++ {
		<-result
	}

	fmt.Printf("spider stop %s\n", time.Now())
}
