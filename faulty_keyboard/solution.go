package solution

func finalString(s string) string {
	value := []byte{}
	for i := 0; i < len(s); i++ {
		if s[i] == 'i' {
			value = reverse(value)
		} else {
			value = append(value, s[i])
		}
	}
	return string(value)
}
func reverse(s []byte) []byte {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return s
}
