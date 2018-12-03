package main

/*
DecodeLastRune(p []byte) (r rune, size int)
解码 []byte p 中最后一个 UTF-8 编码序列，返回该码值和长度。

DecodeLastRuneInString(s string) (r rune, size int)
解码 string s 中最后一个 UTF-8 编码序列，返回该码值和长度。

DecodeRune(p []byte) (r rune, size int)
解码 []byte p 中第一个 UTF-8 编码序列，返回该码值和长度。

DecodeRuneInString(s string) (r rune, size int)
解码 string s 中第一个 UTF-8 编码序列，返回该码值和长度。

EncodeRune(p []byte, r rune) int
将 rune r 的 UTF-8 编码序列写入 []byte p，并返回写入的字节数。p 满足足够的长度。

FullRune(p []byte) bool
检测 []byte p 是否包含一个完整 UTF-8 编码。

FullRuneInString(s string) bool
检测 string s 是否包含一个完整 UTF-8 编码。

RuneCount(p []byte) int
返回 []byte p 中的 UTF-8 编码的码值的个数。

RuneCountInString(s string) (n int)
返回 string s 中的 UTF-8 编码的码值的个数。

RuneLen(r rune) int
返回 rune r 编码后的字节数。

RuneStart(b byte) bool
检测字节 byte b 是否可以作为某个 rune 编码的第一个字节。

Valid(p []byte) bool
检测切片 []byte p 是否包含完整且合法的 UTF-8 编码序列。

ValidRune(r rune) bool
检测字符 rune r 是否包含完整且合法的 UTF-8 编码序列。

ValidString(s string) bool
检测字符串 string s 是否包含完整且合法的 UTF-8 编码序列。
 */

import (
	"unicode/utf8"
	"fmt"
)

func main() {
	// 35838 就是课的 unicode 码值
	utf8.DecodeLastRune([]byte("小韩说课")) // 返回 35838 3

	// 35838 就是课的 unicode 码值
	utf8.DecodeLastRuneInString("小韩说课") // 返回 35838 3

	// 23567 就是 小 的 unicode 码值
	utf8.DecodeRune([]byte("小韩说课")) // 返回 23567 3


	// 23567 就是 小 的 unicode 码值
	utf8.DecodeRuneInString("小韩说课") // 返回 23567 3

	buf := make([]byte, 3)
	n := utf8.EncodeRune(buf, '康')
	fmt.Println(buf, n) // 输出 [229 186 183] 3

	buf = []byte{229, 186, 183} // 康
	utf8.FullRune(buf) // 返回 true
	utf8.FullRune(buf[:2]) // 返回 false

	buf2 := "康" // 康
	utf8.FullRuneInString(buf2) // 返回 true
	utf8.FullRuneInString(buf2[:2]) // 返回 false

	buf3 := []byte("小韩说课")
	fmt.Println(len(buf3)) // 返回 12
	utf8.RuneCount(buf3) // 返回 4

	buf4 := "小韩说课"
	fmt.Println(len(buf4)) // 返回 12
	utf8.RuneCountInString(buf4) // 返回 4

	utf8.RuneLen('康') // 返回 3
	utf8.RuneLen('H') // 返回 1

	buf5 := "小韩说课"
	utf8.RuneStart(buf5[0]) // 返回 true
	utf8.RuneStart(buf5[1]) // 返回 false
	utf8.RuneStart(buf5[3]) // 返回 true

	valid := []byte("小韩说课")
	invalid := []byte{0xff, 0xfe, 0xfd}
	utf8.Valid(valid) // 返回 true
	utf8.Valid(invalid) // 返回 false

	valid2 := 'a'
	invalid2 := rune(0xfffffff)
	fmt.Println(utf8.ValidRune(valid2)) // 返回 true
	fmt.Println(utf8.ValidRune(invalid2)) // 返回 false

	valid3 := "小韩说课"
	invalid3 := string([]byte{0xff, 0xfe, 0xfd})
	fmt.Println(utf8.ValidString(valid3)) // 返回 true
	fmt.Println(utf8.ValidString(invalid3)) // 返回 false
}
