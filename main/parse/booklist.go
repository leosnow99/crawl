package parse

import (
	"crawl/main/engine"
	"regexp"
)

const BookListRe = `<a href="([^"]+)" title="([^"]+)"`

func BookList(contents []byte) engine.ParseResult {
	bookListReg := regexp.MustCompile(BookListRe)
	
	matches := bookListReg.FindAllSubmatch(contents, -1)
	
	result := engine.ParseResult{}
	
	for _, match := range matches {
		result.Items = append(result.Items, match[2])
		bookName := string(match[2])
		result.Requests = append(result.Requests, engine.Request{
			Url: string(match[1]),
			ParseFunc: func(content []byte) engine.ParseResult {
				return BookDetail(content, bookName)
			},
		})
	}
	
	return result
}
