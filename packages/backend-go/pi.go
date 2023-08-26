package main

func CalcPI(N int) float64 {
	// pi/4 = (1/N)Sigma(i to N) 1/(1+ ((i-1/2)/N)^2)

	var summedup float64 = 0

	for i := 0; i < N; i++ {
		var sq float64 = (float64(i) - 0.5) / float64(N)
		sq *= sq
		var innerExpr = 1 / (1 + sq)
		summedup += innerExpr
	}

	return summedup * 4 / float64(N)
}
