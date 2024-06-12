package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	nom := os.Args
	i := 2
	for i < len(nom) {
		a := nom[i-1]
		b := nom[i]
		if b < a {
			nom[i-1] = b
			nom[i] = a
			i = 2
		} else {
			i++
		}
	}
	for k := 1; k < len(nom); k++ {
		for _, j := range nom[k] {
			z01.PrintRune(j)
		}
		z01.PrintRune('\n')
	}
}
