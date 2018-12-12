package main

/*
HasSuffix 判断字符串 s 是否以 suffix 结尾:

strings.HasSuffix(s string, suffix string) bool
*/

import (
	"fmt"
	"strings"
)

func main() {
	suffix := "ct!"
	str1 := "This is a Go project!"
	fmt.Println(strings.HasSuffix(str1, suffix))
}
