// https://leetcode.com/problems/two-sum/
package main

func main(){
}
func twoSum(nums []int, target int) (res []int) {
    var index = make(map[int]int, len(nums))

    for i, num := range nums {
        index[num] = i
    }
    
    for i, num := range nums {
        num = num - target
        if j, ok := index[-num]; ok && i != j {
            res = append(res, i, j)
            break
        }
    }

    return
}
