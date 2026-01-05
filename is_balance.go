package main

func isBalanced(s string) string {
	// Write your code here
	openBracket := map[rune]rune{'{': '}', '(': ')', '[': ']'}
	stack := make([]rune, 0)

	if len(s) == 1 {
		return "NO"
	}

	// {[()]}
	// ()(){}
	// {({})}
	for _, b := range s {
		switch b {
		case '(', '[', '{':
			stack = append(stack, b)

		case ']', ')', '}':
			if len(stack) == 0 {
				return "NO"
			}
			topBrack := stack[len(stack)-1]
			if openBracket[topBrack] != b {
				return "NO"
			}
			stack = stack[:len(stack)-1]
		}
	}

	if len(stack) > 0 {
		return "NO"
	}
	return "YES"
}
