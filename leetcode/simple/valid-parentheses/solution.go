package valid_parentheses

/*
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

	1 左括号必须用相同类型的右括号闭合。
	2 左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。
 */

// 有效的括号
func isValid(s string) bool {
	var topChar byte
	m := map[byte]byte{')':'(','}':'{',']':'['}
	b := []byte(s)
	stack := []byte{}

	for i := 0; i < len(s); i++ {
		if b[i] == '(' || b[i] == '{' || b[i] == '[' {
			stack = append(stack,b[i])
		} else {
			if len(stack) == 0 {
				return false
			}

			topChar = stack[len(stack) - 1]
			stack = stack[:len(stack) - 1]
			if m[b[i]] == '(' && topChar != '(' {
				return false
			}
			if m[b[i]] == '[' && topChar != '[' {
				return false
			}
			if m[b[i]] == '{' && topChar != '{' {
				return false
			}
		}
	}

	return len(stack) == 0
}
