// https://leetcode.com/problems/reverse-bits/?envType=study-plan&id=algorithm-i
package main

func main(){
}
func reverseBits(num uint32) uint32 {
  var (
    result uint32
    i uint32
  )

  for i = 0; i < 32; i++ {
    if (num & (1 << i)) == (1 << i) {
      result = result | (1 << (31 - i))
    }
  }

  return result
}
