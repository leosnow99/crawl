package parse

import (
	"crawl/main/engine"
	"crawl/main/model"
	"regexp"
	"strconv"
)

var authorRe = regexp.MustCompile(`<span class="pl"> 作者</span>:[\d\D]*?<a.*?>([^<]+)</a>`)
var publisherRe = regexp.MustCompile(`<span class="pl">出版社:</span> ([^<]+)<br`)
var pageRe = regexp.MustCompile(`<span class="pl">页数:</span> ([^<]+)<br`)
var priceRe = regexp.MustCompile(`<span class="pl">定价:</span>([^<]+)<br`)
var scoreRe = regexp.MustCompile(`<strong class="ll rating_num " property="v:average">([^<]+)</strong>`)
var infoRe = regexp.MustCompile(`<div class="intro">[\d\D]*?<p>([^<]+)</p></div>`)

func BookDetail(content []byte, name string) engine.ParseResult {
	bookDetail := model.BookDetail{
		BookName:  name,
		Author:    extraStr(content, authorRe),
		Publisher: extraStr(content, publisherRe),
		Price:     extraStr(content, priceRe),
		Score:     extraStr(content, scoreRe),
		Info:      extraStr(content, infoRe),
	}
	page, err := strconv.Atoi(extraStr(content, pageRe))
	if err != nil {
		bookDetail.Page = 0
	} else {
		bookDetail.Page = page
	}
	
	result := engine.ParseResult{
		Items: []interface{}{
			bookDetail,
		},
	}
	return result
}

func extraStr(contents []byte, re *regexp.Regexp) string {
	matches := re.FindSubmatch(contents)
	
	if len(matches) >= 2 {
		return string(matches[1])
	}
	return ""
}
