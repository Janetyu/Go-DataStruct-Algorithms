package main

/*
Join 用于将元素类型为 string 的 slice 使用分割符号来拼接组成一个字符串:

strings.Join(a []string, sep string) string
*/

import (
	"fmt"
	"strings"
)

func main() {
	var s []string = []string{"I", "love", "you"}
	fmt.Println(strings.Join(s, " ")) //I love you
}
