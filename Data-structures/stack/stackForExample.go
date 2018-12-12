package stack

import "reflect"

func isValid(s string) bool {
	var topChar byte
	m := map[byte]byte{')': '(', '}': '{', ']': '['}
	b := []byte(s)
	stack := CreatedefaultStack(reflect.Uint8)

	for i := 0; i < len(s); i++ {
		if b[i] == '(' || b[i] == '{' || b[i] == '[' {
			stack.Push(b[i])
		} else {
			if stack.IsEmpty() {
				return false
			}

			topChar = stack.Pop().(uint8)
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

	return stack.IsEmpty()
}
