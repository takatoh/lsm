package lsm

func LSM1(xs, ys []float64) (float64, float64) {
	var sum_xy, sum_x, sum_y, sum_x2 float64

	n := len(xs)
	for i := 0; i < n; i++ {
		sum_xy += x[i] * y[i]
		sum_x += x[i]
		sum_y += y[i]
		sum_x2 += x[i] * x[i]
	}

	nf = float64(n)
	a := (nf * sum_xy - sum_x * sum_y) / (nf * sum_x2 - sum_x * sum_x)
	b := (sum_x2 * sum_y - sum_xy * sum_x) / (nf * sum_x2 - sum_x * sum_x)

	return a, b
}
