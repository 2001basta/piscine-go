package piscine

func PodiumPosition(podium [][]string) [][]string {
	j := 1
	for i := len(podium) - 2; i >= 0; i-- {
		podium[j-1], podium[i+1] = podium[i+1], podium[j-1]
		j++
	}
	return podium
}
