package piscine

func Split(s, sep string) []string {
	var slic []string
	var x string
	if len(sep) == 1 {
		for i := 0; i < len(s); i++ {
			if string(s[i]) != sep {
				x += string(s[i])
			} else if string(s[i]) == sep && i != len(s)-1 {
				slic = append(slic, x)
				x = ""
			} else {
				slic = append(slic, x)
			}
		}
	} else {
		j := 1
		for j < len(s) {
			if s[j] != sep[1] || s[j-1] != sep[0] {
				x += string(s[j-1])
			} else {
				slic = append(slic, x)
				x = ""
				for k := 1; k < len(sep); k++ {
					j++
				}
			}

			if j == len(s)-1 {
				x += string(s[j])
				slic = append(slic, x)
			}
			j++

		}
	}
	return slic
}
