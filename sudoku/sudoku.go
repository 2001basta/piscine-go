package main

import (
	"fmt"
	"os"
)

func main() {
	table := os.Args
	var sudoku []int
	var table_final [][]int
	for i := 1; i < len(table); i++ {
		for k := 0; k < len(table[i]); k++ {
			if table[i][k] == '.' {
				sudoku = append(sudoku, 0)
			} else {
				sudoku = append(sudoku, int(table[i][k]-'0'))
			}
		}
		table_final = append(table_final, sudoku)
		sudoku = nil
	}
	for _, t := range table_final {
		fmt.Println(t)
	}
	fmt.Print("\n\n")
	y := []int{}
	for i := 0; i < len(table_final); i++ {
		for j := 0; j < len(table_final[i]); j++ {
			if table_final[i][j] == 0 {
				k := 1
				for k <= 9 {
					w := 0
					for w < len(table_final[i]) {
						if table_final[i][w] == k {
							k++
							w = 0
						} else {
							w++
						}
					}
					table_final[i][j] = k
					break
				}
			} else {
				continue
			}
		}
		fmt.Println(table_final[i])
	}
	fmt.Print("\n\n")

	/*table := os.Args
	var sudoku []int
	var table_final [][]int
	for i := 1; i < len(table); i++ {
		for k := 0; k < len(table[i]); k++ {
			if table[i][k] == '.' {
				sudoku = append(sudoku, 0)
			} else {
				sudoku = append(sudoku, int(table[i][k]-'0'))
			}
		}
		table_final = append(table_final, sudoku)
		sudoku = nil
	}
	for _, t := range table_final {
		fmt.Println(t)
	}
	fmt.Print("\n\n")

	for i := 0; i < len(table_final); i++ {
		for j := 0; j < len(table_final[i]); j++ {
			if table_final[i][j] == 0 {

				k := 1
				w := 0
				// clone := 0
				// index := 0
				for k <= 9 {
					for w < len(table_final[i]) {
						if table_final[i][w] == k {
							k++
							w = 0
						} else {
							w++
						}
					}
					y = append(y, k)
					break
				}
				// k = 1

				/*for index <= 3 {
					for k <= 9 {
						for clone < len(table_final) {
							if table_final[clone][index] == k {
								k++
								clone = 0
							} else {
								clone++
							}
						}
						y = append(y, k)
					}
					index++

				}

			} else {
				continue
			}
		}
		fmt.Println(table_final[i])
	}*/
	fmt.Print(y)

	fmt.Print("\n\n")
}
