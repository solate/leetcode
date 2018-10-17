package main

import (
	"fmt"
	"math"
)

func main() {

	num := 123

	fmt.Println(reverse(num))

}

/*

//pop operation:
pop = x % 10;
x /= 10;

//push operation:
temp = rev * 10 + pop; //这句可能会导致溢出   pop > 7 会导致溢出
rev = temp;

因为2的31次方到 2的31次方-1

最大值是2147483647
最小值-2147483648

所以尾数大于7会溢出

*/

//反转整数
func reverse(x int) int {

	var rev int
	for x != 0 {
		pop := x % 10
		x /= 10

		if rev > math.MaxInt32/10 || (rev == math.MaxInt32/10 && pop > 7) {
			return 0
		}
		if rev < math.MinInt32/10 || (rev == math.MinInt32/10 && pop < -8) {
			return 0
		}

		rev = rev*10 + pop

	}

	return rev
}
