package leetcode

import (
	"sort"
)

//给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重
//复的三元组。
//
// 注意：答案中不可以包含重复的三元组。
//
//
//
// 示例 1：
//
//
//输入：nums = [-1,0,1,2,-1,-4]
//输出：[[-1,-1,2],[-1,0,1]]
//
//
// 示例 2：
//
//
//输入：nums = []
//输出：[]
//
//
// 示例 3：
//
//
//输入：nums = [0]
//输出：[]
//
//
//
//
// 提示：
//
//
// 0 <= nums.length <= 3000
// -10⁵ <= nums[i] <= 10⁵
//
// Related Topics 数组 双指针 排序 👍 4163 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
func threeSum(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)

	//最小数字是k,   双指针i,j 分别为 k, n-1

	for k := 0; k < n; k++ {

		// 如果最小数，nums[k] >= num[i] >= num[j] > 0, 则没有
		if nums[k] > 0 {
			break
		}

		// 当 nums[k] == nums[k-1] 说明两个数字相同，已加入跳过
		if k > 0 && nums[k] == nums[k-1] {
			continue
		}

		i, j := k+1, n-1
		for i < j {
			//fmt.Println(i, j, k, "--")
			s := nums[k] + nums[i] + nums[j]
			if s < 0 {
				i++
				for i < j && nums[i] == nums[i-1] {
					i++
				}
			} else if s > 0 {
				j--
				for i < j && nums[j] == nums[j+1] {
					j--
				}
			} else {
				ans = append(ans, []int{nums[k], nums[i], nums[j]})
				i++
				j--
				for i < j && nums[i] == nums[i-1] {
					i++
				}
				for i < j && nums[j] == nums[j+1] {
					j--
				}
			}

		}

	}

	return ans
}

//leetcode submit region end(Prohibit modification and deletion)

func threeSum1(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)

	// 枚举 a
	for first := 0; first < n; first++ {

		// 排除到相同的数字
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}

		// c 对应的指针初始值指向数组的最右端
		third := n - 1
		target := -1 * nums[first] // 两数只和解法升级版 a + b = -c

		// 枚举 b
		for second := first + 1; second < n; second++ {

			//排除相同的值
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}

			// 需要保证 b的指针在c的指针的左侧
			for second < third && nums[second]+nums[third] > target {
				third--
			}
			// 如果指针重合，随着 b 后续的增加
			// 就不会有满足 a+b+c=0 并且 b<c 的 c 了，可以退出循环
			if second == third {
				break
			}

			if nums[second]+nums[third] == target {
				ans = append(ans, []int{nums[first], nums[second], nums[third]})
			}

		}

	}

	return ans
}

func threeSum2(nums []int) [][]int {
	n := len(nums)
	sort.Ints(nums)
	ans := make([][]int, 0)

	//最小数字是k,   双指针i,j 分别为 k, n-1

	for k := 0; k < n; k++ {

		// 如果最小数，nums[k] >= num[i] >= num[j] > 0, 则没有
		if nums[k] > 0 {
			break
		}

		// 当 nums[k] == nums[k-1] 说明两个数字相同，已加入跳过
		if k > 0 && nums[k] == nums[k-1] {
			continue
		}

		i, j := k+1, n-1
		for i < j {
			//fmt.Println(i, j, k, "--")
			s := nums[k] + nums[i] + nums[j]
			if s < 0 {
				i++
				for i < j && nums[i] == nums[i-1] {
					i++
				}
			} else if s > 0 {
				j--
				for i < j && nums[j] == nums[j+1] {
					j--
				}
			} else {
				ans = append(ans, []int{nums[k], nums[i], nums[j]})
				i++
				j--
				for i < j && nums[i] == nums[i-1] {
					i++
				}
				for i < j && nums[j] == nums[j+1] {
					j--
				}
			}

		}

	}

	return ans
}
