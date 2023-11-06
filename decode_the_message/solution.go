package solution

import (
	"strings"
)

func decodeMessage(key string, message string) string {
	mappings := make(map[byte]byte)
	count := 0

	for i := 0; i < len(key) && count < 26; i++ {
		if key[i] == ' ' {
			continue
		}
		if _, ok := mappings[key[i]]; !ok {
			mappings[key[i]] = byte(count + 97)
			count++
		}
	}
	decoded := strings.Builder{}

	for i := 0; i < len(message); i++ {
		if l, ok := mappings[message[i]]; ok {
			decoded.WriteByte(l)
		} else {
			decoded.WriteByte(message[i])
		}
	}
	return decoded.String()
}
