// https://leetcode.com/problems/group-anagrams/submissions/927911382/
package main

func main(){

}
func groupAnagrams(strs []string) [][]string {
  groups := make(map[[26]int][]string)
  
  for i := 0; i < len(strs); i++ {
    counts := [26]int{}

    for j := 0; j < len(strs[i]); j++ {
      counts[strs[i][j] - 'a']++
    }

    groups[counts] = append(groups[counts], strs[i])
  }

  results := [][]string{}

  for key := range groups {
    results = append(results, groups[key])
  }

  return results
}
