package math

func Sum(x ...int) int {
	r := 0
	for _, v := range x {
		r += v
	}
	return r
}

func Multuply(a, b int) int {
	return a * b
}
