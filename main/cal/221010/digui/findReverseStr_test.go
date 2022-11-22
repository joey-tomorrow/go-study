package digui

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	//nums := []int{4,5,6,7,0,1,2}
	//fmt.Println(findBinary(nums, 0, 6, 0))

	nums1 := []int{3,4,5,6,7,0,1,2}
	fmt.Println(findBinary(nums1, 0, 7, 2))
}

func findBinary(nums []int, start, end, target int) int{
	if len(nums) == 0 {
		return -1
	}

	if end - start <= 1{
		if nums[start] == target {
			return start
		}

		if nums[end] == target {
			return end
		}

		return -1
	}

	mid := (start + end) / 2

	if nums[start] < nums[mid] && nums[mid] >= target && nums[start] <= target{
		return findBinary(nums, start, mid, target)
	} else {
		return findBinary(nums, mid+1, len(nums) -1, target)
	}

	return -1
}

