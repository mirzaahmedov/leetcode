// https://leetcode.com/problems/move-zeroes/?envType=study-plan&id=algorithm-i

package main

import "fmt"

func main(){
  nums := []int{ 0,1,0,3,12 }

  moveZeroes(nums)

  fmt.Println(nums)
}
func moveZeroes(nums []int){
  var (
    i int = 0
    j int = 0
  )
  
  for i < len(nums) {
    
    if nums[i] != 0 {
      nums[j], nums[i] = nums[i], nums[j]
      j++
    }

    i++
  }
}
