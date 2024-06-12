package piscine

func ShoppingListSort(slice []string) []string {
	i := 1
	for i < len(slice) {
		if len(slice[i]) < len(slice[i-1]) {
			slice[i-1], slice[i] = slice[i], slice[i-1]

			i = 1
		} else {
			i++
			continue
		}
	}
	return slice
}
