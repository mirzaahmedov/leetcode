// https://leetcode.com/problems/maximum-number-of-words-found-in-sentences/

package main

import (
  "fmt"
)

func main() {
  fmt.Println(mostWordsFound([]string{ 
    "alice and bob love leetcode",
    "i think so too",
    "this is great thanks very much",
  }))
}

func mostWordsFound(sentences []string) int {
    maxCount := 0

    for i := 0; i < len(sentences); i++ {
      count := 0
      for j := 0; j < len(sentences[i]); j++ {
        if sentences[i][j] == ' ' {
          count += 1
        }
      }

      count += 1

      if count > maxCount {
        maxCount = count
      }
    }

    return maxCount
}
