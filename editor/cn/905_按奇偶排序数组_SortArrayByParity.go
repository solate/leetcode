package leetcode

//给定一个非负整数数组 A，返回一个数组，在该数组中， A 的所有偶数元素之后跟着所有奇数元素。
//
// 你可以返回满足此条件的任何数组作为答案。
//
//
//
// 示例：
//
// 输入：[3,1,2,4]
//输出：[2,4,3,1]
//输出 [4,2,3,1]，[2,4,1,3] 和 [4,2,1,3] 也会被接受。
//
//
//
//
// 提示：
//
//
// 1 <= A.length <= 5000
// 0 <= A[i] <= 5000
//
// Related Topics 数组 双指针 排序 👍 228 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func sortArrayByParity(nums []int) []int {

	left, right := 0, len(nums)-1

	for left < right {

		if nums[left]%2 == 0 {
			left++
		} else if nums[right]%2 == 1 {
			right--
		} else {
			nums[left], nums[right] = nums[right], nums[left]
		}
	}

	return nums

}

//leetcode submit region end(Prohibit modification and deletion)

func sortArrayByParity1(nums []int) []int {

	j := 0
	n := len(nums)

	for i := 0; i < n; i++ {
		if nums[i]%2 == 0 {
			nums[i], nums[j] = nums[j], nums[i]
			j++
		}
	}

	return nums

}

func sortArrayByParity2(nums []int) []int {

	left, right := 0, len(nums)-1

	for left < right {

		if nums[left]%2 == 0 {
			left++
		} else if nums[right]%2 == 1 {
			right--
		} else {
			nums[left], nums[right] = nums[right], nums[left]
		}
	}

	return nums

}
