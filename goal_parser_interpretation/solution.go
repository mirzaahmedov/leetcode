package solution

import (
	"strings"
)

func interpret(command string) string {
	result := &strings.Builder{}
	value := &strings.Builder{}

	for i := 0; i < len(command); i++ {
		switch command[i] {
		case '(':
			value.WriteByte('(')
		case ')':
			value.WriteByte(')')
			if value.String() == "()" {
				result.WriteByte('o')
			} else if value.String() == "(al)" {
				result.WriteString("al")
			}
			value.Reset()
		case 'G':
			result.WriteByte('G')
		default:
			value.WriteByte(command[i])
		}
	}

	return result.String()
}
