package main

/*
Contains 判断字符串 s 是否包含 substr :

strings.Contains(s string, substr string) bool

ContainsAny 判断字符串 s 中是否包含 chars 中的任何一个字符
如果 chars 为空，则返回 false

func ContainsAny(s, chars string) bool

ContainsRune 判断字符串 s 中是否包含字符 r

func ContainsRune(s string, r rune) bool
*/

import (
	"fmt"
	"strings"
)

func main() {
	substr := "is"
	str1 := "This is a Go project!"
	fmt.Println(strings.Contains(str1, substr))

	s := "Hello,世界!"
	b := strings.ContainsAny(s, "abc")
	fmt.Println(b) // false
	b = strings.ContainsAny(s, "def")
	fmt.Println(b) // true

	b := strings.ContainsRune(s, '\n')
	fmt.Println(b) // false
	b = strings.ContainsRune(s, '界')
	fmt.Println(b) // true
}
