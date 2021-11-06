package solution

//84. 柱状图中最大的矩形
//给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。
//
//求在该柱状图中，能够勾勒出来的矩形的最大面积。
//示例 1:
//输入：heights = [2,1,5,6,2,3]
//输出：10
//解释：最大的矩形为图中红色区域，面积为 10
//示例 2：
//输入： heights = [2,4]
//输出： 4

//提示：
//
//1 <= heights.length <=105
//0 <= heights[i] <= 104
//暴力解答93/96超时
//func largestRectangleArea(heights []int) int {
//	var maxArea = heights[0]
//	for i, height := range heights {
//		var left, right int
//		for left = i - 1; left >= 0; left-- {
//			if heights[left] < height {
//				break
//			}
//		}
//		for right = i; right < len(heights); right++ {
//			if heights[right] < height {
//				break
//			}
//		}
//		curArea := height * (right - left - 1)
//		if curArea > maxArea {
//			maxArea = curArea
//		}
//	}
//	return maxArea
//}

func largestRectangleArea(heights []int) int {
	heights = append(heights, 0)
	var maxArea = heights[0]
	var stack []int
	for i, height := range heights {
		for len(stack) > 0 && heights[stack[len(stack)-1]] >= height {
			maxHeight := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			sidx := -1
			if len(stack) > 0 {
				sidx = stack[len(stack)-1]
			}
			curArea := maxHeight * (i - sidx - 1)
			if curArea > maxArea {
				maxArea = curArea
			}
		}
		stack = append(stack, i)
	}
	return maxArea
}
