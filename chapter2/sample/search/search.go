package search

import (
	"log"
	"sync"
)

/* 注册用于搜索的匹配器的映射
   这个变量没有定义在任何函数作用域内，所以会被当成包级变量
   使用 var 声明，声明为 Matcher 类型的映射(map)，
   这个映射以string类型值作为键，Matcher类型值作为映射后的值。
*/
var matchers = make(map[string]Matcher)

// Run执行搜索逻辑
func Run(searchTerm string)  {
	/* 获取需要搜索的数据源列表
	   Go语言允许一个函数返回多个值，
	   例如 RetrieveFeeds 函数，返回一个值和一个错误值
	   如果发生了错误，永远不要使用该函数返回的另外一个值（也有允许同时返回数据和错误的函数，
	   但自己实现的函数，需要遵守这个原则，保持含义足够明确），
	   这时必须忽略另外一个值，否则程序会产生更多的错误，甚至崩溃
	 */
	feeds, err := RetrieveFeeds()
	if err != nil {
		log.Fatal(err)
	}

	// 创造一个无缓冲的通道，接收匹配后的结果
	results := make(chan *Result)

	// 构造一个 waitGroup ，以便处理所有的数据源
	var waitGroup sync.WaitGroup

	// 设置需要等待处理
	// 每个数据源的 goroutine 的数量
	waitGroup.Add(len(feeds))

	// 为每个数据源启动一个 goroutine 来查找结果
	for _, feed := range feeds {
		// 获取一个匹配器用于查找
		matcher, exists := matchers[feed.Type]
		if !exists {
			matcher = matchers["default"]
		}

		// 启动一个 goroutine 来执行搜索
		go func(matcher Matcher, feed *Feed) {
			Match(matcher, feed, searchTerm, results)
			waitGroup.Done()
		}(matcher, feed)
	}

	// 启动一个 goroutine 来监控是否所有的工作都做完了
	go func() {
		// 等待所有任务完成
		waitGroup.Wait()

		// 用关闭通道的方式，通知 Display 函数可以退出函数了
		close(results)
	}()

	// 启动函数，显示返回结果，并且在最后一个结果显示完后返回
	Display(results)
}