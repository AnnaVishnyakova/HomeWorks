package Average

func Average(xs []float64) float64 {
	var total float64
	var NotZeroCount int
	if len(xs) == 0 {
		return total
	}
	for _, x := range xs {
		if x != 0 {
			NotZeroCount++
		}
		total += x
	}
	return total / float64(NotZeroCount)
}

func Sum(yz []float64) float64 {
	var total2 float64

	if len(yz) == 0 {
		return total2
	}
	for _, y := range yz {

		total2 += y
	}
	return total2
}
