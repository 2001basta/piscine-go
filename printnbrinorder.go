package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrInOrder(n int) {
	x := n
	var som rune
	var y []rune
	for i := 0; i <= n; i++ {
		som = rune('0' + x%10)
		x = x / 10
		y = append(y, som)
		if x == 0 {
			break
		}
	}

	i := 1
	for i < len(y) {
		x := y[i-1]
		if y[i-1] > y[i] {
			y[i-1] = y[i]
			y[i] = x
			i = 1
		} else {
			i++
		}
	}
	for k := 0; k < len(y); k++ {
		z01.PrintRune(y[k])
	}
}
