package js

import "math"

type Float float64

func (f Float) Pow(n float64) float64 {
	return math.Pow(float64(f), n)
}
