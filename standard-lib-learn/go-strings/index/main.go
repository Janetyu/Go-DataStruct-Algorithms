package main

/*
Index 返回字符串 substr 在字符串 s 中的索引( substr 的第一个字符的索引)，
-1 表示字符串 s 不包含字符串 substr :

strings.Index(s string, sbustr string) int

IndexAny 返回字符串 chars 中的任何一个字符在字符串 s 中第一次出现的位置
如果找不到，则返回 -1，如果 chars 为空，则返回 -1

func IndexAny(s, chars string) int
*/

import (
	"fmt"
	"strings"
)

func main() {
	substr := "is"
	substr2 := "Hi"
	str1 := "This is a Go program!"
	fmt.Println(strings.Index(str1, substr))  //输出2
	fmt.Println(strings.Index(str1, substr2)) //输出-1

	s := "Hello,世界! Hello!"
	i := strings.IndexAny(s, "abc")
	fmt.Println(i) // -1
	i = strings.IndexAny(s, "dof")
	fmt.Println(i) // 1
}
