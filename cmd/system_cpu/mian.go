/*
 * @Description:
 * @Author: shangke
 * @Date: 2024/10/6
 * @LastEditTime: 2024/10/6
 * @LastEditors: shangke
 */
package main

import (
	"fmt"
	"runtime"
)

//识别system cpu

func main() {
	// 获取操作系统和 CPU 架构
	os := runtime.GOOS
	arch := runtime.GOARCH

	// 获取当前程序运行的 Go 版本
	goVersion := runtime.Version()

	// 获取当前可用的处理器数量
	numCPU := runtime.NumCPU()

	// 获取当前 Goroutine 的数量
	numGoroutine := runtime.NumGoroutine()

	// 打印信息
	fmt.Println("操作系统:", os)
	fmt.Println("CPU 架构:", arch)
	fmt.Println("Go 版本:", goVersion)
	fmt.Println("可用处理器数量:", numCPU)
	fmt.Println("当前 Goroutine 数量:", numGoroutine)

	// 设置最大并发数
	runtime.GOMAXPROCS(runtime.NumCPU())
	fmt.Println("设置的最大并发数:", runtime.GOMAXPROCS(-1))
}
