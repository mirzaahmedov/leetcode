package solution

import (
	"fmt"
	"strings"
)

func main() {
	// fmt.Println(simplifyPath("/home/"))
	// fmt.Println(simplifyPath("/../"))
	// fmt.Println(simplifyPath("/home/foo"))
	fmt.Println(simplifyPath("/..."))
	fmt.Println(simplifyPath("/.../"))
	fmt.Println(simplifyPath("/..hidden"))
	fmt.Println(simplifyPath("/.....hidden"))
	fmt.Println(simplifyPath("/hello../world"))
}
func simplifyPath(path string) string {
	var (
		level     = 0
		prev      byte
		items     []string
		curr      string
		dotsCount = 0
		result    = "/"
	)

	for i := 0; i < len(path); i++ {
		switch path[i] {
		case '/':
			if prev != '/' && dotsCount == 0 && curr != "" {
				items = append(items, curr)
				level += 1
			} else if dotsCount == 2 && curr == "" {
				if level > 0 {
					level -= 1
					if len(items) > 0 {
						items = items[0 : len(items)-1]
					}
				}
			} else if dotsCount > 0 && curr != "" {
				items = append(items, curr+strings.Repeat(".", dotsCount))
				level += 1
			} else if dotsCount > 2 {
				items = append(items, curr+strings.Repeat(".", dotsCount))
				level += 1
			}
			dotsCount = 0
			curr = ""

		case '.':
			dotsCount += 1
		default:
			if dotsCount > 0 {
				curr += strings.Repeat(".", dotsCount)
			}
			curr += string(path[i])

			dotsCount = 0
		}
		prev = path[i]
	}

	if dotsCount == 2 && len(items) > 0 && curr == "" {
		items = items[0 : len(items)-1]
	} else if dotsCount > 2 {
		curr += strings.Repeat(".", dotsCount)
	} else if curr != "" && dotsCount > 0 {
		curr += strings.Repeat(".", dotsCount)
	}

	if len(curr) > 0 {
		items = append(items, curr)
	}

	if len(items) == 0 {
		return result
	}

	for i := 0; i < len(items); i++ {
		result += items[i] + "/"
	}

	return result[0 : len(result)-1]
}
