package solution

import "fmt"

func Solution() {
	//ribbons := []int{100000,100000,100000,100000,100000,100000,100000,100000,100000,100000,100000,100000,100000,100000,100000,100000,100000,100000,100000,100000,1,100000,100000,100000,100000,100000,100000,100000,100000,100000,100000,100000,100000,100000,100000,100000,100000,100000,100000,100000,100000,100000,100000,100000,100000,100000,100000,100000,100000,100000}
	heights := []int{2, 1, 5, 6, 2, 3}
	ret := largestRectangleArea(heights)
	fmt.Println(ret)
}
