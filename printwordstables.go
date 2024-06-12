package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for i := 0; i < len(a); i++ {
		x := []rune(a[i])
		for _, y := range x {
			z01.PrintRune(y)
		}
		z01.PrintRune('\n')
	}
}
