package model

import "strconv"

type BookDetail struct {
	BookName  string
	Author    string
	Publisher string
	Page      int
	Price     string
	Score     string
	Info      string
}

func (b BookDetail) String() string {
	return "书名: " + b.BookName + "作者: " + b.Author + " 出版社: " + b.Publisher + " 页数: " + strconv.Itoa(b.
		Page) + " 价格: " + b.Price + " 评分: " + b.Score + " \n简介: " + b.Info
}
