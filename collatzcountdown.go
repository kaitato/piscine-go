package piscine

func CollatzCountdown(start int) int {
	count := 0
	if start < 1 {
		return -1
	}
	if start != 1 {
		if start%2 == 0 {
			start /= 2
		} else {
			start = start*3 + 1
		}
		count++
	}
	return count
}
