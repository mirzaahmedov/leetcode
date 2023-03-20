// https://leetcode.com/problems/longest-substring-without-repeating-characters/?envType=study-plan&id=algorithm-i

package main

import (
  "math"
)

func lengthOfLongestSubstring(s string) int {
  if len(s) < 2 {
    return len(s)
  }

  memo := map[byte]int{}

  i := 0
  max := 0

  for j := 0; j < len(s) && len(s) - i > max; j++ {
    if index, ok := memo[s[j]]; ok && index >= i {
      i = memo[s[j]] + 1
    }

    max = int(math.Max(float64(max), float64(j - i + 1)))

    memo[s[j]] = j
  }

  return max
}
