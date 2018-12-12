package main

/*
HasPrefix判断字符串s是否以prefix开头

strings.HaxPrefix(s string, prefix string) bool
*/

import (
	"fmt"
	"strings"
)

func main() {
	pre := "Thi"
	str1 := "This is a Go project"
	fmt.Println(strings.HasPrefix(str1, pre))
}
