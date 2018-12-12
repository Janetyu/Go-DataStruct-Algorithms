package main

import "unicode"

/*
Is(rangeTab *RangeTable, r rune) bool
检测 rune r 是否在 rangeTable 指定的字符范围内。
rangeTable 一个 Unicode 码值集合，通常使用 unicode 包中定义的集合

In(r rune, ranges …*RangeTable) bool
检测 rune r 是否在多个 rangeTable 指定的字符范围内。
rangeTable 一个 Unicode 码值集合，通常使用 unicode 包中定义的集合

IsOneOf(ranges []*RangeTable, r rune) bool
检测 rune r 是否在 rangeTable ranges 指定的字符范围内。与 In 功能类似，推荐使用 In

IsSpace(r rune) bool
检测字符 rune r 是否是空白字符。在Latin-1字符空间中，空白字符为：
'\t', '\n', '\v', '\f', '\r', ' ', U+0085 (NEL), U+00A0 (NBSP)
其它的空白字符请参见策略Z和属性Pattern_White_Space

IsDigit(r rune) bool
检测字符 rune r 是否是十进制数字字符

IsNumber(r rune) bool
检测字符 rune r 是否是 Unicode 数字字符

IsLetter(r rune) bool
检测一个字符 rune r 是否是字母

IsGraphic(r rune) bool
一个字符 rune r 是否是 unicode 图形字符。图形字符包括字母、标记、数字、符号、标点、空白

IsControl(r rune) bool
检测一个字符 rune r 是否是 unicode 控制字符

IsMark(r rune) bool
检测一个字符 rune r 是否是标记字符

IsPrint(r rune) bool
检测一个字符 rune r 是否是的可打印字符，基本与图形字符一致，除ASCII空白字符U+0020

IsPunct(r rune) bool
检测一个字符 rune r 是否是 unicode标点字符

IsSymbol(r rune) bool
检测一个字符 rune r 是否是 unicode 符号字符

IsLower(r rune) bool
检测一个字符 rune r 是否是小写字母

IsUpper(r rune) bool
检测一个字符 rune r 是否是大写字母

IsTitle(r rune) bool
检测一个字符 rune r 是否是Title字符。
大部分字符的 Title 格式就是其大写格式，少数字符的 Title 格式是特殊字符，例如 ᾏᾟᾯ。

To(_case int, r rune) rune
将字符 rune r 转换为指定的格式，格式_case支持：
unicode.UpperCase、unicode.LowerCase、unicode.TitleCase

ToLower(r rune) rune
将字符 rune r 转换为小写。

func (SpecialCase) ToLower
将字符 rune r 转换为小写。优先使用映射表 SpecialCase。

ToUpper(r rune) rune
将字符 rune r 转换为大写。

func (SpecialCase) ToUpper
将字符 rune r 转换为大写。优先使用映射表 SpecialCase。

ToTitle(r rune) rune
将字符 rune r 转换为 Title 字符。

func (SpecialCase) ToTitle
将字符 rune r 转换为 Title 字符。优先使用映射表 SpecialCase。

SimpleFold(r rune) rune
在 unicode 标准字符映射中查找与 rune r 互相对应的 unicode 码值。向码值大的方向循环查找。
互相对应指的是同一个字符可能出现的多种写法。
*/

func main() {
	// 判断字符是否出现在汉字集合中
	unicode.Is(unicode.Scripts["Han"], 'k') // 返回 false
	unicode.Is(unicode.Scripts["Han"], '康') // 返回 true

	// 检测 rune r 是否在多个 rangeTable 指定的字符范围内
	unicode.In('康', unicode.Scripts["Han"], unicode.Scripts["Latin"]) // 返回 true
	unicode.In('k', unicode.Scripts["Han"], unicode.Scripts["Latin"]) // 返回 true

	// 检测字符 rune r 是否是十进制数字字符
	unicode.IsDigit('9') // 返回 true
	unicode.IsDigit('k') // 返回 false

	// 检测一个字符 rune r 是否是字母
	unicode.IsLetter('9') // 返回 false
	unicode.IsLetter('k') // 返回 true

	// 一个字符 rune r 是否是 unicode 图形字符。图形字符包括字母、标记、数字、符号、标点、空白。
	unicode.IsGraphic('9') // 返回 true
	unicode.IsGraphic(',') // 返回 true

	// 检测一个字符 rune r 是否是 unicode标点字符。
	unicode.IsPunct('9') // 返回 false
	unicode.IsPunct(',') // 返回 true

	// 检测一个字符 rune r 是否是小写字母。
	unicode.IsLower('h') // 返回 true
	unicode.IsLower('H') // 返回 false

	// 检测一个字符 rune r 是否是大写字母。
	unicode.IsUpper('h') // 返回 false
	unicode.IsUpper('H') // 返回 true

	unicode.IsTitle('ᾯ') // 返回 true
	unicode.IsTitle('h') // 返回 false
	unicode.IsTitle('H') // 返回 true

	unicode.To(unicode.UpperCase, 'h') // 返回 H

	// 将字符 rune r 转换为小写。
	unicode.ToLower('H') // 返回 h

	// 映射表 SpecialCase 是特定语言环境下大小写的映射表。主要应用于一些欧洲字符，
	// 例如土耳其 TurkishCase。
	unicode.TurkishCase.ToLower('İ') // 返回 i

	// 将字符 rune r 转换为大写。
	unicode.ToUpper('h') // 返回 H

	// 映射表 SpecialCase 是特定语言环境下大小写的映射表。主要应用于一些欧洲字符，
	// 例如土耳其 TurkishCase。
	unicode.TurkishCase.ToUpper('i') // 返回 İ

	// 将字符 rune r 转换为 Title 字符。
	unicode.ToTitle('h') // 返回 H

	// 映射表 SpecialCase 是特定语言环境下大小写的映射表。
	// 主要应用于一些欧洲字符，例如土耳其 TurkishCase。
	unicode.TurkishCase.ToTitle('i') // 返回 İ

	// 在 unicode 标准字符映射中查找与 rune r 互相对应的 unicode 码值
	unicode.SimpleFold('H') // 返回 h
	unicode.SimpleFold('Φ') // 返回 φ
}
