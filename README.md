# golang_three_sum

Given an integer array nums, return all the triplets `[nums[i], nums[j], nums[k]]` such that `i != j`, `i != k`, and `j != k`, and `nums[i] + nums[j] + nums[k] == 0`.

Notice that the solution set must not contain duplicate triplets.

## Examples

**Example 1:**

```
Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]
Explanation:
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
The distinct triplets are [-1,0,1] and [-1,-1,2].
Notice that the order of the output and the order of the triplets does not matter.

```

**Example 2:**

```
Input: nums = [0,1,1]
Output: []
Explanation: The only possible triplet does not sum up to 0.

```

**Example 3:**

```
Input: nums = [0,0,0]
Output: [[0,0,0]]
Explanation: The only possible triplet sums up to 0.

```

**Constraints:**

- `3 <= nums.length <= 3000`
- $`-10^5$ <= nums[i] <= $10^5$`

## 解析

給定一個整數陣列 nums

要求寫一個演算法找在所有在 nums 中3個和= 0 數字的所有不重複複組合

已知目標是找到 i, j , k , i < j < k , 使得 nums[i] + nums[j] + nums[k] =0

可以注意到 當 nums[j] 確定時, nums[i] + nums[k] 只能 - nums[j] 

所以可透過 類似 two sum 的作法

另外為了可以使用 two pointer 來縮小範圍

所以將原本的nums 做 sort 來做由小到大的排列

初始化 start = 0, index = 1, end = len(nums) - 1, result = []

當針對所有可能的 index  做檢查(使用 two pointer)

![](https://i.imgur.com/D9AOUSA.png)

為了避免重複檢查

所以如果 nums[index] = nums[index-1] 時 且 index > 1 代表已經nums[index]找過了

這時需要把 start = index - 1 讓搜索範圍變小

而對於 start < index 且 end > index 時

為了避免重複檢查 

當nums[start] = nums[start-1] 且 start > 0 代表此搜索範圍已經找過了

這時需要把 start += 1 重新搜索 

當nums[end] = nums[end+1] 且 end < len(nums) - 1 代表此搜索範圍已經找過了

這時需要把 end -= 1 重新搜索 

令 addNum =  nums[start] + nums[index] + nums[end]

當 addNum = 0 代表找到符合條件的組合

把 [nums[start], nums[index], nums[end]] 加入 result , 更新 start += 1, end -= 1

當 addNum > 0 因為 start 已經是最小值 所以只能把 end 往左移

更新 end -= 1

當 addNum < 0 因為 end 已經是最大值 所以只能把  start 往右移

更新 start += 1

## 程式碼
```go
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

```
## 困難點

1. 需要想出避免重複的方法
2. 需要想出對邊界值移動的條件

## Solve Point

- [x]  初始化 i = 1 , result = []
- [x]  針對每個 i = 1.. len(nums)-1 做以下檢查
- [x]  初始化 start = 0, end = len(nums) - 1
- [x]  當 i > 1 且 num[i] == nums[i-1] 代表此數值已經搜尋過， 則更新 start = i - 1
- [x]  當 i > start 且 i < end 做以下檢查
- [x]  如果 start > 0 且 nums[start-1] = nums[start] 代表此範圍已經搜尋過，更新 start += 1 重新搜尋
- [x]  如果 end < len(nums) - 1 且 nums[end+1] = nums[end] 代表此範圍已經搜尋過，更新 end-= 1 重新搜尋
- [x]  令 addNum = nums[start] + nums[i] + nums[end]
- [x]  當 addNum == 0 時 ,  新增 [nums[start], nums[i], nums[end]], 更新 start+=1, end -= 1, 繼續搜尋
- [x]  當 addNum > 0 時 ,  因為 start 已經是最小值所以只能把 end 左移, 更新 end -= 1, 繼續搜尋
- [x]  當 addNum < 0 時 ,  因為 end 已經是最大值所以只能把 start 右移, 更新 start += 1, 繼續搜尋
- [x]  回傳 result