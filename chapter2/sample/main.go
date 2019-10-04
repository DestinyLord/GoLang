package main

import (
	_ "./matchers"
	"./search"
	"log"
	"os"
)

// init在main之前调用
func init()  {
	// 将日志输出到标准输出
	log.SetOutput(os.Stdout)
}

func main() {
	// 使用特定的项做搜索
	search.Run("president")
}