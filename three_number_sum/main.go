package main

import (
	"fmt"
	"math"
	"sort"
)

func getnum(nums []int, target int) int {
	sort.Ints(nums)
	closesum := nums[0] + nums[1] + nums[2]
	mindiff := math.Abs(float64(target - closesum))

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left := i + 1
		right := len(nums) - 1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			diff := math.Abs(float64(target - sum))
			if diff < mindiff {
				closesum = sum
				mindiff = diff
			}

			if sum < target {
				left++
			} else if sum > target {
				right--
			} else {
				return closesum
			}
		}

	}

	return closesum
}

func main() {
	nums := []int{-9,8,2,3}

	target := 1
	fmt.Println(getnum(nums, target))
}
