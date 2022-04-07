/*
[Introduction] - 3
Given an integer array nums sorted in non-decreasing order, return an array of the squares of each number sorted in non-decreasing order.
*/
package main

import "fmt"

func main() {
	nums := []int{-7, -5, -3, 2, 3, 9, 11}
	newNums := sortedSquares(nums)
	fmt.Println((newNums))
}

/* Solution */
func sortedSquares(nums []int) []int {
	var x, y = 0, len(nums) - 1
	arr := make([]int, y+1)
	for z := y; z >= 0; z-- {
		if nums[x]*nums[x] > nums[y]*nums[y] {
			arr[z] = nums[x] * nums[x]
			x++
		} else {
			arr[z] = nums[y] * nums[y]
			y--
		}
	}
	return arr
}
