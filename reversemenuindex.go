package piscine

func ReverseMenuIndex(menu []string) []string {
	var w []string
	for i := len(menu) - 1; i >= 0; i-- {
		w = append(w, menu[i])
	}
	return w
}
