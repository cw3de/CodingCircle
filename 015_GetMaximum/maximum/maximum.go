package maximum

import "math"

// Finde das Maxmimum ohne Verwendung von if/switch/ternär, Vergleichen oder abs()
func Maximum(a, b float64) float64 {
	// build abs with sqrt
	return ((a + b) + math.Sqrt((a-b)*(a-b))) / 2
}

func IntMax(a, b int) int {
	// build abs with bitshift
	return a + (a-b)*((a-b)>>31)
}
