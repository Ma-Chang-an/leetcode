package solution

import "sort"

func minMeetingRooms(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	var items [][]int
	for _, interval := range intervals {
		if len(items) == 0 {
			items = append(items, interval)
			continue
		}
		for i, item := range items {
			if interval[0] >= item[1] {
				items[i] = interval
				break
			}
			if i == len(items)-1 {
				items = append(items, interval)
			}
		}
	}
	return len(items)
}
