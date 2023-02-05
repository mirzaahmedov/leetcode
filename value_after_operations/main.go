// https://leetcode.com/problems/final-value-of-variable-after-performing-operations/

package main

import "fmt"

func main() {
	fmt.Println(finalValueAfterOperations([]string{"--X", "X++", "X++"}))
}

func finalValueAfterOperations(operations []string) int {
	value := 0

	for i := 0; i < len(operations); i++ {
		if operations[i] == "X++" || operations[i] == "++X" {
			value += 1
		} else {
			value -= 1
		}
	}

	return value
}
