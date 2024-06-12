package piscine

func ConcatParams(args []string) string {
	var w string
	var y string
	for i := 0; i < len(args); i++ {
		for _, j := range args[i] {
			w += string(j)
		}
		if i < len(args)-1 {
			y += w + string("\n")
			w = ""
		} else {
			y += w
		}

	}
	return y
}
