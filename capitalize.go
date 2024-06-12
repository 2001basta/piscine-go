package piscine

func Capitalize(s string) string {
	var w string
	var x byte
	for i := 0; i < len(s); i++ {
		x = s[i]
		if i == 0 {
			if x >= 'a' && x <= 'z' {
				x = x - 32
			} else {
				x = x + 0
			}
		} else {
			y := s[i-1]
			if (x >= 'a' && x <= 'z') || (x >= 'A' && x <= 'Z') {
				if x >= 'a' && x <= 'z' {
					if (y >= 'a' && y <= 'z') || (y >= 'A' && y <= 'Z') || (y >= '0' && y <= '9') {
						x = x + 0
					} else {
						x = x - 32
					}
				} else {
					if (y >= 'a' && y <= 'z') || (y >= '0' && y <= '9') || (y >= 'A' && y <= 'Z') {
						x = x + 32
					} else {
						x = x + 0
					}
				}
			}
		}
		w += string(x)
	}
	return w
}
