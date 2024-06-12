package main

import (
	"os"
)

func main() {
	err := "No modulo by 0"
	er := "No division by 0"
	done := os.Args[1:]
	var x, y, result int
	if len(done) == 3 {
		if Checkstring(done[0]) && Checkstring(done[2]) {
			x = atoi(done[0])
			y = atoi(done[2])
			if x >= 9223372036854775807 || y >= 9223372036854775807 || x <= -9223372036854775807 || y <= -9223372036854775807 {
				return
			}
		} else {
			return
		}
		if tchek(done[1]) >= 0 && tchek(done[1]) <= 4 {
			i := tchek(done[1])
			if i == 0 {
				result = x + y
			} else if i == 1 {
				result = x - y
			} else if i == 2 {
				if y != 0 {
					result = x / y
				} else {
					os.Stderr.WriteString(er + "\n")
					return
				}
			} else if i == 3 {
				result = x * y
			} else if i == 4 {
				if y != 0 {
					result = x % y
				} else {
					os.Stderr.WriteString(err + "\n")
					return
				}
			} else {
				os.Stderr.WriteString("\n")
				return
			}
			a := intstr(result)
			os.Stderr.WriteString(a + "\n")

		} else {
			return
		}
	} else {
		return
	}
}

// ///////////////////////////////////////////
func Checkstring(s string) bool {
	for i := 0; i < len(s); i++ {
		if i == 0 {
			if s[i] >= 48 && s[i] <= 57 {
				for j := 0; j < len(s); j++ {
					if s[j] >= 48 && s[j] <= 57 {
						continue
					} else {
						return false
					}
				}
				return true
			} else if s[i] == 45 {
				for j := i + 1; j < len(s); j++ {
					if s[j] >= 48 && s[j] <= 57 {
						continue
					} else {
						return false
					}
				}
				return true
			} else {
				return false
			}
		}
	}
	return false
}

// ///////////////////////////////////////////
func tchek(a string) int {
	operator := []string{"+", "-", "/", "*", "%"}
	for i := 0; i < len(operator); i++ {
		if a == string(operator[i]) {
			return i
		}
	}
	return 5
}

// /////////////////////////////////////////////
func atoi(s string) int {
	sum := 0
	for i := 0; i < len(s); i++ {
		if i == 0 {
			if s[i] >= 48 && s[i] <= 57 {
				for j := 0; j < len(s); j++ {
					if s[j] >= 48 && s[j] <= 57 {
						sum = sum*10 + int(s[j]-'0')
					} else {
						return 0
					}
				}
				return sum
			} else if s[i] == 45 {
				for j := i + 1; j < len(s); j++ {
					if s[j] >= 48 && s[j] <= 57 {
						sum = sum*10 + int(s[j]-'0')
					} else {
						return 0
					}
				}
				return sum * (-1)
			} else if s[i] == 43 {
				for j := i + 1; j < len(s); j++ {
					if s[j] >= 48 && s[j] <= 57 {
						sum = sum*10 + int(s[j]-'0')
					} else {
						return 0
					}
				}
				return sum
			} else {
				return 0
			}
		}
	}
	return 0
}

// ///////////////////////////////////////////////
func intstr(n int) string {
	a := ""
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
			a += string(y[i])
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
		a += string('-')
		for j := len(x) - 1; j >= 0; j-- {
			a += string(x[j])
		}
	}
	return a
}
