// https://leetcode.com/problems/integer-to-roman/
package main

import "fmt"

// 9 - 10 = -1
// 9 - 5 = 4
// 4 - 1 = 3
// 3 - 1 = 2
// 2 - 1 = 1
// counter = 3
// 

func main(){
  // fmt.Println(intToRoman(3))
  // fmt.Println(intToRoman(4)) // IV
  // fmt.Println(intToRoman(5))
  // fmt.Println(intToRoman(6))
  // fmt.Println(intToRoman(9)) // IV
  // fmt.Println(intToRoman(12))
  fmt.Println(intToRoman(1994))
}

func intToRoman(num int) string {
  roman := ""

  symbols := map[int]byte {
    1: 'I',
    5: 'V', 
    10: 'X',
    50: 'L',
    100: 'C',
    500: 'D',
    1000: 'M',
  }

  for _, value := range []int{ 1000, 500, 100, 50, 10, 5, 1 } {
    symbol := symbols[value]

    if num - value >= 0 {
      for num - value >= 0 {
        roman += string(symbol)
        num -= value
      }
    } 
    if num - value < 0 {
      rem := value - num

      if value >= 500 {
        if rem - 100 <= 0 {
          roman += "C" + string(symbol)
          num = 100 - rem
        }
      } else if value >= 50 {
        if rem - 10 <= 0 {
          roman += "X" + string(symbol)
          num = 10 - rem
        }
      } else if value >= 5 {
        if rem - 1 <= 0 {
          roman += "I" + string(symbol)
          num = 1 - rem
        }
      }

    }
  }

  return roman
}
