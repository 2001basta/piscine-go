package piscine

func BasicJoin(elems []string) string {
	var x string
	for i := 0; i < len(elems); i++ {
		x += elems[i]
	}
	return x
}
