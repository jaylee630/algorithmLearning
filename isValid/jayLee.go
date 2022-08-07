package isValid

/**
 * 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串 s ，判断字符串是否有效。
 * 有效字符串需满足：
 * - 左括号必须用相同类型的右括号闭合。
 * - 左括号必须以正确的顺序闭合。
 */
func isValid(s string) bool {
	h := map[byte]byte{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	var stack []byte
	for i := 0; i < len(s); i++ {
		sl := len(stack)
		if sl == 0 || s[i] == '(' || s[i] == '{' || s[i] == '[' {
			stack = append(stack, s[i]) // 压栈
			continue
		}

		var bottom byte
		bottom, stack = stack[sl-1], stack[:sl-1] // 推栈
		if h[bottom] != s[i] {
			return false
		}
	}
	if len(stack) != 0 {
		return false
	}
	return true
}
