package piscine

func Compact(ptr *[]string) int {
	count := 0
	var w []string
	for _, y := range *ptr {
		if y != "" {
			w = append(w, y)
			count++
		}
	}
	*ptr = w
	return count
}
