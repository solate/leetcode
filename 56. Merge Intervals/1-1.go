package main

import (
	"fmt"
	"sort"
)

/**
Range (Int, Int):

Examples:

(1, 5) => [1, 2, 3, 4, 5]

(2, 2) => [2]

mergeRanges examples:

1. Input: [(1, 3), (2, 4)] => output: [(1, 4)]

2. Input: [(1, 2), (4, 5)] => output: [(1, 2), (4, 5)]

3. Input: [(1, 4), (2, 3), (6, 6)] => output: [(1, 4), (6, 6)]

[1, 1], [3, 3], [5, 5], [7, 7]

[{1,1}, {3,3}] [{5,5}, {7, 7}]

{1,1} {5,5}  => 取是否可以组合在一起 [{1,1}, {5, 5}]


*/

//https://zhuanlan.zhihu.com/p/33114095 参考 因为用了collections.sort所以是O(nlogn)
func main() {

	qs := [][]Interval{
		{
			Interval{1, 2},
			Interval{1, 3},
			Interval{1, 4},
			Interval{3, 3},
		},
		{
			Interval{1, 3},
			Interval{2, 4},
		},
		{
			Interval{1, 2},
			Interval{4, 5},
		},
		{
			Interval{1, 4},
			Interval{2, 3},
			Interval{6, 6},
		},
	}

	for _, v := range qs {
		fmt.Println(merge(v))
	}

}

type Interval struct {
	Start int
	End   int
}

//时间复杂度O(nlogn)
func merge(its []Interval) []Interval {

	if len(its) < 2 {

		return its
	}

	//排序, 保证都是从小到大
	sort.Slice(its, func(i, j int) bool {
		return its[i].Start < its[j].Start
	})

	length := len(its)
	result := make([]Interval, 0, length)

	//比较目标
	target := its[0] //第一个元素

	for i := 1; i < length; i++ {
		if its[i].Start <= target.End { //下个元素开始小于目标元素的结束,也就是有交集
			target.End = max(target.End, its[i].End) //那么结束就是最大的那个值
		} else {
			//没交集
			result = append(result, target) //就存到结果里
			target = its[i]                 //目标换成当前没交集的这个元素
		}
	}

	//将最后一个区间存入
	result = append(result, target)

	return result

}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
