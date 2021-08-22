package main

import (
	"chapter2-sample/search"
	"log"
	"os"

	// _可以导入一个包，但是并不使用它，这样做的原因是为了初始化这个包，
	// 导入会自动执行包中所有文件的init()函数，
	// 在这里这样做的原因是为了调用matchers包中rss.go文件里的init()，注册RSS匹配器
	_ "chapter2-sample/matchers"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}
