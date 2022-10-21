package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(nums, target)) // output: [1,0]
}
func twoSum(nums []int, target int) []int {
	indices := []int{0, 0}
	for index1, value1 := range nums {

		for index2, value2 := range nums {
			if index1 == index2 {
				continue
			}
			if value1+value2 == target {
				indices[0] = index1
				indices[1] = index2
				break
			}
		}
	}
	return indices
}

/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.
/*