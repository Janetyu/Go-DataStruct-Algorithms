package main

/*
Repeat 用于重复 count 次字符串 s 并返回一个新的字符串:

strings.Repeat(s string, count int) string
*/

import (
	"fmt"
	"strings"
)

func main() {
	s := "This is a go program!"
	fmt.Println(strings.Repeat(s, 2)) //This is a go program!This is a go program!
}
