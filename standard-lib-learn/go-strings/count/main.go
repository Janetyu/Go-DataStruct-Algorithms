package main

/*
Count 用于计算字符串 substr 在字符串 s 中出现的非重叠次数:

strings.Count(s string, substr string) int
*/

import (
	"fmt"
	"strings"
)

func main() {
	substr := "is"
	s := "This is a go program!"
	fmt.Println(strings.Count(s, substr)) //2
}
