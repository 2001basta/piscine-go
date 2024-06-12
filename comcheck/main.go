package main

import (
	"fmt"
	"os"
)

func main() {
	w := []string{"01", "galaxy", "galaxy 01"}
	x := os.Args[1:]
	a := 0
	for i := 0; i < len(x); i++ {
		for _, j := range w {
			if x[i] == j {
				a++
			} else {
				continue
			}
		}
	}
	if a != 0 {
		fmt.Println("Alert!!!")
	}
}
