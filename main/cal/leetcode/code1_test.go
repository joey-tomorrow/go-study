package leetcode

import (
	"fmt"
	"testing"
)

/*
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。



示例:

给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/two-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func twoSum(nums []int, target int) []int {
	indexMap := make(map[int][]int)
	for k, v := range nums {
		value, ok := indexMap[v]
		if ok {
			value = append(value, k)
			indexMap[v] = value
		} else {
			indexMap[v] = []int{k}
		}
	}

	result := make([]int, 0)
	for k, v := range nums {
		tmp, ok := indexMap[target-v]
		if ok {
			if target-v == v {
				if len(tmp) == 2 {
					result = append(result, tmp[0], tmp[1])
					return result
				}
			} else {
				result = append(result, k, tmp[0])
				return result
			}
		}
	}

	//for i := 0; i <= target; i++ {
	//	first, ok := indexMap[i]
	//	if ok {
	//		if len(first) == 1 {
	//			second, yes := indexMap[target-i]
	//			if yes {
	//				result = append(result, first[0], second[0])
	//				return result
	//			}
	//		} else {
	//			if i*2 == target {
	//				result = append(result, first[0], first[1])
	//				return result
	//			}
	//		}
	//	}
	//}

	return result
}

func twoSum1(nums []int, target int) []int {
	result := make([]int, 0)
	indexMap := make(map[int]int)
	for k, v := range nums {
		left := target - v
		tmp, ok := indexMap[left]
		if ok {
			result = append(result, k, tmp)
		} else {
			indexMap[v] = k
		}
	}

	return result
}

func Test_TwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}

	fmt.Println(twoSum1(nums, 26))

	nums1 := []int{3, 3}

	fmt.Println(twoSum1(nums1, 6))

	nums2 := []int{-3, 4, 3, 90}
	fmt.Println(twoSum1(nums2, 0))

	nums3 := []int{3, 2, 4}
	fmt.Println(twoSum1(nums3, 6))
}
