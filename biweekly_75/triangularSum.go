package main

import "fmt"

func triangularSum(nums []int) int {
	n := len(nums)
	m := make([]int, n)

	for i := n - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			m[j] = (nums[j] + nums[j+1]) % 10
		}
		copy_slice(m, nums, i+1)
	}
	return nums[0]

}

func copy_slice(source, target []int, len int) {
	for i := 0; i < len; i++ {
		target[i] = source[i]
	}
}

func main() {
	a := []int{1, 2, 3, 4, 5}
	sum := triangularSum(a)
	fmt.Println(sum)
}
