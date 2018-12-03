package main

/*
你可以使用 strings.TrimSpace(s) 来剔除字符串开头和结尾的空白符号;

如果你想要剔除指定字符，则可以使用 strings.Trim(s,cutset) 来将开头和结尾的 cutset 去除掉。

strings.Trim(s string cutset string) string

该函数的第二个参数可以包含任何字符，如果你只想剔除开头或者结尾的字符串，
则可以使用 TrimLeft 或者 TrimRight 来实现。

strings.TrimLeft(s string, cutset string) string
strings.TrumRight(s string, cutset string) string

TrimPrefix 删除 s 头部的 prefix 字符串
如果 s 不是以 prefix 开头，则返回原始 s

func TrimPrefix(s, prefix string) string

TrimSuffix 删除 s 尾部的 suffix 字符串
如果 s 不是以 suffix 结尾，则返回原始 s

func TrimSuffix(s, suffix string) string
 */

import (
	"fmt"
	"strings"
)

func main() {
	s := " ThisThis is a Go program!This "
	fmt.Println(s) //原生打印
	fmt.Println(strings.TrimSpace(s)) //剔除空格后的字符串

	s1 := "ThisThis is a Go program!This"
	fmt.Println(strings.Trim(s1, "This")) // is a Go program!
	fmt.Println(strings.TrimLeft(s1, "This"))//is a Go program!This
	fmt.Println(strings.TrimRight(s1, "This"))//ThisThis is a Go program!

	//注：TrimSuffix只是去掉s字符串结尾的suffix字符串，只是去掉１次，
	// 而TrimRight是一直去掉s字符串右边的字符串，只要有响应的字符串就去掉，
	// 是一个多次的过程，这也是二者的本质区别．

	s = "Hello 世界!"
	ts := strings.TrimPrefix(s, "Hello")
	fmt.Printf("%q\n", ts) // " 世界"

	s = "Hello 世界!!!!!"
	ts = strings.TrimSuffix(s, "!!!!")
	fmt.Printf("%q\n", ts) // " 世界"
}
