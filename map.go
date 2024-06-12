package piscine

func Map(f func(int) bool, a []int) []bool {
	var w []bool
	var x bool
	for i := 0; i < len(a); i++ {
		x = f(a[i])
		w = append(w, x)
	}
	return w
}
