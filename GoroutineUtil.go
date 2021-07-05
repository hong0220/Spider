package main

import (
	"bytes"
	"runtime"
	"strconv"
)

type GoroutineUti struct {
}

/**
 * 获取goroutineID
 * 堆栈信息：
 * goroutine 1 [running]:
 * main.main()
 *     C:/Users/zuma/Desktop/gowork/src/github.com/0990/gotool/main.go:10 +0x74
 */
func (goroutineUti GoroutineUti) GetGoroutineID() uint64 {
	b := make([]byte, 64)
	runtime.Stack(b, false)

	// 注意最后一个空格
	b = bytes.TrimPrefix(b, []byte("goroutine "))
	b = b[:bytes.IndexByte(b, ' ')]

	// 以10进制方式解析b，保存为int64类型
	n, _ := strconv.ParseUint(string(b), 10, 64)
	return n
}
