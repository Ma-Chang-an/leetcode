package solution

func findStrobogrammatic(n int) []string {
	if n == 0 {
		return nil
	}
	if n == 1 {
		return []string{"0", "1", "8"}
	}
	if n == 2 {
		return []string{"11", "69", "88", "96"}
	}
	tmp := findStrobogrammatic(n - 2)
	var ret []string
	if n%2 == 0 {
		for _, item := range tmp {
			ret = append(ret, item[0:len(item)/2]+"00"+item[len(item)/2:])
			ret = append(ret, item[0:len(item)/2]+"11"+item[len(item)/2:])
			ret = append(ret, item[0:len(item)/2]+"69"+item[len(item)/2:])
			ret = append(ret, item[0:len(item)/2]+"88"+item[len(item)/2:])
			ret = append(ret, item[0:len(item)/2]+"96"+item[len(item)/2:])
		}
		return ret
	}
	for _, item := range tmp {
		ret = append(ret, "1"+item+"1")
		ret = append(ret, "8"+item+"8")
		ret = append(ret, "6"+item+"9")
		ret = append(ret, "9"+item+"6")
	}
	return ret
}
