package piscine

func StringToIntSlice(str string) []int {
	var w []int
	new := []rune(str)
	for i := 0; i < len(new); i++ {
		w = append(w, int(new[i]))
	}
	return w
}
