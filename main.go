package main

import (
	"bufio"
	"os"
	"strings"
	"fmt"
)
// 第3个亿！！！
func main() {
	r := bufio.NewReader(os.Stdin)

	for {
		c, err := r.ReadString('\n')
		if err == nil {
			c = strings.Replace(c, "吗", "", -1)
			c = strings.Replace(c, "？", "！", -1)
			c = strings.Replace(c, "？", "！", -1)
			fmt.Print(c)
		}
	}
}
