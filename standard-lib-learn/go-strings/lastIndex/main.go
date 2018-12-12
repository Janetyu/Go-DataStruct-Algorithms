package main

/*
LastIndex 返回字符串 substr 在字符串 s 中最后出现位置的索引( substr 的第一个字符的索引)，-1 表示

strings.LastIndex(s string, substr string) int
*/

import (
	"fmt"
	"strings"
)

func main() {
	substr := "is"
	substr2 := "Hi"
	str1 := "This is a Go program!"
	fmt.Println(strings.LastIndex(str1, substr))  //输出5
	fmt.Println(strings.LastIndex(str1, substr2)) //输出-1
}
