package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	nom := os.Args[0]
	slic := []rune(nom)
	for i := 2; i < len(slic); i++ {
		z01.PrintRune(slic[i])
	}
	z01.PrintRune('\n')
}
