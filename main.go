package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

/**
* @author leo
* @date 18-9-9 下午14:44
 */

func main() {

	// Colly的主要实体是一个Collector对象。Collector在收集器作业运行时管理网络通信并负责执行附加的回调。
	// 初始化Collector
	c := colly.NewCollector()
	// 可以将不同类型的回调函数添加到Collector中，来控制收集作业或者处理网络信息
	// 查找和访问所有链接
	// OnResponse如果收到的内容是HTML则立即调用

	c.OnHTML("a[href]", func(element *colly.HTMLElement) {
		element.Request.Visit(element.Attr("href"))
	})

	// 在请求Request之前调用
	c.OnRequest(func(request *colly.Request) {
		fmt.Println("正在访问--->>>", request.URL)
	})

	// 收到回复Response后调用
	c.OnResponse(func(response *colly.Response) {
		fmt.Println("响应response--->>>", response.StatusCode)
	})

	c.Visit("http://www.zhihu.com/")
}
