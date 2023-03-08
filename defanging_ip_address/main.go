// https://leetcode.com/problems/defanging-an-ip-address/

package main

import "fmt"

func main(){
  fmt.Println(defangIPaddr("12.122.33.22"))
}

func defangIPaddr(address string) string {
  res := make([]byte, len(address) + 6)
  j := 0

  for i := 0; i < len(address); i++ {
    if address[i] == '.' {
      res[j] = '['
      res[j + 1] = '.'
      res[j + 2] = ']'

      j += 3
    } else {
      res[j] = address[i]
      j += 1
    }
  } 

  return string(res)
}
