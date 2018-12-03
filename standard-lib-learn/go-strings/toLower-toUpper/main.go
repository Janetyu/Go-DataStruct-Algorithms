package main

/*
ToLower 将字符串中的 Unicode 字符全部转换为相应的小写字符:

strings.ToLower(s string) string

ToUpper 将字符串中的 Unicode 字符全部转换为相应的大写字符:

strings.ToUpper(s string) string

ToTitle 将 s 中的所有字符修改为其 Title 格式
大部分字符的 Title 格式就是其 Upper 格式
只有少数字符的 Title 格式是特殊字符
这里的 ToTitle 主要给 Title 函数调用

func ToTitle(s string) string
 */

import (
	"fmt"
	"strings"
)

func main() {
	s := "This is a Go program!"
	fmt.Println(strings.ToLower(s)) //this is a go program!
	fmt.Println(strings.ToUpper(s)) //THIS IS A GO PROGRAM!
	fmt.Println(strings.ToTitle(s)) //THIS IS A GO PROGRAM!
}
