package main

import (
	"fmt"
)

func main() {
	fmt.Println(countAndSay(1))
	fmt.Println(countAndSay(2))
	fmt.Println(countAndSay(3))
	fmt.Println(countAndSay(4))
}
func countAndSay(n int) string {
	if n < 2 {
		return "1"
	}

	return count(2, n, "1")
}

func count(n int, max int, res string) string {
	if n == max {
		return say(res)
	}

	r := say(res)

	return count(n+1, max, r)
}
func say(s string) string {
	r := make([]byte, 0)

	c, t := byte(0), byte(100)
	for i := 0; i < len(s); i++ {
		if s[i] < 48 || s[i] > 59 {
			continue
		}
		if s[i] == t {
			c++
		} else {
			if t != 100 {
				r = append(r, c+48, t)
			}

			c = 1
			t = s[i]
		}
	}

	if c > 0 {
		r = append(r, c+48, t)
	}

	return string(r)
}
