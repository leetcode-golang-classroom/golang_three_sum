package sol

import "sort"

func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	result := [][]int{}
	sort.Ints(nums)
	n := len(nums)
	// fixed index pivot
	for pivot := 1; pivot < n-1; pivot++ {
		start, end := 0, n-1
		// do two sum for fixed index pivot
		if pivot > 1 && nums[pivot] == nums[pivot-1] {
			start = pivot - 1
		}
		for start < pivot && pivot < end {
			if start > 0 && nums[start] == nums[start-1] { // update search range
				start++
				continue
			}
			if end < n-1 && nums[end] == nums[end+1] { // update search range
				end--
				continue
			}
			sum := nums[start] + nums[pivot] + nums[end]
			if sum == 0 {
				result = append(result, []int{nums[start], nums[pivot], nums[end]})
				start++
				end--
			} else if sum > 0 {
				end--
			} else {
				start++
			}
		}
	}
	return result
}
