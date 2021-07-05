package main

type SpiderMaster struct {
}

func (spiderMaster SpiderMaster) start(jobs chan UrlSeed, result chan int, count int) {
	// 1个协程
	for i := 0; i < 1; i++ {
		// slaver启动
		go SpiderSlaver{}.work(jobs, result)
	}

	// 1个任务
	for i := 0; i < count; i++ {
		urlSeed := UrlSeed{url: "https://tech.meituan.com/", deep: 1}
		jobs <- urlSeed
	}
}
