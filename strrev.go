package piscine

func StrRev(s string) string {
	var x string
	var y string
	for i := len(s) - 1; i >= 0; i-- {
		x = string(s[i])
		y += x
	}
	return y
}
