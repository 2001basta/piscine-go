package piscine

func Fibonacci(index int) int {
	x := index - 1
	if index == 0 {
		return 0
	} else if index == 1 {
		return 1
	} else if index < 0 {
		return -1
	} else {
		return Fibonacci(x) + Fibonacci(x-1)
	}
}
