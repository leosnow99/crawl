package main

import (
	"fmt"
	"regexp"
)

func testRegexp() {
	str := "I love leosnow@gmail.com\n"
	re := regexp.MustCompile(`[0-9a-zA-Z]*@gmail.com`)
	findString:= re.FindString(str)
	fmt.Println(findString)
}
