/*
[Introduction] - 1

Given a binary array nums,
return the maximum number of consecutive 1's in the array.
*/

package main

import "fmt"

func main() {

	nums := []int{1, 1, 0, 1, 1, 1, 1, 0}
	max := findMaxConsecutiveOnes(nums)
	fmt.Println(max)
}

/* Solution */

func findMaxConsecutiveOnes(nums []int) int {
	var max, cnt int = 0, 0
	for _, num := range nums {
		if num == 1 {
			cnt++
			if max < cnt {
				max = cnt
			}
		} else {
			cnt = 0
		}
	}
	return max
}
