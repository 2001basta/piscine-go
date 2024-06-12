package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	nom := os.Args
	for i := len(nom) - 1; i > 0; i-- {
		for _, j := range nom[i] {
			z01.PrintRune(j)
		}
		z01.PrintRune('\n')
	}
}
