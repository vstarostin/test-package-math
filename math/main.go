package math

func Sum(x ...int) int {
	r := 0
	for _, c := range x {
		r += c
	}
	return r
}

func Multuply(x ...int) int {
	r := 1
	for _, c := range x {
		r *= c
	}
	return r
}

func Pow(x, p int) float64 {
	n := p
	if p < 1 {
		n = -p
	}
	xx := 1.0
	for i := 0; i < n; i++ {
		xx *= float64(x)
	}
	if p < 1 {
		xx = 1 / xx
	}
	return xx
}
