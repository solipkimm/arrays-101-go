/*
[Introduction] - 2
Given an array nums of integers, return how many of them contain an even number of digits.
*/

package main

import "fmt"

func main() {
	nums := []int{12, 345, 2, 66, 7896, 12, 345, 2, 6, 7896}
	cnt := findNumbers2(nums)
	fmt.Println(cnt)

}

func findNumbers(nums []int) int {
	cnt := 0
	for _, num := range nums {
		dgt := 0
		for i := 1; i <= 100000; i *= 10 {
			if num >= i {
				dgt++
			}
		}
		if dgt%2 == 0 {
			cnt++
		}
	}
	return cnt
}

func findNumbers2(nums []int) int {
	cnt := 0
	for _, num := range nums {
		dgt := 0
		for num != 0 {
			num /= 10
			dgt++
		}
		if dgt%2 == 0 {
			cnt++
		}
	}
	return cnt
}
