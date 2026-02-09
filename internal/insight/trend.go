package insight

func PercentChange(prev, curr int) float64 {
	if prev == 0 {
		if curr == 0 {
			return 0
		}
		return 100
	}
	return (float64(curr-prev) / float64(prev)) * 100
}
