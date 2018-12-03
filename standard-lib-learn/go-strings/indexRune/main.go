package main

/*
如果 r 是非 ASCII 编码的字符，建议使用以下函数来对字符进行定位:

strings.IndexRune(s string, r rune) int
 */

import (
	"fmt"
	"strings"
)

func main() {
	substr := '中' // int32
	substr2 := '天'
	str1 := "我爱你中国"
	fmt.Println(strings.IndexRune(str1, substr)) //输出9
	fmt.Println(strings.IndexRune(str1, substr2))//输出-1

	substr3 := "中"
	substr4 := "天"
	str2 := "我爱你𤋮中国"

	//常见的汉字在utf8中占3个字节，某些特殊的汉字(如𤋮)在utf8中占四个字节，所以输出结果是13而不是12
	fmt.Println(strings.Index(str2, substr3)) //输出13
	fmt.Println(strings.Index(str2, substr4))//输出-1
}