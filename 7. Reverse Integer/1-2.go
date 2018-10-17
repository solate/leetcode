package main

import (
	"fmt"
	"math"
)

func main() {

	num := 123

	fmt.Println(reverse2(num))

}

//反转整数
func reverse2(x int) int {

	var rev int64
	for x != 0 {
		pop := x % 10
		x /= 10

		rev = rev*10 + int64(pop)

	}

	if rev > math.MaxInt32 || rev < math.MinInt32 {
		return 0
	}

	return int(rev)
}
