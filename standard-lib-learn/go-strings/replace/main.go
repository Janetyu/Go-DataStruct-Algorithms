package main

/*
Replace 用于将字符串s 中的前 n 个字符串 old 替换为字符串 new ，并返回一个新的字符串，
如果 n = -1 则替换所有字符串 old 为字符串 new :

strings.Replace(s string, old string, new string, n int) string

Replacer 根据一个替换列表执行替换操作
type Replacer struct {
	Replace(s string) string
	WriteString(w io.Writer, s string) (n int, err error)
}

NewReplacer 通过“替换列表”创建一个 Replacer 对象。
按照“替换列表”中的顺序进行替换，只替换非重叠部分。
如果参数的个数不是偶数，则抛出异常。
如果在“替换列表”中有相同的“查找项”，则后面重复的“查找项”会被忽略
func NewReplacer(oldnew ...string) *Replacer

Replace 返回对 s 进行“查找和替换”后的结果
Replace 使用的是 Boyer-Moore 算法，速度很快
func (r *Replacer) Replace(s string) string

WriteString 对 s 进行“查找和替换”，然后将结果写入 w 中
func (r *Replacer) WriteString(w io.Writer, s string) (n int, err error)
*/

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	str1 := "我爱你熙中国熙"
	fmt.Println(strings.Replace(str1, "熙", "", 1))  //我爱你中国熙
	fmt.Println(strings.Replace(str1, "熙", "", -1)) //我爱你中国

	srp := strings.NewReplacer("Hello", "你好", "World", "世界", "!", "！")
	s := "Hello World!Hello World!hello world!"
	rst := srp.Replace(s)
	fmt.Println(rst) // 你好 世界！你好 世界！hello world！

	wl2 := []string{"Hello", "Hi", "Hello", "你好"}
	srp2 := strings.NewReplacer(wl2...)
	s2 := "Hello World! Hello World! hello world!"
	rst2 := srp2.Replace(s2)
	fmt.Println(rst2) // Hi World! Hi World! hello world!

	wl3 := []string{"Hello", "你好", "World", "世界", "!", "！"}
	srp3 := strings.NewReplacer(wl3...)
	s3 := "Hello World!Hello World!hello world!"
	srp3.WriteString(os.Stdout, s3) // 你好 世界！你好 世界！hello world！
}
