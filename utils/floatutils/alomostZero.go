package floatutils

import "math"

func AlmostEquals(a, b, e float64) bool {

	if a == b {
		return true
	}

	d := math.Abs(a - b)

	if b == 0 {
		return d < e
	}

	return (d / math.Abs(b)) < e
}
