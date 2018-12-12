package main

/*
strings.Split(s,sep)用于自定义分割符号来对指定字符串进行分割，同样返回slice。

因为这 2 个函数都会返回 slice，所以习惯使用 for-range 循环来对其进行处理

strings.Split(s string, sep string) []string

SplitN 以 sep 为分隔符，将 s 切分成多个子串，结果中不包含 sep 本身
如果 sep 为空，则将 s 切分成 Unicode 字符列表。
如果 s 中没有 sep 子串，则将整个 s 作为 []string 的第一个元素返回
参数 n 表示最多切分出几个子串，超出的部分将不再切分。
如果 n 为 0，则返回 nil，如果 n 小于 0，则不限制切分个数，全部切分

func SplitN(s, sep string, n int) []string

SplitAfterN 以 sep 为分隔符，将 s 切分成多个子串，结果中包含 sep 本身
如果 sep 为空，则将 s 切分成 Unicode 字符列表。
如果 s 中没有 sep 子串，则将整个 s 作为 []string 的第一个元素返回
参数 n 表示最多切分出几个子串，超出的部分将不再切分。
如果 n 为 0，则返回 nil，如果 n 小于 0，则不限制切分个数，全部切分

func SplitAfterN(s, sep string, n int) []string
*/

import (
	"fmt"
	"strings"
)

func main() {
	s := "This is a Go program!"
	result := strings.Split(s, "Go")
	fmt.Printf("%v\n", result) //[This is a  program!]
	for _, value := range result {
		fmt.Printf("%s\n", value)
	}

	s = "Hello, 世界! Hello!"
	ss := strings.SplitN(s, " ", 2)
	fmt.Printf("%q\n", ss) // ["Hello," "世界! Hello!"]
	ss = strings.SplitN(s, " ", -1)
	fmt.Printf("%q\n", ss) // ["Hello," "世界!" "Hello!"]
	ss = strings.SplitN(s, "", 3)
	fmt.Printf("%q\n", ss) // ["H" "e" "llo, 世界! Hello!"]

	s = "Hello, 世界! Hello!"
	ss = strings.SplitAfterN(s, " ", 2)
	fmt.Printf("%q\n", ss) // ["Hello, " "世界! Hello!"]
	ss = strings.SplitAfterN(s, " ", -1)
	fmt.Printf("%q\n", ss) // ["Hello, " "世界! " "Hello!"]
	ss = strings.SplitAfterN(s, "", 3)
	fmt.Printf("%q\n", ss) // ["H" "e" "llo, 世界! Hello!"]
}
