package piscine

func Atoi(s string) int {
	sum := 0
	var x int
	var y int
	for i := 0; i < len(s); i++ {
		if i == 0 {
			if s[i] >= 48 && s[i] <= 57 {
				sum = sum*10 + int(s[i]-'0')
				x = sum
			} else if s[i] == 45 {
				for j := i + 1; j < len(s); j++ {
					if s[j] >= 48 && s[j] <= 57 {
						sum = sum*10 + int(s[j]-'0')
						y = sum * (-1)
					} else {
						return 0
					}
				}
				return y
			} else if s[i] == 43 {
				for j := i + 1; j < len(s); j++ {
					if s[j] >= 48 && s[j] <= 57 {
						sum = sum*10 + int(s[j]-'0')
						y = sum
					} else {
						return 0
					}
				}
				return y
			} else {
				return 0
			}
		} else {
			if s[i] >= 48 && s[i] <= 57 {
				sum = sum*10 + int(s[i]-'0')
				x = sum
			} else {
				return 0
			}
		}
	}
	return x
}
