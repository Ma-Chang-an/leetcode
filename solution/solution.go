package solution

import "fmt"

func Solution() {
	costs := [][]int{{17, 2, 17}, {16, 16, 5}, {14, 3, 19}}
	ret := minCost(costs)
	fmt.Println(ret)
}
