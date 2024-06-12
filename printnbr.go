package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbr(n int) {
	if n >= 0 {
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
		for i := len(y) - 1; i >= 0; i-- {
			z01.PrintRune(y[i])
		}
	} else {
		y := n
		var some rune
		var x []rune
		for i := 0; i <= (n / 2 * (-1)); i++ {
			some = rune('0' + ((y % 10) * (-1)))
			y = y / 10
			x = append(x, some)
			if y == 0 {
				break
			}
		}
		z01.PrintRune('-')
		for j := len(x) - 1; j >= 0; j-- {
			z01.PrintRune(x[j])
		}
	}
}
