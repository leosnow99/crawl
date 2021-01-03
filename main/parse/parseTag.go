package parse

import (
	"crawl/main/engine"
	"fmt"
	"regexp"
)

const regexpStr = `<a href="([^"]+)" class="tag">([^"]+)</a>`

func Content(content []byte) (result engine.ParseResult) {
	reg := regexp.MustCompile(regexpStr)
	
	for _, match := range reg.FindAllSubmatch(content, -1) {
		fmt.Printf("url: %s\n", "https://book.douban.com"+string(match[1]))
		result.Items = append(result.Items, match[2])
		result.Requests = append(result.Requests, engine.Request{
			Url:       "https://book.douban.com" + string(match[1]),
			ParseFunc: BookList,
		})
	}
	return
}
