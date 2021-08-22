package search

import (
	"log"
	"sync"
)

// 注册用于搜索的匹配器映射
var matchers = make(map[string]Matcher)

// Run执行搜索逻辑
func Run(searchTerm string) {
	// 构造一个witGroup，以便处理所有数据源
	var waitGroup sync.WaitGroup

	// 获取需要搜索的数据源列表
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}
	// 设置需要等待处理
	// 每个数据源的goroutine的数量
	waitGroup.Add(len(feeds))

	// 无缓冲通道，接受匹配后的结果
	results := make(chan *Result)

	// 为每个数据源启动一个goroutine来查找结果
	for _, feed := range feeds {
		// 获取一个匹配器用于查找
		matcher, exists := matchers[feed.Type]
		// 如果不存在则使用默认匹配器
		if !exists {
			matcher = matchers["default"]
		}

		// 启动一个goroutine来执行搜索
		go func(matcher Matcher, feed *Feed) {
			// 搜索数据源的数据，将匹配结果输出到results通道
			Match(matcher, feed, searchTerm, results)
			waitGroup.Done()
		}(matcher, feed)
	}

	// 启动一个goroutine来监控是否所有工作都完成了
	go func() {
		waitGroup.Wait()

		// 关闭通道，通知Display函数
		close(results)
	}()

	// 打印函数，在收到完成通知后结束
	Display(results)
}
