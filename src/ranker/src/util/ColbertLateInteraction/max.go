package ColbertLateInteraction

func getMax(row []float64) float64 {
	var max float64 = 0.0
	for _, e := range row {
		if e > max {
			max = e
		}
	}
	return max
}
