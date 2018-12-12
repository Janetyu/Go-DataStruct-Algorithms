package main

/*
utf-8 是变长字符集，英文标点占用1个字节，中文占用3个字节。
*/

import "fmt"

func main() {
	len := len("Hank康")
	fmt.Println(len) // 7
}
