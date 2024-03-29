package array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

/**
给定一个排序数组，你需要在 原地 删除重复出现的元素，使得每个元素只出现一次，返回移除后数组的新长度。

不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。


输入：nums = [1,1,2]
输出：2, nums = [1,2]
解释：函数应该返回新的长度 2 ，并且原数组 nums 的前两个元素被修改为 1, 2 。不需要考虑数组中超出新长度后面的元素。

输入：nums = [0,0,1,1,1,2,2,3,3,4]
输出：5, nums = [0,1,2,3,4]
解释：函数应该返回新的长度 5 ， 并且原数组 nums 的前五个元素被修改为 0, 1, 2, 3, 4 。不需要考虑数组中超出新长度后面的元素。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-duplicates-from-sorted-array
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func removeDuplicates(nums []int) int {
	size := len(nums)

	if size <= 1 {
		return size
	}
	slow := 0
	for fast := 1; fast < size; fast++ {
		if nums[slow] != nums[fast] {
			nums[slow+1] = nums[fast]
			slow++
		}
	}
	return slow + 1
}

func TestRemoveDuplicates(t *testing.T) {
	nums := []int{1, 1, 2}
	size := removeDuplicates(nums)
	assert.Equal(t, 2, size)

	nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	size = removeDuplicates(nums)
	assert.Equal(t, 5, size)
}
