package validparentheses

//Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
func isValid(s string) bool {
	//
	mapLeft := map[rune]rune{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	mapRight := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	//
	stack := make([]rune, 0)
	//
	for _, v := range s {
		_, ok := mapLeft[v]
		if ok {
			stack = append(stack, v)
		} else {
			if len(stack) > 0 {
				pop := stack[len(stack)-1]
				stack = stack[0 : len(stack)-1]
				expect := mapRight[v]
				if pop != expect {
					return false
				}
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}
